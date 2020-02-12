// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// Status represents the plugin configuration for Status config information.
type Status struct {
	// Kubernetes resources to watch status of
	Resources []string
	// total time allowed to watch Kubernetes resources
	Timeout time.Duration
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
