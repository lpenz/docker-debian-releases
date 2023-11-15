// Copyright (C) 2020 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"../../common"
)

func gitAllBranches() []string {
	out, err := exec.Command("git", "for-each-ref", "--format=\"%(refname:short)\"").Output()
	if err != nil {
		return []string{}
	}
	ret := []string{}
	for _, rev0 := range strings.Split(string(out), "\n") {
		rev := strings.Replace(rev0, "origin/", "", -1)
		if len(rev) < 2 {
			continue
		}
		rev = rev[1 : len(rev)-1]
		ret = append(ret, rev)
	}
	return ret
}

func cmdLineParse() (string, error) {
	if len(os.Args) > 2 {
		return "", fmt.Errorf("could not parse %s", os.Args)
	}
	return os.Args[1], nil
}

func main() {
	srcfilename, err := cmdLineParse()
	common.PanicIf(err)
	releaseInfos := common.ReadReleaseInfo(srcfilename)
	allBranches := map[string]bool{}
	for _, ri := range *releaseInfos {
		for _, branchname := range ri.BranchNames() {
			if branchname == "" {
				continue
			}
			allBranches[branchname] = true
		}
	}
	// Add some innocent branches to list:
	allBranches["HEAD"] = true
	allBranches["master"] = true
	allBranches["main"] = true
	allBranches["devel"] = true
	allBranches["gh-pages"] = true
	for _, b := range gitAllBranches() {
		_, found := allBranches[b]
		if !found {
			fmt.Fprintf(os.Stderr, "Warning: found spurious branch %s\n", b)
		}
	}
}
