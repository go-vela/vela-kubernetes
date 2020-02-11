// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
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

func TestKubernetes_applyCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"apply",
		"--output",
		"json",
		"--filename",
		"apply.yml",
	)

	got := applyCmd("apply.yml")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("applyCmd is %v, want %v", got, want)
	}
}

func TestKubernetes_patchCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"patch",
		"--output",
		"json",
		"--filename",
		"patch.yml",
	)

	got := patchCmd("patch.yml")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("patchCmd is %v, want %v", got, want)
	}
}

func TestKubernetes_replaceCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"replace",
		"--output",
		"json",
		"--filename",
		"replace.yml",
	)

	got := replaceCmd("replace.yml")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("replaceCmd is %v, want %v", got, want)
	}
}

func TestKubernetes_statusWatchCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"rollout",
		"status",
		"--watch",
		"true",
		"myStatus",
	)

	got := statusWatchCmd("myStatus")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("statusWatchCmd is %v, want %v", got, want)
	}
}

func TestKubernetes_statusNoWatchCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"rollout",
		"status",
		"--watch",
		"false",
		"myStatus",
	)

	got := statusNoWatchCmd("myStatus")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("statusNoWatchCmd is %v, want %v", got, want)
	}
}

func TestKubernetes_versionCmd(t *testing.T) {
	// setup types
	want := exec.Command(
		kubectlBin,
		"version",
	)

	got := versionCmd()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("versionCmd is %v, want %v", got, want)
	}
}
