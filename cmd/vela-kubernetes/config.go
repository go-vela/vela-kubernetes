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

// Config represents the plugin configuration for Kubernetes information.
type Config struct {
	// cluster configuration file to interact with Kubernetes
	File string
	// name of the configuration cluster to use for interactions with Kubernetes
	Cluster string
	// name of the configuration context to use for interactions with Kubernetes
	Context string
	// name of the configuration namespace to use for interactions with Kubernetes
	Namespace string
}

// Validate verifies the Config is properly configured.
func (c *Config) Validate() error {
	logrus.Trace("validating config configuration")

	// verify file is provided
	if len(c.File) == 0 {
		return fmt.Errorf("no config file provided")
	}

	// verify cluster is provided
	if len(c.Cluster) == 0 {
		return fmt.Errorf("no config cluster provided")
	}

	// verify context is provided
	if len(c.Context) == 0 {
		return fmt.Errorf("no config context provided")
	}

	// verify namespace is provided
	if len(c.Namespace) == 0 {
		return fmt.Errorf("no config namespace provided")
	}

	return nil
}

// Write creates a .kube/config file in the home directory of the current user.
func (c *Config) Write() error {
	logrus.Trace("writing Kubernetes configuration file")

	// use custom filesystem which enables us to test
	a := &afero.Afero{
		Fs: appFS,
	}

	// check if config file content is provided
	if len(c.File) == 0 {
		return nil
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

	return a.WriteFile(path, []byte(c.File), 0600)
}
