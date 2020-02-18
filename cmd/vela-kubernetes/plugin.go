// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/sirupsen/logrus"
)

// Plugin represents the configuration loaded for the plugin.
type Plugin struct {
	// apply arguments loaded for the plugin
	Apply *Apply
	// config arguments loaded for the plugin
	Config *Config
	// patch arguments loaded for the plugin
	Patch *Patch
	// status arguments loaded for the plugin
	Status *Status
}

// Exec formats and runs the commands for managing resources in Kubernetes.
func (p *Plugin) Exec() error {
	logrus.Debug("running plugin with provided configuration")

	// create kubectl configuration file for authentication
	err := p.Config.Write()
	if err != nil {
		return err
	}

	// output kubectl version for troubleshooting
	err = execCmd(versionCmd(p.Config))
	if err != nil {
		return err
	}

	// apply the Kubernetes resource files
	err = p.Apply.Exec(p.Config)
	if err != nil {
		return err
	}

	// watch the status for the specified resources
	err = p.Status.Exec(p.Config)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the plugin is properly configured.
func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate apply configuration
	err := p.Apply.Validate()
	if err != nil {
		return err
	}

	// validate config configuration
	err = p.Config.Validate()
	if err != nil {
		return err
	}

	// validate patch configuration
	err = p.Patch.Validate()
	if err != nil {
		return err
	}

	// validate status configuration
	err = p.Status.Validate()
	if err != nil {
		return err
	}

	return nil
}
