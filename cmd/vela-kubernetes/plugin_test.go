// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestKubernetes_Plugin_Exec_Apply_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	p := &Plugin{
		Apply: &Apply{
			DryRun: "false",
			Files:  []string{"apply.yml"},
			Output: "json",
		},
		Config: &Config{
			Action:    "apply",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
			Path:      "~/.kube/config",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Plugin_Exec_Patch_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "patch",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
			Path:      "~/.kube/config",
		},
		Patch: &Patch{
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
		},
		Status: &Status{},
	}

	err := p.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Plugin_Exec_Status_Error(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "status",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
			Path:      "~/.kube/config",
		},
		Patch: &Patch{},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
			Watch:     false,
		},
	}

	err := p.Exec()
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_Apply(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{
			DryRun: "false",
			Files:  []string{"apply.yml"},
			Output: "json",
		},
		Config: &Config{
			Action:    "apply",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate_Patch(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "patch",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
		},
		Patch: &Patch{
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
		},
		Status: &Status{},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate_Status(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "status",
			Cluster:   "cluster",
			Context:   "context",
			File:      "file",
			Namespace: "namespace",
		},
		Patch: &Patch{},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
			Watch:     false,
		},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate_InvalidAction(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "foobar",
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoConfig(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply:  &Apply{},
		Config: &Config{},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoApply(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "apply",
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoPatch(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "patch",
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoStatus(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			Action:    "status",
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch:  &Patch{},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
