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
type WorkflowRuns struct {
	TotalCount   int `json:"total_count"`
	WorkflowRuns []struct {
		ID               int64         `json:"id"`
		Name             string        `json:"name"`
		NodeID           string        `json:"node_id"`
		HeadBranch       string        `json:"head_branch"`
		HeadSha          string        `json:"head_sha"`
		RunNumber        int64         `json:"run_number"`
		Event            string        `json:"event"`
		Status           string        `json:"status"`
		Conclusion       string        `json:"conclusion"`
		WorkflowID       int64         `json:"workflow_id"`
		CheckSuiteID     int64         `json:"check_suite_id"`
		CheckSuiteNodeID string        `json:"check_suite_node_id"`
		URL              string        `json:"url"`
		HTMLURL          string        `json:"html_url"`
		PullRequests     []interface{} `json:"pull_requests"`
		CreatedAt        time.Time     `json:"created_at"`
		UpdatedAt        time.Time     `json:"updated_at"`
		JobsURL          string        `json:"jobs_url"`
		LogsURL          string        `json:"logs_url"`
		CheckSuiteURL    string        `json:"check_suite_url"`
		ArtifactsURL     string        `json:"artifacts_url"`
		CancelURL        string        `json:"cancel_url"`
		RerunURL         string        `json:"rerun_url"`
		WorkflowURL      string        `json:"workflow_url"`
		HeadCommit       struct {
			ID        string    `json:"id"`
			TreeID    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"head_commit"`
		Repository struct {
			ID       int64  `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int64  `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ForksURL         string `json:"forks_url"`
			KeysURL          string `json:"keys_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			TeamsURL         string `json:"teams_url"`
			HooksURL         string `json:"hooks_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			EventsURL        string `json:"events_url"`
			AssigneesURL     string `json:"assignees_url"`
			BranchesURL      string `json:"branches_url"`
			TagsURL          string `json:"tags_url"`
			BlobsURL         string `json:"blobs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitRefsURL       string `json:"git_refs_url"`
			TreesURL         string `json:"trees_url"`
			StatusesURL      string `json:"statuses_url"`
			LanguagesURL     string `json:"languages_url"`
			StargazersURL    string `json:"stargazers_url"`
			ContributorsURL  string `json:"contributors_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			CommitsURL       string `json:"commits_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			CommentsURL      string `json:"comments_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			ContentsURL      string `json:"contents_url"`
			CompareURL       string `json:"compare_url"`
			MergesURL        string `json:"merges_url"`
			ArchiveURL       string `json:"archive_url"`
			DownloadsURL     string `json:"downloads_url"`
			IssuesURL        string `json:"issues_url"`
			PullsURL         string `json:"pulls_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			LabelsURL        string `json:"labels_url"`
			ReleasesURL      string `json:"releases_url"`
			DeploymentsURL   string `json:"deployments_url"`
		} `json:"repository"`
		HeadRepository struct {
			ID       int64  `json:"id"`
			NodeID   string `json:"node_id"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			Private  bool   `json:"private"`
			Owner    struct {
				Login             string `json:"login"`
				ID                int64  `json:"id"`
				NodeID            string `json:"node_id"`
				AvatarURL         string `json:"avatar_url"`
				GravatarID        string `json:"gravatar_id"`
				URL               string `json:"url"`
				HTMLURL           string `json:"html_url"`
				FollowersURL      string `json:"followers_url"`
				FollowingURL      string `json:"following_url"`
				GistsURL          string `json:"gists_url"`
				StarredURL        string `json:"starred_url"`
				SubscriptionsURL  string `json:"subscriptions_url"`
				OrganizationsURL  string `json:"organizations_url"`
				ReposURL          string `json:"repos_url"`
				EventsURL         string `json:"events_url"`
				ReceivedEventsURL string `json:"received_events_url"`
				Type              string `json:"type"`
				SiteAdmin         bool   `json:"site_admin"`
			} `json:"owner"`
			HTMLURL          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			URL              string `json:"url"`
			ForksURL         string `json:"forks_url"`
			KeysURL          string `json:"keys_url"`
			CollaboratorsURL string `json:"collaborators_url"`
			TeamsURL         string `json:"teams_url"`
			HooksURL         string `json:"hooks_url"`
			IssueEventsURL   string `json:"issue_events_url"`
			EventsURL        string `json:"events_url"`
			AssigneesURL     string `json:"assignees_url"`
			BranchesURL      string `json:"branches_url"`
			TagsURL          string `json:"tags_url"`
			BlobsURL         string `json:"blobs_url"`
			GitTagsURL       string `json:"git_tags_url"`
			GitRefsURL       string `json:"git_refs_url"`
			TreesURL         string `json:"trees_url"`
			StatusesURL      string `json:"statuses_url"`
			LanguagesURL     string `json:"languages_url"`
			StargazersURL    string `json:"stargazers_url"`
			ContributorsURL  string `json:"contributors_url"`
			SubscribersURL   string `json:"subscribers_url"`
			SubscriptionURL  string `json:"subscription_url"`
			CommitsURL       string `json:"commits_url"`
			GitCommitsURL    string `json:"git_commits_url"`
			CommentsURL      string `json:"comments_url"`
			IssueCommentURL  string `json:"issue_comment_url"`
			ContentsURL      string `json:"contents_url"`
			CompareURL       string `json:"compare_url"`
			MergesURL        string `json:"merges_url"`
			ArchiveURL       string `json:"archive_url"`
			DownloadsURL     string `json:"downloads_url"`
			IssuesURL        string `json:"issues_url"`
			PullsURL         string `json:"pulls_url"`
			MilestonesURL    string `json:"milestones_url"`
			NotificationsURL string `json:"notifications_url"`
			LabelsURL        string `json:"labels_url"`
			ReleasesURL      string `json:"releases_url"`
			DeploymentsURL   string `json:"deployments_url"`
		} `json:"head_repository"`
	} `json:"workflow_runs"`
}

func getGithubWorkflowRun(page int) *WorkflowRuns {
	urlbase := "https://api.github.com/repos/lpenz/docker-debian-releases/actions/runs?per_page=100"
	client := &http.Client{}
	url := fmt.Sprintf("%s&page=%d", urlbase, page)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	common.PanicIf(err)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
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
		panic(fmt.Sprintf("page %d got status %d", page, resp.StatusCode))
	}
	workflowRuns := WorkflowRuns{}
	err = json.NewDecoder(resp.Body).Decode(&workflowRuns)
	common.PanicIf(err)
	return &workflowRuns
}

func getGithubWorkflowRuns() *WorkflowRuns {
	workflowRuns := getGithubWorkflowRun(1)
	page := 1
	for {
		if len(workflowRuns.WorkflowRuns) >= workflowRuns.TotalCount {
			break
		}
		newrun := getGithubWorkflowRun(page)
		page += 1
		workflowRuns.WorkflowRuns = append(workflowRuns.WorkflowRuns, newrun.WorkflowRuns...)
	}
	return workflowRuns
}

type branchWorkflowRun struct {
	Name      string
	State     string    `json:"state"`
	UpdatedAt time.Time `json:"finished"`
}

func getBranchWorkflowRuns(workflowRuns *WorkflowRuns) *map[string]branchWorkflowRun {
	branchBuilds := map[string]branchWorkflowRun{}
	for _, info := range workflowRuns.WorkflowRuns {
		setit := true
		if val, ok := branchBuilds[info.HeadBranch]; ok {
			if val.UpdatedAt.After(info.UpdatedAt) {
				setit = false
			}
		}
		if setit {
			branchBuilds[info.HeadBranch] = branchWorkflowRun{
				Name:      info.HeadBranch,
				State:     info.Conclusion,
				UpdatedAt: info.UpdatedAt,
			}
		}
	}
	return &branchBuilds
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
	targetfilename, err := cmdLineParse()
	common.PanicIf(err)
	githubWorkflowRuns := getGithubWorkflowRuns()
	branchWorkflowRuns := getBranchWorkflowRuns(githubWorkflowRuns)
	common.JsonOutput(targetfilename, "branchWorkflowRuns", branchWorkflowRuns)
}
