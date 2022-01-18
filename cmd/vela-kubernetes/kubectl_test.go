// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"testing"

	"github.com/spf13/afero"
)

func TestKubernetes_install(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// run test
	err := install("v0.17.0", "v0.17.0")
	if err != nil {
		t.Errorf("install returned err: %v", err)
	}
}

func TestKubernetes_install_NoBinary(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	// run test
	err := install("v0.15.0", "v0.17.0")
	if err == nil {
		t.Errorf("install should have returned err")
	}
}

func TestKubernetes_install_NotWritable(t *testing.T) {
	// setup filesystem
	appFS = afero.NewMemMapFs()

	a := &afero.Afero{
		Fs: appFS,
	}

	// create binary file
	err := a.WriteFile(_kubectl, []byte("!@#$%^&*()"), 0777)
	if err != nil {
		t.Errorf("Unable to write file %s: %v", _kubectl, err)
	}

	// run test
	err = install("v0.15.0", "v0.17.0")
	if err == nil {
		t.Errorf("install should have returned err")
	}
}
