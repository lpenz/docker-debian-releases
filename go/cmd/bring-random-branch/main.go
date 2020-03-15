// Copyright (C) 2020 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	"../../common"
)

func gitRevParse(rev string) string {
	out, err := exec.Command("git", "rev-parse", rev).Output()
	if err != nil {
		return ""
	}
	return strings.Split(string(out), "\n")[0]
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
	head := gitRevParse("HEAD")
	if head == "" {
		panic("invalid HEAD")
	}
	releaseInfos := common.ReadReleaseInfo(srcfilename)
	options := []common.ReleaseInfo{}
	foundNew := false
	for _, ri := range *releaseInfos {
		for _, branchname := range ri.BranchNames() {
			if branchname == "" {
				continue
			}
			rev := gitRevParse(fmt.Sprintf("origin/%s", branchname))
			if rev == "" {
				// Found brand new release, process it only
				fmt.Printf("New release found! %s\n", branchname)
				foundNew = true
				options = []common.ReleaseInfo{ri}
				break
			}
			if rev != head {
				options = append(options, ri)
				break
			}
		}
		if foundNew {
			break
		}
	}
	if len(options) == 0 {
		return
	}
	rand.Seed(time.Now().Unix())
	chosen := options[rand.Intn(len(options))]
	for _, ri := range *releaseInfos {
		if ri.Origin != ri.Origin || ri.Codename != chosen.Codename {
			continue
		}
		for _, branchname := range ri.BranchNames() {
			if branchname == "" {
				continue
			}
			cmdstr := []string{"git", "push", "-f", "origin", fmt.Sprintf("HEAD:%s", branchname)}
			fmt.Printf("+ %s\n", strings.Join(cmdstr, " "))
			cmd := exec.Command(cmdstr[0], cmdstr[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			common.PanicIf(err)
		}
	}
}
