// Copyright (C) 2020 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package main

import (
	"archive/tar"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"../../common"
)

type Manifest []struct {
	Config   string
	Layers   interface{}
	RepoTags interface{}
}

type Config struct {
	Architecture    string        `json:"architecture"`
	Comment         string        `json:"comment"`
	Config          interface{}   `json:"config"`
	ContainerConfig interface{}   `json:"container_config"`
	Created         interface{}   `json:"created"`
	DockerVersion   interface{}   `json:"docker_version"`
	History         []interface{} `json:"history"`
	Os              string        `json:"os"`
	Rootfs          interface{}   `json:"rootfs"`
}

func tarGetContents(tarfilename string, filename string) []byte {
	tarfile, err := os.Open(tarfilename)
	common.PanicIf(err)
	defer tarfile.Close()
	tarReader := tar.NewReader(tarfile)
	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		common.PanicIf(err)
		if hdr.Name != filename {
			continue
		}
		jsonString, err := ioutil.ReadAll(tarReader)
		common.PanicIf(err)
		return jsonString
	}
	return nil
}

func tarGetManifest(tarfilename string, filename string) Manifest {
	jsonString := tarGetContents(tarfilename, filename)
	var tmp Manifest
	err := json.Unmarshal(jsonString, &tmp)
	common.PanicIf(err)
	return tmp
}

func tarGetConfig(tarfilename string, filename string) Config {
	jsonString := tarGetContents(tarfilename, filename)
	var tmp Config
	err := json.Unmarshal(jsonString, &tmp)
	common.PanicIf(err)
	return tmp
}

func tarCopyReplace(tarinfile string, taroutfile string, filename string, contents []byte) {
	tarout, err := os.Create(taroutfile)
	common.PanicIf(err)
	defer tarout.Close()
	tarWriter := tar.NewWriter(tarout)

	tarin, err := os.Open(tarinfile)
	common.PanicIf(err)
	defer tarin.Close()
	tarReader := tar.NewReader(tarin)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		common.PanicIf(err)
		if hdr.Name != filename {
			// Just copy over
			err = tarWriter.WriteHeader(hdr)
			common.PanicIf(err)
			io.Copy(tarWriter, tarReader)
			common.PanicIf(err)
		} else {
			newHdr := hdr
			newHdr.Size = int64(len(contents))
			err = tarWriter.WriteHeader(newHdr)
			common.PanicIf(err)
			_, err := tarWriter.Write(contents)
			common.PanicIf(err)
		}
	}
	tarWriter.Flush()
}

func cmdLineParse() (string, string, string, error) {
	if len(os.Args) != 4 {
		return "", "", "", fmt.Errorf("could not parse %s", os.Args)
	}
	return os.Args[1], os.Args[2], os.Args[3], nil
}

func main() {
	arch, infilename, outfilename, err := cmdLineParse()
	common.PanicIf(err)
	manifest := tarGetManifest(infilename, "manifest.json")
	configfilename := manifest[0].Config
	config := tarGetConfig(infilename, configfilename)
	config.Architecture = arch
	jsonstr, err := json.Marshal(config)
	common.PanicIf(err)
	tarCopyReplace(infilename, outfilename, configfilename, jsonstr)
}
