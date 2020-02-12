// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestKubernetes_Patch_Command(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Context:   "context",
		Namespace: "namespace",
	}

	p := &Patch{
		Images: []string{"images"},
	}

	want := exec.Command(
		kubectlBin,
		fmt.Sprintf("--namespace=%s", c.Namespace),
		fmt.Sprintf("--context=%s", c.Context),
		"patch",
		"--output=json",
	)

	got := p.Command(c)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Command is %v, want %v", got, want)
	}
}

func TestKubernetes_Patch_Validate(t *testing.T) {
	// setup types
	p := &Patch{
		Images: []string{"images"},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Patch_Validate_NoImages(t *testing.T) {
	// setup types
	p := &Patch{}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
