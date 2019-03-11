//usr/bin/env go run "$0" "$@"; exit "$?"
// -*- go -*-

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

func panicIf(err error, transforms ...func(error) error) {
	if err == nil {
		return
	}
	for _, transform := range transforms {
		err = transform(err)
	}
	panic(err)
}

func getLinksFromURL(url string) ([]string, error) {
	links := []string{}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		panicIf(err)
	}()
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				for _, a := range t.Attr {
					if a.Key == "href" {
						links = append(links, url+"/"+strings.TrimSuffix(a.Val, "/"))
						break
					}
				}
			}
		case tt == html.ErrorToken:
			// End of the document, we're done
			return links, nil
		}
	}
}

type releaseInfo struct {
	Origin        string
	Label         string
	Suite         string
	Version       string
	Date          time.Time
	Codename      string
	URL           string
	Architectures []string
}

func getReleaseInfo(url string) (releaseInfo, error) {
	ri := releaseInfo{}
	resp, err := http.Get(url)
	if err != nil {
		return ri, err
	}
	defer func() {
		err := resp.Body.Close()
		panicIf(err)
	}()
	lineRe := regexp.MustCompile(`(^[A-Z][a-z]+): (.*)$`)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		kv := lineRe.FindStringSubmatch(line)
		if kv == nil {
			continue
		}
		switch kv[1] {
		case "Origin":
			ri.Origin = kv[2]
		case "Label":
			ri.Label = kv[2]
		case "Suite":
			ri.Suite = kv[2]
		case "Version":
			ri.Version = kv[2]
		case "Codename":
			ri.Codename = kv[2]
		case "Date":
			d, err := time.Parse("Mon, _2 Jan 2006 15:04:05 MST", kv[2])
			if err == nil {
				ri.Date = d
			}
		case "Architectures":
			ri.Architectures = strings.Split(kv[2], " ")
		}
	}
	if reflect.DeepEqual(ri, releaseInfo{}) {
		return ri, fmt.Errorf("could not parse lines in %s", url)
	}
	ri.URL = url
	return ri, err
}

func getAptmirrorsReleaseInfos() []releaseInfo {
	mirrors := []string{
		"http://archive.debian.org/debian",
		"http://deb.debian.org/debian",
		"http://ports.ubuntu.com/ubuntu-ports",
		"http://archive.ubuntu.com/ubuntu",
		"http://old-releases.ubuntu.com/ubuntu",
	}
	releaseInfos := []releaseInfo{}
	for _, m := range mirrors {
		url := m + "/dists" // /${dist}/main/binary-${arch}/Release
		links, err := getLinksFromURL(url)
		panicIf(err)
		for _, link := range links {
			ri, err := getReleaseInfo(link + "/Release")
			if err != nil {
				continue
			}
			releaseInfos = append(releaseInfos, ri)
		}
	}
	sort.SliceStable(releaseInfos, func(i, j int) bool {
		ri := releaseInfos[i]
		rj := releaseInfos[j]
		if ri.Origin != rj.Origin {
			return ri.Origin < rj.Origin
		}
		return ri.Date.Before(rj.Date)
	})
	return releaseInfos
}

func cmdLineParse() (string, error) {
	if len(os.Args) > 2 {
		return "", fmt.Errorf("could not parse %s", os.Args)
	}
	if len(os.Args) == 1 {
		return "-", nil
	}
	return os.Args[1], nil
}

func main() {
	filename, err := cmdLineParse()
	panicIf(err)
	releaseInfos := getAptmirrorsReleaseInfos()
	json, err := json.MarshalIndent(releaseInfos, "", "  ")
	panicIf(err)
	jsonstr := string(json) + "\n"
	if filename == "" || filename == "-" {
		fmt.Print(jsonstr)
	} else {
		err := ioutil.WriteFile(filename, []byte(jsonstr), 0644)
		panicIf(err)
	}
}
