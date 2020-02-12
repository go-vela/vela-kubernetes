// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Patch represents the plugin configuration for Patch config information.
type Patch struct {
	// container images from files to patch
	Images []string
}

// Command formats and outputs the Patch command from the
// provided configuration to patch resources.
func (p *Patch) Command(c *Config) *exec.Cmd {
	logrus.Trace("creating kubectl patch command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// check if config namespace is provided
	if len(c.Namespace) > 0 {
		// add flag for namespace from provided config namespace
		flags = append(flags, fmt.Sprintf("--namespace=%s", c.Namespace))
	}

	// check if config context is provided
	if len(c.Context) > 0 {
		// add flag for context from provided config context
		flags = append(flags, fmt.Sprintf("--context=%s", c.Context))
	}

	// add flag for apply kubectl command
	flags = append(flags, "patch")

	// add flag for output
	flags = append(flags, "--output=json")

	return exec.Command(kubectlBin, flags...)
}

// Validate verifies the Patch is properly configured.
func (p *Patch) Validate() error {
	logrus.Trace("validating patch configuration")

	// verify images are provided
	if len(p.Images) == 0 {
		return fmt.Errorf("no patch images provided")
	}

	return nil
}
