// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
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
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
		Path:      "~/.kube/config",
	}

	s := &Status{
		Resources: []string{"resources"},
		Timeout:   5 * time.Minute,
		Watch:     true,
	}

	// nolint: gosec // testing purposes
	for _, resource := range s.Resources {
		want := exec.Command(
			_kubectl,
			fmt.Sprintf("--kubeconfig=%s", c.Path),
			fmt.Sprintf("--cluster=%s", c.Cluster),
			fmt.Sprintf("--context=%s", c.Context),
			fmt.Sprintf("--namespace=%s", c.Namespace),
			"rollout",
			"status",
			resource,
			fmt.Sprintf("--timeout=%s", s.Timeout),
			fmt.Sprintf("--watch=%t", s.Watch),
		)

		got := s.Command(c, resource)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

func TestKubernetes_Status_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "status",
		File:      "file",
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	s := &Status{
		Resources: []string{"resources"},
		Timeout:   5 * time.Minute,
		Watch:     true,
	}

	err := s.Exec(c)
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Status_Validate(t *testing.T) {
	// setup types
	s := &Status{
		Resources: []string{"resources"},
		Timeout:   5 * time.Minute,
		Watch:     true,
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
		Watch:   true,
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
		Watch:     true,
	}

	err := s.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
