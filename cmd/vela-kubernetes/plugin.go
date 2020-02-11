// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Plugin represents the configuration loaded for the plugin.
type Plugin struct {
	// kubernetes arguments loaded for the plugin
	Kubernetes *Kubernetes
}

// Command formats and outputs the command necessary for
// kubectl to manage Kubernetes resources.
func (p *Plugin) Command() *exec.Cmd {
	logrus.Debug("creating kubectl command from plugin configuration")

	// variable to store flags for command
	var flags []string

	return exec.Command(kubectlBin, flags...)
}

// Exec formats and runs the commands for managing resources in Kubernetes.
func (p *Plugin) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	return nil
}

// Validate verifies the plugin is properly configured.
func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate Kubernetes configuration
	err := p.Kubernetes.Validate()
	if err != nil {
		return err
	}

	return nil
}
