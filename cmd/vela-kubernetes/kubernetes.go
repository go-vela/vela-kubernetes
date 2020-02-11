// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/spf13/afero"

	"github.com/sirupsen/logrus"
)

var appFS = afero.NewOsFs()

// Kubernetes represents the plugin configuration for Kubernetes information.
type Kubernetes struct {
	// cluster configuration to use for interactions
	Config string
	// cluster context to use for interactions
	Context string
	// cluster namespace to use for interactions
	Namespace string
}

// Write creates a .kube/config file in the home directory of the current user.
func (k *Kubernetes) Write() error {
	logrus.Trace("writing kubernetes configuration file")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// set default home directory for root user
	home := "/root"

	// capture current user running commands
	u, err := user.Current()
	if err == nil {
		// set home directory to current user
		home = u.HomeDir
	}

	// create full path for .kube/config file
	path := filepath.Join(home, ".kube/config")

	return a.WriteFile(path, []byte(k.Config), 0600)
}

// Validate verifies the Kubernetes is properly configured.
func (k *Kubernetes) Validate() error {
	logrus.Trace("validating kubernetes configuration")

	// verify config is provided
	if len(k.Config) == 0 {
		return fmt.Errorf("no kubernetes config provided")
	}

	// verify context is provided
	if len(k.Context) == 0 {
		return fmt.Errorf("no kubernetes context provided")
	}

	// verify namespace is provided
	if len(k.Namespace) == 0 {
		return fmt.Errorf("no kubernetes namespace provided")
	}

	return nil
}
