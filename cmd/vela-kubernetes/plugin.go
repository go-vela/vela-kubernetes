// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	// ErrInvalidAction defines the error type when the
	// Action provided to the Plugin is unsupported.
	ErrInvalidAction = errors.New("invalid action provided")
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

	// execute action specific configuration
	switch p.Config.Action {
	case applyAction:
		// execute apply action
		return p.Apply.Exec(p.Config)
	case patchAction:
		// execute patch action
		return p.Patch.Exec(p.Config)
	case statusAction:
		// execute status action
		return p.Status.Exec(p.Config)
	default:
		return fmt.Errorf(
			"%w: %s (Valid actions: %s, %s, %s)",
			ErrInvalidAction,
			p.Config.Action,
			applyAction,
			patchAction,
			statusAction,
		)
	}
}

// Validate verifies the plugin is properly configured.
func (p *Plugin) Validate() error {
	logrus.Debug("validating plugin configuration")

	// validate config configuration
	err := p.Config.Validate()
	if err != nil {
		return err
	}

	// validate action specific configuration
	switch p.Config.Action {
	case applyAction:
		// validate apply configuration
		return p.Apply.Validate()
	case patchAction:
		// validate patch configuration
		return p.Patch.Validate()
	case statusAction:
		// validate status configuration
		return p.Status.Validate()
	default:
		return fmt.Errorf(
			"%w: %s (Valid actions: %s, %s, %s)",
			ErrInvalidAction,
			p.Config.Action,
			applyAction,
			patchAction,
			statusAction,
		)
	}
}
