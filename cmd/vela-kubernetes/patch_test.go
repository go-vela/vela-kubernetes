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
		Action:    "patch",
		File:      "file",
		Cluster:   "cluster",
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
		DryRun:        false,
		Files:         []string{"patch.yml"},
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	for _, file := range p.Files {
		for _, container := range p.Containers {
			pattern := fmt.Sprintf(patchPattern, container.Name, container.Image)

			want := exec.Command(
				kubectlBin,
				fmt.Sprintf("--namespace=%s", c.Namespace),
				fmt.Sprintf("--cluster=%s", c.Cluster),
				fmt.Sprintf("--context=%s", c.Context),
				"patch",
				fmt.Sprintf("--local=%t", p.DryRun),
				fmt.Sprintf("--filename=%s", file),
				fmt.Sprintf("--patch=%s", pattern),
				fmt.Sprintf("--output=%s", p.Output),
			)

			got := p.Command(c, file, container)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Command is %v, want %v", got, want)
			}
		}
	}
}

func TestKubernetes_Patch_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "patch",
		File:      "file",
		Cluster:   "cluster",
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
		DryRun:        false,
		Files:         []string{"patch.yml"},
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
		DryRun:        false,
		Files:         []string{"patch.yml"},
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Patch_Validate_NoFiles(t *testing.T) {
	// setup types
	p := &Patch{
		DryRun:        false,
		Output:        "json",
		RawContainers: `[{"name": "container", "image": "alpine"}]`,
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_Invalid(t *testing.T) {
	// setup types
	p := &Patch{
		DryRun:        false,
		Files:         []string{"patch.yml"},
		Output:        "json",
		RawContainers: "!@#$%^&*()",
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_NoRawContainers(t *testing.T) {
	// setup types
	p := &Patch{
		DryRun: false,
		Files:  []string{"patch.yml"},
		Output: "json",
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Patch_Validate_NoRawContainerName(t *testing.T) {
	// setup types
	p := &Patch{
		DryRun:        false,
		Files:         []string{"patch.yml"},
		Output:        "json",
		RawContainers: `[{"image": "alpine"}]`,
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
