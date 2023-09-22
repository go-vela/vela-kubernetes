// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// Container represents the plugin configuration for patching a container.
//
// TODO: Enable patching of other container attributes?
// * environment variables
// * image pull policy
// * ports
// * resources
//
// This is possible if we made Container a map[string]string.
type Container struct {
	// name of the container to patch
	Name string
	// image of the container to patch
	Image string
}

// Validate verifies the Container is properly configured.
func (c *Container) Validate() error {
	logrus.Tracef("validating container configuration for %s", c.Name)

	// verify name is provided
	if len(c.Name) == 0 {
		return fmt.Errorf("no container name provided")
	}

	// verify image is provided
	if len(c.Image) == 0 {
		return fmt.Errorf("no container image provided")
	}

	return nil
}
