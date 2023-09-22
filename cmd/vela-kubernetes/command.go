// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

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

// versionCmd is a helper function to output
// the client and server version information.
func versionCmd(c *Config) *exec.Cmd {
	logrus.Trace("creating kubectl version command")

	// variable to store flags for command
	var flags []string

	// check if config path is provided
	if len(c.Path) > 0 {
		// add flag for path from provided config path
		flags = append(flags, fmt.Sprintf("--kubeconfig=%s", c.Path))
	}

	// check if config cluster is provided
	if len(c.Cluster) > 0 {
		// add flag for cluster from provided config cluster
		flags = append(flags, fmt.Sprintf("--cluster=%s", c.Cluster))
	}

	// check if config context is provided
	if len(c.Context) > 0 {
		// add flag for context from provided config context
		flags = append(flags, fmt.Sprintf("--context=%s", c.Context))
	}

	// check if config namespace is provided
	if len(c.Namespace) > 0 {
		// add flag for namespace from provided config namespace
		flags = append(flags, fmt.Sprintf("--namespace=%s", c.Namespace))
	}

	// add flag for version kubectl command
	flags = append(flags, "version", "--output=yaml")

	return exec.Command(_kubectl, flags...)
}
