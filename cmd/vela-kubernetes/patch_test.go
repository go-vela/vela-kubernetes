// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
)

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
