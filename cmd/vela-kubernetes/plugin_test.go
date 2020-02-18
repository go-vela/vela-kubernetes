// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
	"time"
)

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
		Apply: &Apply{
			Files:  []string{"apply.yml"},
			Output: "json",
		},
		Config: &Config{
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Containers: []*Container{
				{
					Name:  "container",
					Image: "alpine",
				},
			},
			Output:        "json",
			RawContainers: `[{"name": "container", "image": "alpine"}]`,
		},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
		},
	}

	err := p.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Plugin_Validate_NoApply(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{},
		Config: &Config{
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Containers: []*Container{
				{
					Name:  "container",
					Image: "alpine",
				},
			},
			Output:        "json",
			RawContainers: `[{"name": "container", "image": "alpine"}]`,
		},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
		},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoConfig(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{
			Files:  []string{"apply.yml"},
			Output: "json",
		},
		Config: &Config{},
		Patch: &Patch{
			Containers: []*Container{
				{
					Name:  "container",
					Image: "alpine",
				},
			},
			Output:        "json",
			RawContainers: `[{"name": "container", "image": "alpine"}]`,
		},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
		},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoPatch(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{
			Files:  []string{"apply.yml"},
			Output: "json",
		},
		Config: &Config{
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{},
		Status: &Status{
			Resources: []string{"resources"},
			Timeout:   5 * time.Minute,
		},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Plugin_Validate_NoStatus(t *testing.T) {
	// setup types
	p := &Plugin{
		Apply: &Apply{
			Files: []string{"apply.yml"},
		},
		Config: &Config{
			File:      "file",
			Cluster:   "cluster",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Containers: []*Container{
				{
					Name:  "container",
					Image: "alpine",
				},
			},
			Output:        "json",
			RawContainers: `[{"name": "container", "image": "alpine"}]`,
		},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
