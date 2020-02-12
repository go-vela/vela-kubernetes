// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Patch represents the plugin configuration for Patch config information.
type Patch struct {
	// container images from files to patch
	Images []string
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
