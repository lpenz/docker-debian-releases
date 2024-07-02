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

	"../../common"
)

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
	token := os.Getenv("DOCKERHUB_TOKEN")
	err := dockerSetDescrs(token, image, description, fullDescription)
	if err != nil {
		panic(fmt.Sprintf("error setting descriptions: %s", err))
	}
}
