// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"

	"../../common"
)

func getLinksFromURL(url string) ([]string, error) {
	links := []string{}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		common.PanicIf(err)
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

func getReleaseInfo(url string) (common.ReleaseInfo, error) {
	ri := common.ReleaseInfo{}
	resp, err := http.Get(url)
	if err != nil {
		return ri, err
	}
	defer func() {
		err := resp.Body.Close()
		common.PanicIf(err)
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
			if ri.Origin == "Raspberry Pi Foundation" {
				ri.Origin = "rpios"
			}
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
			for _, a := range strings.Split(kv[2], " ") {
				if a != "" && a != "all" {
					ri.Architectures = append(ri.Architectures, a)
				}
			}
		}
	}
	if reflect.DeepEqual(ri, common.ReleaseInfo{}) {
		return ri, fmt.Errorf("could not parse lines in %s", url)
	}
	return ri, err
}

func hasUrl(ri *common.ReleaseInfo, url string) bool {
	for _, u := range ri.URL {
		if u == url {
			return true
		}
	}
	return false
}

func getAptmirrorsReleaseInfos() []*common.ReleaseInfo {
	mirrors := []string{
		"http://archive.debian.org/debian",
		"http://deb.debian.org/debian",
		"http://ports.ubuntu.com/ubuntu-ports",
		"http://archive.ubuntu.com/ubuntu",
		"http://old-releases.ubuntu.com/ubuntu",
		"http://archive.raspbian.org/raspbian",
		"http://raspbian.raspberrypi.org/raspbian",
		"http://archive.raspberrypi.org/debian",
		"http://deb.devuan.org/merged",
		"http://deb.devuan.org/devuan",
	}
	releaseInfos := []*common.ReleaseInfo{}
	riMap := map[string]*common.ReleaseInfo{}
	for _, m := range mirrors {
		url := m + "/dists" // /${dist}/main/binary-${arch}/Release
		links, err := getLinksFromURL(url)
		common.PanicIf(err)
		for _, link := range links {
			ri, err := getReleaseInfo(link + "/Release")
			if err != nil || ri.Codename == "None" || ri.Label == "None" {
				continue
			}
			riKey := fmt.Sprintf("%q", ri)
			riCurr, present := riMap[riKey]
			if present {
				if !hasUrl(riCurr, url) {
					riCurr.URL = append(riCurr.URL, url)
				}
			} else {
				ri.URL = append(ri.URL, url)
				releaseInfos = append(releaseInfos, &ri)
				riMap[riKey] = &ri
			}
		}
	}
	now := time.Now()
	sort.SliceStable(releaseInfos, func(i, j int) bool {
		ri := releaseInfos[i]
		rj := releaseInfos[j]
		if ri.Origin != rj.Origin {
			return ri.Origin < rj.Origin
		}
		// Ignore dates if both today
		riToday := now.Sub(ri.Date).Hours() < 24
		rjToday := now.Sub(ri.Date).Hours() < 24
		if !riToday || !rjToday {
			if ri.Date != rj.Date {
				return ri.Date.Before(rj.Date)
			}
		}
		if ri.Version != rj.Version {
			// Empty versions evaluated as z
			riV := ri.Version
			rjV := ri.Version
			if riV == "" {
				riV = "z"
			}
			if rjV == "" {
				rjV = "z"
			}
			return riV < rjV
		}
		if ri.Date != rj.Date {
			return ri.Date.Before(rj.Date)
		}
		return ri.Codename < rj.Codename
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
	common.PanicIf(err)
	common.JsonOutput(filename, "releaseInfos", getAptmirrorsReleaseInfos())
}
