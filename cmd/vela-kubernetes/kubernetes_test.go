// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestKubernetes_Kubernetes_Validate(t *testing.T) {
	// setup types
	k := &Kubernetes{
		Config:    "config",
		Context:   "context",
		Namespace: "namespace",
	}

	err := k.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Kubernetes_Validate_NoConfig(t *testing.T) {
	// setup types
	k := &Kubernetes{
		Context:   "context",
		Namespace: "namespace",
	}

	err := k.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Kubernetes_Validate_NoContext(t *testing.T) {
	// setup types
	k := &Kubernetes{
		Config:    "config",
		Namespace: "namespace",
	}

	err := k.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Kubernetes_Validate_NoNamespace(t *testing.T) {
	// setup types
	k := &Kubernetes{
		Config:  "config",
		Context: "context",
	}

	err := k.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}

func TestKubernetes_Kubernetes_Write(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// setup types
	k := &Kubernetes{
		Config:    "config",
		Context:   "context",
		Namespace: "namespace",
	}

	err := k.Write()
	if err != nil {
		t.Errorf("Write returned err: %v", err)
	}
}
