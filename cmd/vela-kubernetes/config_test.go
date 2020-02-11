// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
	"time"
)

func TestKubernetes_Config_Validate(t *testing.T) {
	// setup types
	c := &Config{
		Files:    []string{"files"},
		Images:   []string{"images"},
		Statuses: []string{"statuses"},
		Timeout:  5 * time.Minute,
	}

	err := c.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Config_Validate_NoFiles(t *testing.T) {
	// setup types
	c := &Config{
		Images:   []string{"images"},
		Statuses: []string{"statuses"},
		Timeout:  5 * time.Minute,
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Config_Validate_NoImages(t *testing.T) {
	// setup types
	c := &Config{
		Files:    []string{"files"},
		Statuses: []string{"statuses"},
		Timeout:  5 * time.Minute,
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Config_Validate_NoStatuses(t *testing.T) {
	// setup types
	c := &Config{
		Files:   []string{"files"},
		Images:  []string{"images"},
		Timeout: 5 * time.Minute,
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Config_Validate_NoTimeout(t *testing.T) {
	// setup types
	c := &Config{
		Files:    []string{"files"},
		Images:   []string{"images"},
		Statuses: []string{"statuses"},
	}

	err := c.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
