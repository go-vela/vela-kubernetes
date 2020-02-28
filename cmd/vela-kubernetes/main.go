// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := cli.NewApp()

	// Plugin Information

	app.Name = "vela-kubernetes"
	app.HelpName = "vela-kubernetes"
	app.Usage = "Vela Kubernetes plugin for managing Kubernetes resources"
	app.Copyright = "Copyright (c) 2020 Target Brands, Inc. All rights reserved."
	app.Authors = []cli.Author{
		{
			Name:  "Vela Admins",
			Email: "vela@target.com",
		},
	}

	// Plugin Metadata

	app.Compiled = time.Now()
	app.Action = run

	// Plugin Flags

	app.Flags = []cli.Flag{

		cli.StringFlag{
			EnvVar: "PARAMETER_LOG_LEVEL,VELA_LOG_LEVEL,KUBERNETES_LOG_LEVEL",
			Name:   "log.level",
			Usage:  "set log level - options: (trace|debug|info|warn|error|fatal|panic)",
			Value:  "info",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_VERSION,VELA_KUBECTL_VERSION,KUBECTL_VERSION",
			Name:   "kubectl.version",
			Usage:  "set kubectl version for plugin",
		},

		// Apply Flags

		cli.BoolFlag{
			EnvVar: "PARAMETER_DRY_RUN,APPLY_DRY_RUN",
			Name:   "apply.dry_run",
			Usage:  "enables pretending to apply the files",
		},
		cli.StringSliceFlag{
			EnvVar: "PARAMETER_FILES,APPLY_FILES",
			Name:   "apply.files",
			Usage:  "kubernetes files or directories to apply",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_OUTPUT,APPLY_OUTPUT",
			Name:   "apply.output",
			Usage:  "set output for apply - options: (json|yaml|wide)",
		},

		// Config Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_ACTION,CONFIG_ACTION,KUBE_ACTION",
			Name:   "config.action",
			Usage:  "action to perform against Kubernetes",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CLUSTER,CONFIG_CLUSTER,KUBE_CLUSTER",
			Name:   "config.cluster",
			Usage:  "kubectl cluster for interacting with Kubernetes",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CONTEXT,CONFIG_CONTEXT,KUBE_CONTEXT",
			Name:   "config.context",
			Usage:  "kubectl context for interacting with Kubernetes",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CONFIG,CONFIG_FILE,KUBE_CONFIG",
			Name:   "config.file",
			Usage:  "kubectl configuration for interacting with Kubernetes",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_NAMESPACE,CONFIG_NAMESPACE,KUBE_NAMESPACE",
			Name:   "config.namespace",
			Usage:  "kubectl namespace for interacting with Kubernetes",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CONFIG_PATH,CONFIG_PATH,KUBE_CONFIG_PATH",
			Name:   "config.path",
			Usage:  "path to kubectl configuration file",
		},

		// Patch Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_CONTAINERS,PATCH_CONTAINERS",
			Name:   "patch.containers",
			Usage:  "containers from files to patch",
		},
		cli.BoolFlag{
			EnvVar: "PARAMETER_DRY_RUN,PATCH_DRY_RUN",
			Name:   "patch.dry_run",
			Usage:  "enables pretending to patch the containers",
		},
		cli.StringSliceFlag{
			EnvVar: "PARAMETER_FILES,PATCH_FILES",
			Name:   "patch.files",
			Usage:  "kubernetes files or directories to patch",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_OUTPUT,PATCH_OUTPUT",
			Name:   "patch.output",
			Usage:  "set output for patch - options: (json|yaml|wide)",
		},

		// Status Flags

		cli.StringSliceFlag{
			EnvVar: "PARAMETER_STATUSES,STATUS_RESOURCES",
			Name:   "status.resources",
			Usage:  "kubernetes resources to watch status on",
		},
		cli.DurationFlag{
			EnvVar: "PARAMETER_TIMEOUT,STATUS_TIMEOUT",
			Name:   "status.timeout",
			Usage:  "maximum duration to watch status on kubernetes resources",
			Value:  5 * time.Minute,
		},
		cli.BoolTFlag{
			EnvVar: "PARAMETER_WATCH,STATUS_WATCH",
			Name:   "status.watch",
			Usage:  "enables watching the status until the rollout completes",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

// run executes the plugin based off the configuration provided.
func run(c *cli.Context) error {
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
			DryRun: c.Bool("apply.dry_run"),
			Files:  c.StringSlice("apply.files"),
			Output: c.String("apply.output"),
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
			DryRun:        c.Bool("patch.dry_run"),
			Files:         c.StringSlice("patch.files"),
			Output:        c.String("patch.output"),
			RawContainers: c.String("patch.containers"),
		},
		// status configuration
		Status: &Status{
			Resources: c.StringSlice("status.resources"),
			Timeout:   c.Duration("status.timeout"),
			Watch:     c.BoolT("status.watch"),
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
