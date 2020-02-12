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

func TestKubernetes_Apply_Command(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Context:   "context",
		Namespace: "namespace",
	}

	a := &Apply{
		Files: []string{"apply.yml"},
	}

	for _, file := range a.Files {
		want := exec.Command(
			kubectlBin,
			fmt.Sprintf("--namespace=%s", c.Namespace),
			fmt.Sprintf("--context=%s", c.Context),
			"apply",
			fmt.Sprintf("--filename=%s", file),
			"--output=json",
		)

		got := a.Command(c, file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

func TestKubernetes_Apply_Validate(t *testing.T) {
	// setup types
	a := &Apply{
		Files: []string{"apply.yml"},
	}

	err := a.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Apply_Validate_NoFiles(t *testing.T) {
	// setup types
	a := &Apply{}

	err := a.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
