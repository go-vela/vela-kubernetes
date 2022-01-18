// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
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
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
		Path:      "~/.kube/config",
	}

	want := exec.Command(
		_kubectl,
		fmt.Sprintf("--kubeconfig=%s", c.Path),
		fmt.Sprintf("--cluster=%s", c.Cluster),
		fmt.Sprintf("--context=%s", c.Context),
		fmt.Sprintf("--namespace=%s", c.Namespace),
		"version",
	)

	got := versionCmd(c)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("versionCmd is %v, want %v", got, want)
	}
}
