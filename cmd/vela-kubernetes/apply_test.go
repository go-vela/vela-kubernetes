// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os/exec"
	"reflect"
	"testing"
)

func TestKubernetes_Apply_Command(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
		Path:      "~/.kube/config",
	}

	a := &Apply{
		DryRun: "false",
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	//nolint:gosec // testing purposes
	for _, file := range a.Files {
		want := exec.CommandContext(
			t.Context(),
			_kubectl,
			fmt.Sprintf("--kubeconfig=%s", c.Path),
			fmt.Sprintf("--cluster=%s", c.Cluster),
			fmt.Sprintf("--context=%s", c.Context),
			fmt.Sprintf("--namespace=%s", c.Namespace),
			"apply",
			"--dry-run=none",
			fmt.Sprintf("--filename=%s", file),
			fmt.Sprintf("--output=%s", a.Output),
		)

		got := a.Command(c, file)

		if got.Path != want.Path || !reflect.DeepEqual(got.Args, want.Args) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}
func TestKubernetes_Apply_Command_WithDryRunTrue(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
		Path:      "~/.kube/config",
	}

	a := &Apply{
		DryRun: "true",
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	//nolint:gosec // testing purposes
	for _, file := range a.Files {
		want := exec.CommandContext(
			t.Context(),
			_kubectl,
			fmt.Sprintf("--kubeconfig=%s", c.Path),
			fmt.Sprintf("--cluster=%s", c.Cluster),
			fmt.Sprintf("--context=%s", c.Context),
			fmt.Sprintf("--namespace=%s", c.Namespace),
			"apply",
			"--dry-run=client",
			fmt.Sprintf("--filename=%s", file),
			fmt.Sprintf("--output=%s", a.Output),
		)

		got := a.Command(c, file)

		if got.Path != want.Path || !reflect.DeepEqual(got.Args, want.Args) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}
func TestKubernetes_Apply_Command_WithDryRunAnythingNonBoolean(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		Cluster:   "cluster",
		Context:   "context",
		File:      "file",
		Namespace: "namespace",
		Path:      "~/.kube/config",
	}

	a := &Apply{
		DryRun: "server",
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	//nolint:gosec // testing purposes
	for _, file := range a.Files {
		want := exec.CommandContext(
			t.Context(),
			_kubectl,
			fmt.Sprintf("--kubeconfig=%s", c.Path),
			fmt.Sprintf("--cluster=%s", c.Cluster),
			fmt.Sprintf("--context=%s", c.Context),
			fmt.Sprintf("--namespace=%s", c.Namespace),
			"apply",
			fmt.Sprintf("--dry-run=%s", a.DryRun),
			fmt.Sprintf("--filename=%s", file),
			fmt.Sprintf("--output=%s", a.Output),
		)

		got := a.Command(c, file)

		if got.Path != want.Path || !reflect.DeepEqual(got.Args, want.Args) {
			t.Errorf("Command is %v, want %v", got, want)
		}
	}
}

func TestKubernetes_Apply_Exec_Error(t *testing.T) {
	// setup types
	c := &Config{
		Action:    "apply",
		File:      "file",
		Cluster:   "cluster",
		Context:   "context",
		Namespace: "namespace",
	}

	a := &Apply{
		DryRun: "false",
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	err := a.Exec(c)
	if err == nil {
		t.Errorf("Exec should have returned err")
	}
}

func TestKubernetes_Apply_Validate(t *testing.T) {
	// setup types
	a := &Apply{
		DryRun: "false",
		Files:  []string{"apply.yml"},
		Output: "json",
	}

	err := a.Validate()
	if err != nil {
		t.Errorf("Validate returned err: %v", err)
	}
}

func TestKubernetes_Apply_Validate_NoFiles(t *testing.T) {
	// setup types
	a := &Apply{
		DryRun: "false",
		Output: "json",
	}

	err := a.Validate()
	if err == nil {
		t.Errorf("Validate should have returned err")
	}
}
