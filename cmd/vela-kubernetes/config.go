// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Config represents the plugin configuration for Kubernetes config information.
type Config struct {
	// Kubernetes files or directories to apply
	Files []string
	// new container images from files to apply
	Images []string
	// Kubernetes resources to watch status of
	Statuses []string
	// total time allowed to watch Kubernetes resources
	Timeout time.Duration
}

// Validate verifies the Config is properly configured.
func (c *Config) Validate() error {
	logrus.Trace("validating config configuration")

	// verify files are provided
	if len(c.Files) == 0 {
		return fmt.Errorf("no config files provided")
	}

	// verify images are provided
	if len(c.Images) == 0 {
		return fmt.Errorf("no config images provided")
	}

	// verify statuses are provided
	if len(c.Statuses) == 0 {
		return fmt.Errorf("no config statuses provided")
	}

	// verify timeout is provided
	if c.Timeout <= 0 {
		return fmt.Errorf("no config timeout provided")
	}

	return nil
}
