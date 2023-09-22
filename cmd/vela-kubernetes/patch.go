// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

const patchAction = "patch"

// patchPattern represents the pattern needed to
// patch a Kubernetes container with a new image.
//
// spec: (replica set)
//
//	template:
//	  spec: (pod)
//	    containers:
//	      - name:
//	        image:
const patchPattern = `
{
  "spec": {
    "template": {
      "spec": {
        "containers": [
          {
            "name": "%s",
            "image": "%s"
          }
        ]
      }
    }
  }
}
`

// Patch represents the plugin configuration for Patch config information.
type Patch struct {
	// container images from files to patch
	Containers []*Container
	// enables pretending to patch the containers from the files
	DryRun bool
	// Kubernetes files or directories to patch
	Files []string
	// sets the output for the patch command
	Output string
	// raw input of containers provided for plugin
	RawContainers string
}

// Command formats and outputs the Patch command from
// the provided configuration to patch resources.
func (p *Patch) Command(c *Config, file string, container *Container) *exec.Cmd {
	logrus.Tracef("creating kubectl patch command for %s from plugin configuration", container.Name)

	// create pattern for patching containers
	pattern := fmt.Sprintf(patchPattern, container.Name, container.Image)

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

	// add flag for patch kubectl command
	flags = append(flags, "patch")

	// add flag for dry run from provided patch dry run
	flags = append(flags, fmt.Sprintf("--local=%t", p.DryRun))

	// add flag for file from provided patch file
	flags = append(flags, fmt.Sprintf("--filename=%s", file))

	// add flag for the patch to be made
	flags = append(flags, fmt.Sprintf("--patch=%s", pattern))

	// check if patch output is provided
	if len(p.Output) > 0 {
		// add flag for output from provided patch output
		flags = append(flags, fmt.Sprintf("--output=%s", p.Output))
	}

	return exec.Command(_kubectl, flags...)
}

// Exec formats and runs the commands for patching
// the provided configuration to the resources.
func (p *Patch) Exec(c *Config) error {
	logrus.Debug("running patch with provided configuration")

	// iterate through all files to patch
	for _, file := range p.Files {
		// iterate through all images to patch
		for _, container := range p.Containers {
			// create the patch command for the file from the image
			cmd := p.Command(c, file, container)

			// run the patch command for the image
			err := execCmd(cmd)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Unmarshal captures the provided containers and
// serializes them into their expected form.
func (p *Patch) Unmarshal() error {
	logrus.Trace("unmarshaling raw containers")

	// cast raw containers into bytes
	bytes := []byte(p.RawContainers)

	// serialize raw properties into expected Containers type
	err := json.Unmarshal(bytes, &p.Containers)
	if err != nil {
		return err
	}

	return nil
}

// Validate verifies the Patch is properly configured.
func (p *Patch) Validate() error {
	logrus.Trace("validating patch configuration")

	// verify files are provided
	if len(p.Files) == 0 {
		return fmt.Errorf("no patch files provided")
	}

	// verify containers are provided
	if len(p.RawContainers) == 0 {
		return fmt.Errorf("no patch containers provided")
	}

	// serialize provided containers into expected type
	err := p.Unmarshal()
	if err != nil {
		return fmt.Errorf("unable to unmarshal patch containers: %w", err)
	}

	// iterate through all containers
	for _, container := range p.Containers {
		// verify the container is valid
		err := container.Validate()
		if err != nil {
			return fmt.Errorf("invalid patch container provided: %w", err)
		}
	}

	return nil
}
