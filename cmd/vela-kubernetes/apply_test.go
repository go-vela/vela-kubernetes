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
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	a := &Apply{
		DryRun: false,
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	for _, file := range a.Files {
		want := exec.Command(
			kubectlBin,
			fmt.Sprintf("--namespace=%s", c.Namespace),
			fmt.Sprintf("--cluster=%s", c.Cluster),
			fmt.Sprintf("--context=%s", c.Context),
			"apply",
			fmt.Sprintf("--dry-run=%t", a.DryRun),
			fmt.Sprintf("--filename=%s", file),
			fmt.Sprintf("--output=%s", a.Output),
		)

		got := a.Command(c, file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

func TestKubernetes_Apply_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	a := &Apply{
		DryRun: false,
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	err := a.Exec(c)
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Apply_Validate(t *testing.T) {
	// setup types
	a := &Apply{
		DryRun: false,
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	err := a.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Apply_Validate_NoFiles(t *testing.T) {
	// setup types
	a := &Apply{
		DryRun: false,
		Output: "json",
	}

	err := a.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
