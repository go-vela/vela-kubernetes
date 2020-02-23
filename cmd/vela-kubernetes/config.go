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
	// action to perform against Kubernetes
	Action string
	// name of the configuration cluster from file
	Cluster string
	// name of the configuration context from file
	Context string
	// configuration file for communication with Kubernetes
	File string
	// name of the configuration namespace from file
	Namespace string
	// path to configuration file for communcation with Kubernetes
	Path string
}

// Validate verifies the Config is properly configured.
func (c *Config) Validate() error {
	logrus.Trace("validating config configuration")

	// verify action is provided
	if len(c.Action) == 0 {
		return fmt.Errorf("no config action provided")
	}

	// verify file is provided
	if len(c.File) == 0 {
		return fmt.Errorf("no config file provided")
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

	// check if a custom config path was provided
	if len(c.Path) == 0 {
		// create full path for .kube/config file
		c.Path = filepath.Join(home, ".kube/config")
	}

	logrus.Debugf("Creating kubectl configuration file %s", c.Path)

	// send Filesystem call to create directory path for .kube/config file
	err = a.Fs.MkdirAll(filepath.Dir(c.Path), 0777)
	if err != nil {
		return err
	}

	return a.WriteFile(c.Path, []byte(c.File), 0600)
}
