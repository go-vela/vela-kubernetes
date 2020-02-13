// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/sirupsen/logrus"
)

// cronPatch represents the pattern needed to
// patch a Kubernetes CronJob with a new image.
//
// CRONJOB:
//   spec: (job)
//     jobTemplate:
//       spec: (replica)
//         template:
//           spec: (pod)
//             containers:
//           metadata:
//             annotations:
//               commit_sha:
const cronPatch = `
spec:
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: %s
              image: %s
          metadata:
            annotations:
              commit_sha:
`

// deploymentPatch represents the pattern needed to
// patch a Kubernetes Deployment with a new image.
//
// DEPLOYMENT/DAEMONSET:
//   spec: (replica set)
//     template:
//       spec: (pod)
//         containers:
//       metadata:
//         annotations:
//           commit_sha:
const deploymentPatch = `
spec:
  template:
    spec:
      containers:
        - name: %s
          image: %s
      metadata:
        annotations:
          commit_sha:
`

// Container represents the plugin configuration for patching a container.
//
// TODO: Enable patching of other container attributes?
// * environment variables
// * image pull policy
// * ports
// * resources
//
// This is possible if we made Container a map[string]string
type Container struct {
	// name of the container to patch
	Name string
	// image of the container to patch
	Image string
}

// Validate verifies the Container is properly configured.
func (c *Container) Validate() error {
	logrus.Tracef("validating container configuration for %s", c.Name)

	// verify name is provided
	if len(c.Name) == 0 {
		return fmt.Errorf("no container name provided")
	}

	// verify image is provided
	if len(c.Image) == 0 {
		return fmt.Errorf("no container image provided")
	}

	return nil
}

// Patch represents the plugin configuration for Patch config information.
type Patch struct {
	// raw input of containers provided for plugin
	RawContainers string
	// container images from files to patch
	Containers []*Container
}

// Command formats and outputs the Patch command from
// the provided configuration to patch resources.
func (p *Patch) Command(c *Config, container *Container) *exec.Cmd {
	logrus.Tracef("creating kubectl patch command for %s from plugin configuration", container.Name)

	// variable to store flags for command
	var flags []string

	// check if config namespace is provided
	if len(c.Namespace) > 0 {
		// add flag for namespace from provided config namespace
		flags = append(flags, fmt.Sprintf("--namespace=%s", c.Namespace))
	}

	// check if config context is provided
	if len(c.Context) > 0 {
		// add flag for context from provided config context
		flags = append(flags, fmt.Sprintf("--context=%s", c.Context))
	}

	// add flag for patch kubectl command
	flags = append(flags, "patch")

	// add flag for output
	flags = append(flags, "--output=json")

	return exec.Command(kubectlBin, flags...)
}

// Exec formats and runs the commands for patching
// the provided configuration to the resources.
func (p *Patch) Exec(c *Config) error {
	logrus.Debug("running patch with provided configuration")

	// iterate through all images to patch
	for _, container := range p.Containers {
		// create the patch command for the image
		cmd := p.Command(c, container)

		// run the patch command for the image
		err := execCmd(cmd)
		if err != nil {
			return err
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

	// verify containers are provided
	if len(p.RawContainers) == 0 {
		return fmt.Errorf("no patch containers provided")
	}

	// serialize provided containers into expected type
	err := p.Unmarshal()
	if err != nil {
		return fmt.Errorf("unable to unmarshal patch containers: %v", err)
	}

	// iterate through all containers
	for _, container := range p.Containers {
		// verify the container is valid
		err := container.Validate()
		if err != nil {
			return fmt.Errorf("invalid patch container provided: %v", err)
		}
	}

	return nil
}
