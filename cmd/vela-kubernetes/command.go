// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

const kubectlBin = "/bin/kubectl"

// execCmd is a helper function to
// run the provided command.
func execCmd(e *exec.Cmd) error {
	logrus.Tracef("executing cmd %s", strings.Join(e.Args, " "))

	// set command stdout to OS stdout
	e.Stdout = os.Stdout
	// set command stderr to OS stderr
	e.Stderr = os.Stderr

	// output "trace" string for command
	fmt.Println("$", strings.Join(e.Args, " "))

	return e.Run()
}

// applyCmd is a helper function to apply
// the provided configuration for a resource.
func applyCmd(file string) *exec.Cmd {
	logrus.Trace("returning applyCmd")

	return exec.Command(
		kubectlBin,
		"apply",
		"--output",
		"json",
		"--filename",
		file,
	)
}

// patchCmd is a helper function to update
// the fields of a resource using the
// Kubernetes merging strategy.
func patchCmd(file string) *exec.Cmd {
	logrus.Trace("returning patchCmd")

	return exec.Command(
		kubectlBin,
		"patch",
		"--output",
		"json",
		"--filename",
		file,
	)
}

// replaceCmd is a helper function to replace
// the provided configuration for a resource.
func replaceCmd(file string) *exec.Cmd {
	logrus.Trace("returning replaceCmd")

	return exec.Command(
		kubectlBin,
		"replace",
		"--output",
		"json",
		"--filename",
		file,
	)
}

// statusWatchCmd is a helper function to inspect
// the current status of the latest rollout and
// wait for that rollout to finish.
func statusWatchCmd(status string) *exec.Cmd {
	logrus.Trace("returning statusWatchCmd")

	return exec.Command(
		kubectlBin,
		"rollout",
		"status",
		"--watch",
		"true",
		status,
	)
}

// statusNoWatchCmd is a helper function to inspect
// the current status of the latest rollout without
// waiting for that rollout to finish.
func statusNoWatchCmd(status string) *exec.Cmd {
	logrus.Trace("returning statusNoWatchCmd")

	return exec.Command(
		kubectlBin,
		"rollout",
		"status",
		"--watch",
		"false",
		status,
	)
}

// versionCmd is a helper function to output
// the client and server version information
// for the configured context.
func versionCmd() *exec.Cmd {
	logrus.Trace("returning versionCmd")

	return exec.Command(
		kubectlBin,
		"version",
	)
}
