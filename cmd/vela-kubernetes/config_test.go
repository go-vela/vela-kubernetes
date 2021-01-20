// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestKubernetes_Config_Validate(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Config_Validate_NoAction(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Config_Validate_NoFile(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Config_Write(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
	}

	err := c.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}

func TestKubernetes_Config_Write_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewReadOnlyFs(afero.NewMemMapFs())

	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
	}

	err := c.Write()
	if err == nil {
		t.Errorf("Write should have returned err")
	}
}

func TestKubernetes_Config_Write_NoFile(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	err := c.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}
