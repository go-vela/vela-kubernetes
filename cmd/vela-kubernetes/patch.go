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

// Command formats and outputs the Patch command from
// the provided configuration to patch resources.
func (p *Patch) Command(c *Config, image string) *exec.Cmd {
	logrus.Tracef("creating kubectl patch command for %s from plugin configuration", image)

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

	// add flag for patch kubectl command
	flags = append(flags, "patch")

	// add flag for output
	flags = append(flags, "--output=json")

	return exec.Command(kubectlBin, flags...)
}

// Exec formats and runs the commands for patching
// the provided configuration to the resources.
func (p *Patch) Exec(c *Config) error {
	logrus.Debug("running patch with provided configuration")

	// iterate through all images to patch
	for _, image := range p.Images {
		// create the patch command for the image
		cmd := p.Command(c, image)

		// run the patch command for the image
		err := execCmd(cmd)
		if err != nil {
			return err
		}
	}

	return nil
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
