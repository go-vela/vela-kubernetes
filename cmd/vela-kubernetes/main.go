// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/mail"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"

	_ "github.com/joho/godotenv/autoload"

	"github.com/go-vela/vela-kubernetes/version"
)

func main() {
	// capture application version information
	v := version.New()

	// serialize the version information as pretty JSON
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		logrus.Fatal(err)
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(bytes))

	// create new CLI application
	// Plugin Information
	cmd := cli.Command{
		Name:      "vela-kubernetes",
		Usage:     "Vela Kubernetes plugin for managing Kubernetes resources",
		Copyright: "Copyright 2020 Target Brands, Inc. All rights reserved.",
		Authors: []any{
			&mail.Address{
				Name:    "Vela Admins",
				Address: "vela@target.com",
			},
		},
		Version: v.Semantic(),
		Action:  run,
	}

	// Plugin Metadata

	// Plugin Flags

	cmd.Flags = []cli.Flag{

		&cli.BoolFlag{
			Name:  "dry_run",
			Usage: "enables pretending to perform the action",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_DRY_RUN"),
				cli.EnvVar("KUBERNETES_DRY_RUN"),
				cli.EnvVar("VELA_BUILD_NUMBER"),
				cli.File("//vela/parameters/kubernetes/dry_run"),
				cli.File("/vela/secrets/kubernetes/dry_run"),
			),
		},
		&cli.StringSliceFlag{
			Name:  "files",
			Usage: "kubernetes files or directories to perform an action on",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_FILES"),
				cli.EnvVar("KUBERNETES_FILES"),
				cli.File("/vela/parameters/kubernetes/files"),
				cli.File("/vela/secrets/kubernetes/files"),
			),
		},
		&cli.StringFlag{
			Name:  "log.level",
			Usage: "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value: "info",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_LOG_LEVEL"),
				cli.EnvVar("KUBERNETES_LOG_LEVEL"),
				cli.File("/vela/parameters/kubernetes/log_level"),
				cli.File("/vela/secrets/kubernetes/log_level"),
			),
		},
		&cli.StringFlag{
			Name:  "output",
			Usage: "set output for action - options: (json|yaml|wide)",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_OUTPUT"),
				cli.EnvVar("KUBERNETES_OUTPUT"),
				cli.File("/vela/parameters/kubernetes/output"),
				cli.File("/vela/secrets/kubernetes/output"),
			),
		},
		&cli.StringFlag{
			Name:  "kubectl.version",
			Usage: "set kubectl version for plugin",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_VERSION"),
				cli.EnvVar("KUBERNETES_VERSION"),
				cli.File("/vela/parameters/kubernetes/version"),
				cli.File("/vela/secrets/kubernetes/version"),
			),
		},

		// Config Flags

		&cli.StringFlag{
			Name:  "config.action",
			Usage: "action to perform against Kubernetes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_ACTION"),
				cli.EnvVar("KUBERNETES_ACTION"),
				cli.File("/vela/parameters/kubernetes/action"),
				cli.File("/vela/secrets/kubernetes/action"),
			),
		},
		&cli.StringFlag{
			Name:  "config.cluster",
			Usage: "kubectl cluster for interacting with Kubernetes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_CLUSTER"),
				cli.EnvVar("KUBERNETES_CLUSTER"),
				cli.File("/vela/parameters/kubernetes/cluster"),
				cli.File("/vela/secrets/kubernetes/cluster"),
			),
		},
		&cli.StringFlag{
			Name:  "config.context",
			Usage: "kubectl context for interacting with Kubernetes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_CONTEXT"),
				cli.EnvVar("KUBERNETES_CONTEXT"),
				cli.File("/vela/parameters/kubernetes/context"),
				cli.File("/vela/secrets/kubernetes/context"),
			),
		},
		&cli.StringFlag{
			Name:  "config.file",
			Usage: "kubectl configuration for interacting with Kubernetes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_CONFIG"),
				cli.EnvVar("KUBERNETES_CONFIG"),
				cli.EnvVar("KUBE_CONFIG"),
				cli.File("/vela/parameters/kubernetes/config"),
				cli.File("/vela/secrets/kubernetes/config"),
			),
		},
		&cli.StringFlag{
			Name:  "config.namespace",
			Usage: "kubectl namespace for interacting with Kubernetes",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_NAMESPACE"),
				cli.EnvVar("KUBERNETES_NAMESPACE"),
				cli.File("/vela/parameters/kubernetes/namespace"),
				cli.File("/vela/secrets/kubernetes/namespace"),
			),
		},
		&cli.StringFlag{
			Name:  "config.path",
			Usage: "path to kubectl configuration file",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_PATH"),
				cli.EnvVar("KUBERNETES_PATH"),
				cli.File("/vela/parameters/kubernetes/path"),
				cli.File("/vela/secrets/kubernetes/path"),
			),
		},

		// Patch Flags

		&cli.StringFlag{
			Name:  "patch.containers",
			Usage: "containers from files to patch",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_CONTAINERS"),
				cli.EnvVar("KUBERNETES_CONTAINERS"),
				cli.File("/vela/parameters/kubernetes/containers"),
				cli.File("/vela/secrets/kubernetes/containers"),
			),
		},
		// Status Flags

		&cli.StringSliceFlag{
			Name:  "status.resources",
			Usage: "kubernetes resources to watch status on",
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_STATUSES"),
				cli.EnvVar("KUBERNETES_STATUSES"),
				cli.File("/vela/parameters/kubernetes/resources"),
				cli.File("/vela/secrets/kubernetes/resources"),
			),
		},
		&cli.DurationFlag{
			Name:  "status.timeout",
			Usage: "maximum duration to watch status on kubernetes resources",
			Value: 5 * time.Minute,
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_TIMEOUT"),
				cli.EnvVar("KUBERNETES_TIMEOUT"),
				cli.File("/vela/parameters/kubernetes/timeout"),
				cli.File("/vela/secrets/kubernetes/timeout"),
			),
		},
		&cli.BoolFlag{
			Name:  "status.watch",
			Usage: "enables watching the status until the rollout completes",
			Value: true,
			Sources: cli.NewValueSourceChain(
				cli.EnvVar("PARAMETER_WATCH"),
				cli.EnvVar("KUBERNETES_WATCH"),
				cli.File("/vela/parameters/kubernetes/watch"),
				cli.File("/vela/secrets/kubernetes/watch"),
			),
		},
	}

	err = cmd.Run(context.Background(), os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(_ context.Context, c *cli.Command) error {
	// set the log level for the plugin
	switch c.String("log.level") {
	case "t", "trace", "Trace", "TRACE":
		logrus.SetLevel(logrus.TraceLevel)
	case "d", "debug", "Debug", "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "w", "warn", "Warn", "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "e", "error", "Error", "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "f", "fatal", "Fatal", "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "p", "panic", "Panic", "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	case "i", "info", "Info", "INFO":
		fallthrough
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.WithFields(logrus.Fields{
		"code":     "https://github.com/go-vela/vela-kubernetes",
		"docs":     "https://go-vela.github.io/docs/plugins/registry/pipeline/kubernetes",
		"registry": "https://hub.docker.com/r/target/vela-kubernetes",
	}).Info("Vela Kubernetes Plugin")

	// capture custom kubectl version requested
	version := c.String("kubectl.version")

	// check if a custom kubectl version was requested
	if len(version) > 0 {
		// attempt to install the custom kubectl version
		err := install(version, os.Getenv("PLUGIN_KUBECTL_VERSION"))
		if err != nil {
			return err
		}
	}

	// create the plugin
	p := &Plugin{
		// apply configuration
		Apply: &Apply{
			DryRun: c.String("dry_run"),
			Files:  c.StringSlice("files"),
			Output: c.String("output"),
		},
		// config configuration
		Config: &Config{
			Action:    c.String("config.action"),
			Cluster:   c.String("config.cluster"),
			Context:   c.String("config.context"),
			File:      c.String("config.file"),
			Namespace: c.String("config.namespace"),
			Path:      c.String("config.path"),
		},
		// patch configuration
		Patch: &Patch{
			DryRun:        c.Bool("dry_run"),
			Files:         c.StringSlice("files"),
			Output:        c.String("output"),
			RawContainers: c.String("patch.containers"),
		},
		// status configuration
		Status: &Status{
			Resources: c.StringSlice("status.resources"),
			Timeout:   c.Duration("status.timeout"),
			Watch:     c.Bool("status.watch"),
		},
	}

	// validate the plugin
	err := p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
