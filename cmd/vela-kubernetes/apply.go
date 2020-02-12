// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Apply represents the plugin configuration for Apply config information.
type Apply struct {
	// Kubernetes files or directories to apply
	Files []string
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
