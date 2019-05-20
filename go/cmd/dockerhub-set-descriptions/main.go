// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"../../common"
)

func getToken(username string, password string) (string, error) {
	jsonReq, err := json.MarshalIndent(map[string]string{"username": username, "password": password}, "", "")
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", "https://hub.docker.com/v2/users/login/",
		strings.NewReader(string(jsonReq)))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		err := resp.Body.Close()
		common.PanicIf(err)
	}()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status: %d", resp.StatusCode)
	}
	jsonRespStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var respMap map[string]string
	err = json.Unmarshal(jsonRespStr, &respMap)
	if err != nil {
		return "", err
	}
	token, ok := respMap["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response")
	}
	return token, nil
}

func dockerSetDescrs(token string, image string, description string, fullDescription string) error {
	reqData := map[string]string{
		"registry":         "registry-1.docker.io",
		"description":      description,
		"full_description": fullDescription,
	}
	jsonReq, err := json.MarshalIndent(reqData, "", "")
	if err != nil {
		return err
	}
	baseUrl := "https://cloud.docker.com/v2/repositories"
	url := fmt.Sprintf("%s/%s/", baseUrl, image)
	req, err := http.NewRequest("PATCH", url, strings.NewReader(string(jsonReq)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("JWT %s", token))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err := resp.Body.Close()
		common.PanicIf(err)
	}()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status: %d", resp.StatusCode)
	}
	return nil
}

func cmdLineParse() (string, string, string) {
	if len(os.Args) != 4 {
		panic(fmt.Sprintf("could not parse %s", os.Args))
	}
	return os.Args[1], os.Args[2], os.Args[3]
}

func main() {
	image, description, fullDescription := cmdLineParse()
	username := os.Getenv("DOCKERHUB_USERNAME")
	if username == "" {
		panic("DOCKERHUB_USERNAME not set")
	}
	password := os.Getenv("DOCKERHUB_PASSWORD")
	if password == "" {
		panic("DOCKERHUB_USERNAME not set")
	}
	token, err := getToken(username, password)
	if err != nil {
		panic(fmt.Sprintf("error getting token: %s", err))
	}
	err = dockerSetDescrs(token, image, description, fullDescription)
	if err != nil {
		panic(fmt.Sprintf("error setting descriptions: %s", err))
	}
}
