// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"../../common"
)

// Created with help from https://mholt.github.io/json-to-go/
type BranchInfo struct {
	Branch struct {
		RepositoryID int       `json:"repository_id"`
		Duration     int       `json:"duration"`
		Number       string    `json:"number"`
		StartedAt    time.Time `json:"started_at"`
		CommitID     int       `json:"commit_id"`
		PullRequest  bool      `json:"pull_request"`
		ID           int       `json:"id"`
		FinishedAt   time.Time `json:"finished_at"`
		Config       struct {
			Dist string `json:"dist"`
			Jobs struct {
				Include []struct {
					Addons struct {
						Apt struct {
							Packages []string `json:"packages"`
						} `json:"apt"`
					} `json:"addons,omitempty"`
					If      string      `json:"if"`
					Name    string      `json:"name"`
					Script  interface{} `json:"script"`
					Install interface{} `json:"install,omitempty"`
				} `json:"include"`
			} `json:"jobs"`
			Deploy []struct {
				True struct {
					AllBranches bool   `json:"all_branches"`
					Condition   string `json:"condition"`
				} `json:"true"`
				Provider string `json:"provider"`
				Script   string `json:"script"`
			} `json:"deploy"`
			Group        string   `json:"group"`
			BeforeDeploy []string `json:"before_deploy"`
			Result       string   `json:".result"`
			Language     string   `json:"language"`
			Os           string   `json:"os"`
		} `json:"config"`
		JobIds []int  `json:"job_ids"`
		State  string `json:"state"`
	} `json:"branch"`
	Commit struct {
		CommitterName  string      `json:"committer_name"`
		AuthorEmail    string      `json:"author_email"`
		AuthorName     string      `json:"author_name"`
		Sha            string      `json:"sha"`
		CommitterEmail string      `json:"committer_email"`
		Message        string      `json:"message"`
		Branch         string      `json:"branch"`
		CompareURL     string      `json:"compare_url"`
		CommittedAt    interface{} `json:"committed_at"`
		ID             int         `json:"id"`
	} `json:"commit"`
}

func getTravisBranchInfo(branch string) *BranchInfo {
	urlbase := "https://api.travis-ci.com/repos/lpenz/docker-debian-releases/branches"
	client := &http.Client{}
	url := fmt.Sprintf("%s/%s", urlbase, branch)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	common.PanicIf(err)
	req.Header.Add("Accept", "application/vnd.travis-ci.2.1+json")
	resp, err := client.Do(req)
	common.PanicIf(err)
	defer func() {
		err := resp.Body.Close()
		common.PanicIf(err)
	}()
	if resp.StatusCode == 404 {
		return nil
	}
	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("status: %d", resp.StatusCode))
	}
	branchInfo := BranchInfo{}
	err = json.NewDecoder(resp.Body).Decode(&branchInfo)
	common.PanicIf(err)
	return &branchInfo
}

func getTravisBranchInfos(releaseInfos *[]common.ReleaseInfo) *map[string]*BranchInfo {
	ret := map[string]*BranchInfo{}
	for _, release := range *releaseInfos {
		if strings.Index(release.Suite, "-") != -1 {
			continue
		}
		for _, a := range release.Architectures {
			if strings.Index(a, "-") != -1 {
				continue
			}
			suffixes := []string{"", "-minbase"}
			for _, suffix := range suffixes {
				branch := fmt.Sprintf("%s-%s-%s%s", strings.ToLower(release.Origin), release.Codename, a, suffix)
				branchInfo := getTravisBranchInfo(branch)
				if branchInfo != nil {
					ret[branch] = branchInfo
				}
			}
		}
	}
	return &ret
}

type buildInfo struct {
	ID         string    `json:"id"`
	State      string    `json:"state"`
	FinishedAt time.Time `json:"finished"`
	Duration   int       `json:"duration"`
}

func getBranchBuilds(branchInfo *map[string]*BranchInfo) *map[string]buildInfo {
	branchBuilds := map[string]buildInfo{}
	for branch, info := range *branchInfo {
		branchBuilds[branch] = buildInfo{
			ID:         fmt.Sprintf("%d", info.Branch.ID),
			State:      info.Branch.State,
			FinishedAt: info.Branch.FinishedAt,
			Duration:   info.Branch.Duration,
		}
	}
	return &branchBuilds
}

func readReleaseInfo(srcfilename string) *[]common.ReleaseInfo {
	srcfile, err := os.Open(srcfilename)
	common.PanicIf(err)
	defer func() {
		err := srcfile.Close()
		common.PanicIf(err)
	}()
	var releaseInfos map[string][]common.ReleaseInfo
	err = json.NewDecoder(srcfile).Decode(&releaseInfos)
	common.PanicIf(err)
	ret := releaseInfos["releaseInfos"]
	return &ret
}

func cmdLineParse() (string, string, error) {
	if len(os.Args) > 3 {
		return "", "", fmt.Errorf("could not parse %s", os.Args)
	}
	if len(os.Args) == 2 {
		return "", "-", nil
	}
	return os.Args[1], os.Args[2], nil
}

func main() {
	srcfilename, targetfilename, err := cmdLineParse()
	common.PanicIf(err)
	releaseInfos := readReleaseInfo(srcfilename)
	branchInfos := getTravisBranchInfos(releaseInfos)
	common.JsonOutput(targetfilename, "branchBuilds", getBranchBuilds(branchInfos))
}
