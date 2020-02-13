// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// Apply represents the plugin configuration for Apply config information.
type Apply struct {
	// Kubernetes files or directories to apply
	Files []string
	// sets the output for the apply command
	Output string
}

// Command formats and outputs the Apply command from
// the provided configuration to apply to resources.
func (a *Apply) Command(c *Config, file string) *exec.Cmd {
	logrus.Trace("creating kubectl apply command from plugin configuration")

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
	flags = append(flags, "apply")

	// check if file is provided
	if len(file) > 0 {
		// add flag for file from provided apply file
		flags = append(flags, fmt.Sprintf("--filename=%s", file))
	}

	// check if apply output is provided
	if len(a.Output) > 0 {
		// add flag for output from provided apply output
		flags = append(flags, fmt.Sprintf("--output=%s", a.Output))
	}

	return exec.Command(kubectlBin, flags...)
}

// Exec formats and runs the commands for applying
// the provided configuration to the resources.
func (a *Apply) Exec(c *Config) error {
	logrus.Debug("running apply with provided configuration")

	// iterate through all files to apply
	for _, file := range a.Files {
		// create the apply command for the file
		cmd := a.Command(c, file)

		// run the apply command for the file
		err := execCmd(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate verifies the Apply is properly configured.
func (a *Apply) Validate() error {
	logrus.Trace("validating apply configuration")

	// verify files are provided
	if len(a.Files) == 0 {
		return fmt.Errorf("no apply files provided")
	}

	return nil
}
