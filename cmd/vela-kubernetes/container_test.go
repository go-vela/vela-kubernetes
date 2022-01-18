// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
)

func TestKubernetes_Container_Validate(t *testing.T) {
	// setup types
	c := &Container{
		Name:  "container",
		Image: "alpine",
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Container_Validate_NoName(t *testing.T) {
	// setup types
	c := &Container{
		Image: "alpine",
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Container_Validate_NoImage(t *testing.T) {
	// setup types
	c := &Container{
		Name: "container",
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
