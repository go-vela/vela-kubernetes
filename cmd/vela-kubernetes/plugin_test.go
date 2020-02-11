// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os/exec"
	"reflect"
	"testing"
)

func TestKubernetes_Plugin_Command(t *testing.T) {
	// setup types
	p := &Plugin{}

	want := exec.Command(
		kubectlBin,
	)

	// run test
	got := p.Command()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Command is %v, want %v", got, want)
	}
}

func TestKubernetes_Plugin_Exec(t *testing.T) {
	// setup types
	p := &Plugin{}

	err := p.Exec()
	if err != nil {
		t.Errorf("Exec returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate(t *testing.T) {
	// setup types
	p := &Plugin{
		Kubernetes: &Kubernetes{
			Config:    "config",
			Context:   "context",
			Namespace: "namespace",
		},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate_NoKubernetes(t *testing.T) {
	// setup types
	p := &Plugin{
		Kubernetes: &Kubernetes{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
