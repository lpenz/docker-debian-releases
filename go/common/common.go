// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func PanicIf(err error, transforms ...func(error) error) {
	if err == nil {
		return
	}
	for _, transform := range transforms {
		err = transform(err)
	}
	panic(err)
}

type ReleaseInfo struct {
	Origin        string
	Label         string
	Suite         string
	Version       string
	Date          time.Time
	Codename      string
	URL           []string
	Architectures []string
}

func JsonOutput(filename string, title string, data interface{}) {
	info := map[string]interface{}{
		title: data,
	}
	json, err := json.MarshalIndent(info, "", "  ")
	PanicIf(err)
	jsonstr := string(json) + "\n"
	if filename == "" || filename == "-" {
		fmt.Print(jsonstr)
	} else {
		err := ioutil.WriteFile(filename, []byte(jsonstr), 0644)
		PanicIf(err)
	}
}

func ReadReleaseInfo(srcfilename string) *[]ReleaseInfo {
	srcfile, err := os.Open(srcfilename)
	PanicIf(err)
	defer func() {
		err := srcfile.Close()
		PanicIf(err)
	}()
	var releaseInfos map[string][]ReleaseInfo
	err = json.NewDecoder(srcfile).Decode(&releaseInfos)
	PanicIf(err)
	ret := releaseInfos["releaseInfos"]
	return &ret
}
