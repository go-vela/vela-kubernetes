// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

const statusAction = "status"

// Status represents the plugin configuration for Status config information.
type Status struct {
	// Kubernetes resources to watch status of
	Resources []string
	// total time allowed to watch Kubernetes resources
	Timeout time.Duration
	// enables watching the status until the rollout completes
	Watch bool
}

// Command formats and outputs the Status command from the
// provided configuration to watch the status on resources.
func (s *Status) Command(c *Config, resource string) *exec.Cmd {
	logrus.Trace("creating kubectl status command from plugin configuration")

	// variable to store flags for command
	var flags []string

	// check if config namespace is provided
	if len(c.Namespace) > 0 {
		// add flag for namespace from provided config namespace
		flags = append(flags, fmt.Sprintf("--namespace=%s", c.Namespace))
	}

	// check if config cluster is provided
	if len(c.Cluster) > 0 {
		// add flag for cluster from provided config cluster
		flags = append(flags, fmt.Sprintf("--cluster=%s", c.Cluster))
	}

	// check if config context is provided
	if len(c.Context) > 0 {
		// add flag for context from provided config context
		flags = append(flags, fmt.Sprintf("--context=%s", c.Context))
	}

	// add flag for status kubectl command
	flags = append(flags, "rollout", "status")

	// check if resource is provided
	if len(resource) > 0 {
		// add flag for resource from provided status resource
		flags = append(flags, resource)
	}

	// check if status timeout is provided
	if s.Timeout > 0 {
		// add flag for timeout from provided status timeout
		flags = append(flags, fmt.Sprintf("--timeout=%v", s.Timeout))
	}

	// add flag for watching status of rollout until it finishes
	flags = append(flags, fmt.Sprintf("--watch=%v", s.Watch))

	return exec.Command(kubectlBin, flags...)
}

// Exec formats and runs the commands for watching the
// status on resources from the provided configuration.
func (s *Status) Exec(c *Config) error {
	logrus.Debug("running status with provided configuration")

	// iterate through all resources to watch status
	for _, resource := range s.Resources {
		// create the status command for the resource
		cmd := s.Command(c, resource)

		// run the status command for the resource
		err := execCmd(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// Validate verifies the Status is properly configured.
func (s *Status) Validate() error {
	logrus.Trace("validating status configuration")

	// verify resources are provided
	if len(s.Resources) == 0 {
		return fmt.Errorf("no status resources provided")
	}

	// verify timeout is provided
	if s.Timeout <= 0 {
		return fmt.Errorf("no status timeout provided")
	}

	return nil
}
