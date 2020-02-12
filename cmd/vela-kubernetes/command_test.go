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

func TestKubernetes_execCmd(t *testing.T) {
	// setup types
	e := exec.Command("echo", "hello")

	err := execCmd(e)
	if err != nil {
		t.Errorf("execCmd returned err: %v", err)
	}
}

func TestKubernetes_versionCmd(t *testing.T) {
	// setup types
	c := &Config{
		File:      "file",
		Context:   "context",
		Namespace: "namespace",
	}

	want := exec.Command(
		kubectlBin,
		fmt.Sprintf("--namespace=%s", c.Namespace),
		fmt.Sprintf("--context=%s", c.Context),
		"version",
	)

	got := versionCmd(c)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("versionCmd is %v, want %v", got, want)
	}
}
