// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
	"time"
)

func TestKubernetes_Status_Validate(t *testing.T) {
	// setup types
	s := &Status{
		Resources: []string{"resources"},
		Timeout:   5 * time.Minute,
	}

	err := s.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Status_Validate_NoResources(t *testing.T) {
	// setup types
	s := &Status{
		Timeout: 5 * time.Minute,
	}

	err := s.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Status_Validate_NoTimeout(t *testing.T) {
	// setup types
	s := &Status{
		Resources: []string{"resources"},
	}

	err := s.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
