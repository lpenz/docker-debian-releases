// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"../../common"
)

type jsonsFlags []string

func (i *jsonsFlags) String() string {
	return "my string representation"
}

func (i *jsonsFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var myJsonsFlags jsonsFlags

func cmdLineParse() ([]string, string, string, error) {
	flag.Var(&myJsonsFlags, "json", "json to import into scope")
	flag.Parse()
	if flag.NFlag() > 3 || flag.NFlag() < 1 {
		empty := []string{}
		return empty, "", "", fmt.Errorf("could not parse %s", os.Args)
	}
	output := "-"
	if flag.NFlag() == 1 {
		output = flag.Arg(1)
	}
	return myJsonsFlags, flag.Arg(0), output, nil
}

func main() {
	jsonfilenames, tmplfilename, outfilename, err := cmdLineParse()
	common.PanicIf(err)
	info := map[string]interface{}{}
	for _, jsonfilename := range jsonfilenames {
		jsonString, err := ioutil.ReadFile(jsonfilename)
		common.PanicIf(err)
		var tmp map[string]interface{}
		err = json.Unmarshal(jsonString, &tmp)
		for k, v := range tmp {
			info[k] = v
		}
		common.PanicIf(err)
	}
	t, err := template.New("").Funcs(template.FuncMap{
		"stringsIndex": strings.Index,
		"ToLower":      strings.ToLower,
	}).ParseFiles(tmplfilename)
	common.PanicIf(err)
	if outfilename != "-" {
		outfile, err := os.Create(outfilename)
		common.PanicIf(err)
		err = t.ExecuteTemplate(outfile, tmplfilename, info)
		common.PanicIf(err)
	} else {
		err = t.ExecuteTemplate(os.Stdout, tmplfilename, info)
		common.PanicIf(err)
	}
}
