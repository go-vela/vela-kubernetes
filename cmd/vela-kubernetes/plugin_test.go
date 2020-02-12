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
			Files: []string{"files"},
		},
		Config: &Config{
			File:      "file",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Images: []string{"images"},
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
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Images: []string{"images"},
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
			Files: []string{"files"},
		},
		Config: &Config{},
		Patch: &Patch{
			Images: []string{"images"},
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
			Files: []string{"files"},
		},
		Config: &Config{
			File:      "file",
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
			Files: []string{"files"},
		},
		Config: &Config{
			File:      "file",
			Context:   "context",
			Namespace: "namespace",
		},
		Patch: &Patch{
			Images: []string{"images"},
		},
		Status: &Status{},
	}

	err := p.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
