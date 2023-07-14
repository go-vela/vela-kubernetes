// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	getter "github.com/hashicorp/go-getter/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

const (
	_kubectl  = "/bin/kubectl"
	_download = "https://storage.googleapis.com/kubernetes-release/release/%s/bin/%s/%s/kubectl"
)

func install(customVer, defaultVer string) error {
	logrus.Infof("custom kubectl version requested: %s", customVer)

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if the custom version matches the default version
	if strings.EqualFold(customVer, defaultVer) {
		// the kubectl versions match so no action is required
		return nil
	}

	logrus.Debugf("custom version does not match default: %s", defaultVer)
	// rename the old kubectl binary since we can't overwrite it for now
	//
	// https://github.com/hashicorp/go-getter/issues/219
	err := a.Rename(_kubectl, fmt.Sprintf("%s.default", _kubectl))
	if err != nil {
		return err
	}

	// create the download URL to install kubectl
	url := fmt.Sprintf(_download, customVer, runtime.GOOS, runtime.GOARCH)

	logrus.Infof("downloading kubectl version from: %s", url)
	// send the HTTP request to install kubectl
	_, err = getter.GetFile(context.Background(), _kubectl, url)
	if err != nil {
		return err
	}

	logrus.Debugf("changing ownership of file: %s", _kubectl)

	// ensure the kubectl binary is executable
	err = a.Chmod(_kubectl, 0700)
	if err != nil {
		return err
	}

	return nil
}
