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
		Containers: []*Container{
			{
				Name:  "container",
				Image: "alpine",
			},
		},
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	for _, container := range p.Containers {
		want := exec.Command(
			kubectlBin,
			fmt.Sprintf("--namespace=%s", c.Namespace),
			fmt.Sprintf("--context=%s", c.Context),
			"patch",
			fmt.Sprintf("--output=%s", p.Output),
		)

		got := p.Command(c, container)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

func TestKubernetes_Patch_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Context:   "context",
		Namespace: "namespace",
	}

	p := &Patch{
		Containers: []*Container{
			{
				Name:  "container",
				Image: "alpine",
			},
		},
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	err := p.Exec(c)
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Patch_Validate(t *testing.T) {
	// setup types
	p := &Patch{
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Patch_Validate_Invalid(t *testing.T) {
	// setup types
	p := &Patch{
		Output:        "json",
		RawContainers: "!@#$%^&*()",
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_NoContainers(t *testing.T) {
	// setup types
	p := &Patch{}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_NoContainerName(t *testing.T) {
	// setup types
	p := &Patch{
		Output:        "json",
		RawContainers: `[{"image": "alpine"}]`,
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_NoContainerImage(t *testing.T) {
	// setup types
	p := &Patch{
		Output:        "json",
		RawContainers: `[{"name": "container"}]`,
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
