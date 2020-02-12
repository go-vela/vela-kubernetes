// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
	"time"
)

func TestKubernetes_Status_Command(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Context:   "context",
		Namespace: "namespace",
	}

	s := &Status{
		Resources: []string{"resources"},
		Timeout:   5 * time.Minute,
	}

	for _, resource := range s.Resources {
		want := exec.Command(
			kubectlBin,
			fmt.Sprintf("--namespace=%s", c.Namespace),
			fmt.Sprintf("--context=%s", c.Context),
			"rollout",
			"status",
			resource,
			fmt.Sprintf("--timeout=%v", s.Timeout),
			"--watch=true",
		)

		got := s.Command(c, resource)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

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
