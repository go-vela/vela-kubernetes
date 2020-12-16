// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-vela/vela-kubernetes/version"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-kubernetes"
	app.HelpName = "vela-kubernetes"
	app.Usage = "Vela Kubernetes plugin for managing Kubernetes resources"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []*cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Plugin Metadata

	app.Action = run
	app.Compiled = time.Now()
	app.Version = version.New().Semantic()

	// Plugin Flags

	app.Flags = []cli.Flag{

		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_DRY_RUN", "KUBERNETES_DRY_RUN"},
			FilePath: "/vela/parameters/kubernetes/dry_run,/vela/secrets/kubernetes/dry_run",
			Name:     "dry_run",
			Usage:    "enables pretending to perform the action",
		},
		&cli.StringSliceFlag{
			EnvVars:  []string{"PARAMETER_FILES", "KUBERNETES_FILES"},
			FilePath: "/vela/parameters/kubernetes/files,/vela/secrets/kubernetes/files",
			Name:     "files",
			Usage:    "kubernetes files or directories to perform an action on",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_LOG_LEVEL", "KUBERNETES_LOG_LEVEL"},
			FilePath: "/vela/parameters/kubernetes/log_level,/vela/secrets/kubernetes/log_level",
			Name:     "log.level",
			Usage:    "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:    "info",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_OUTPUT", "KUBERNETES_OUTPUT"},
			FilePath: "/vela/parameters/kubernetes/output,/vela/secrets/kubernetes/output",
			Name:     "output",
			Usage:    "set output for action - options: (json|yaml|wide)",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_VERSION", "KUBERNETES_VERSION"},
			FilePath: "/vela/parameters/kubernetes/version,/vela/secrets/kubernetes/version",
			Name:     "kubectl.version",
			Usage:    "set kubectl version for plugin",
		},

		// Config Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_ACTION", "KUBERNETES_ACTION"},
			FilePath: "/vela/parameters/kubernetes/action,/vela/secrets/kubernetes/action",
			Name:     "config.action",
			Usage:    "action to perform against Kubernetes",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CLUSTER", "KUBERNETES_CLUSTER"},
			FilePath: "/vela/parameters/kubernetes/cluster,/vela/secrets/kubernetes/cluster",
			Name:     "config.cluster",
			Usage:    "kubectl cluster for interacting with Kubernetes",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONTEXT", "KUBERNETES_CONTEXT"},
			FilePath: "/vela/parameters/kubernetes/context,/vela/secrets/kubernetes/context",
			Name:     "config.context",
			Usage:    "kubectl context for interacting with Kubernetes",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONFIG", "KUBERNETES_CONFIG", "KUBE_CONFIG"},
			FilePath: "/vela/parameters/kubernetes/config,/vela/secrets/kubernetes/config",
			Name:     "config.file",
			Usage:    "kubectl configuration for interacting with Kubernetes",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_NAMESPACE", "KUBERNETES_NAMESPACE"},
			FilePath: "/vela/parameters/kubernetes/namespace,/vela/secrets/kubernetes/namespace",
			Name:     "config.namespace",
			Usage:    "kubectl namespace for interacting with Kubernetes",
		},
		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_PATH", "KUBERNETES_PATH"},
			FilePath: "/vela/parameters/kubernetes/path,/vela/secrets/kubernetes/path",
			Name:     "config.path",
			Usage:    "path to kubectl configuration file",
		},

		// Patch Flags

		&cli.StringFlag{
			EnvVars:  []string{"PARAMETER_CONTAINERS", "KUBERNETES_CONTAINERS"},
			FilePath: "/vela/parameters/kubernetes/containers,/vela/secrets/kubernetes/containers",
			Name:     "patch.containers",
			Usage:    "containers from files to patch",
		},

		// Status Flags

		&cli.StringSliceFlag{
			EnvVars:  []string{"PARAMETER_STATUSES", "KUBERNETES_STATUSES"},
			FilePath: "/vela/parameters/kubernetes/resources,/vela/secrets/kubernetes/resources",
			Name:     "status.resources",
			Usage:    "kubernetes resources to watch status on",
		},
		&cli.DurationFlag{
			EnvVars:  []string{"PARAMETER_TIMEOUT", "KUBERNETES_TIMEOUT"},
			FilePath: "/vela/parameters/kubernetes/timeout,/vela/secrets/kubernetes/timeout",
			Name:     "status.timeout",
			Usage:    "maximum duration to watch status on kubernetes resources",
			Value:    5 * time.Minute,
		},
		&cli.BoolFlag{
			EnvVars:  []string{"PARAMETER_WATCH", "KUBERNETES_WATCH"},
			FilePath: "/vela/parameters/kubernetes/watch,/vela/secrets/kubernetes/watch",
			Name:     "status.watch",
			Usage:    "enables watching the status until the rollout completes",
			Value:    true,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
	// capture the version information as pretty JSON
	v, err := json.MarshalIndent(version.New(), "", "  ")
	if err != nil {
		return err
	}

	// output the version information to stdout
	fmt.Fprintf(os.Stdout, "%s\n", string(v))

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
		"docs":     "https://go-vela.github.io/docs/plugins/registry/kubernetes",
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
			DryRun: c.Bool("dry_run"),
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
	err = p.Validate()
	if err != nil {
		return err
	}

	// execute the plugin
	return p.Exec()
}
