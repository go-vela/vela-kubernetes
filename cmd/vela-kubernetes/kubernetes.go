// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Kubernetes represents the plugin configuration for Kubernetes information.
type Kubernetes struct {
	// cluster configuration to use for interactions
	Config string
	// cluster context to use for interactions
	Context string
	// cluster namespace to use for interactions
	Namespace string
}

// Validate verifies the Kubernetes is properly configured.
func (k *Kubernetes) Validate() error {
	logrus.Trace("validating kubernetes configuration")

	// verify config is provided
	if len(k.Config) == 0 {
		return fmt.Errorf("no kubernetes config provided")
	}

	// verify context is provided
	if len(k.Context) == 0 {
		return fmt.Errorf("no kubernetes context provided")
	}

	// verify namespace is provided
	if len(k.Namespace) == 0 {
		return fmt.Errorf("no kubernetes namespace provided")
	}

	return nil
}
