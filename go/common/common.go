// Copyright (C) 2019 Leandro Lisboa Penz <lpenz@lpenz.org>
// This file is subject to the terms and conditions defined in
// file 'LICENSE', which is part of this source code package.

package common

import (
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
