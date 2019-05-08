// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"../../common"
)

func cmdLineParse() (string, string, string, error) {
	if len(os.Args) > 4 || len(os.Args) < 2 {
		return "", "", "", fmt.Errorf("could not parse %s", os.Args)
	}
	output := "-"
	if len(os.Args) == 4 {
		output = os.Args[3]
	}
	return os.Args[1], os.Args[2], output, nil
}

func main() {
	jsonfilename, tmplfilename, outfilename, err := cmdLineParse()
	common.PanicIf(err)
	jsonString, err := ioutil.ReadFile(jsonfilename)
	common.PanicIf(err)
	var releaseInfos interface{}
	err = json.Unmarshal(jsonString, &releaseInfos)
	common.PanicIf(err)
	t, err := template.New("").Funcs(template.FuncMap{
		"stringsIndex": strings.Index,
		"ToLower":      strings.ToLower,
	}).ParseFiles(tmplfilename)
	common.PanicIf(err)
	if outfilename != "-" {
		outfile, err := os.Create(outfilename)
		common.PanicIf(err)
		err = t.ExecuteTemplate(outfile, tmplfilename, releaseInfos)
		common.PanicIf(err)
	} else {
		err = t.ExecuteTemplate(os.Stdout, tmplfilename, releaseInfos)
		common.PanicIf(err)
	}
}
