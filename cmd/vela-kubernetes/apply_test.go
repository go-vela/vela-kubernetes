// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"
)

func TestKubernetes_Apply_Validate(t *testing.T) {
	// setup types
	a := &Apply{
		Files: []string{"files"},
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
