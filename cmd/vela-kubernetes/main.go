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

		// Apply Flags

		cli.StringSliceFlag{
			EnvVar: "PARAMETER_FILES,APPLY_FILES",
			Name:   "apply.files",
			Usage:  "kubernetes files or directories to apply",
		},

		// Config Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_CONFIG,CONFIG_FILE,KUBE_CONFIG",
			Name:   "config.file",
			Usage:  "kubernetes cluster configuration",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CONTEXT,CONFIG_CONTEXT,KUBE_CONTEXT",
			Name:   "config.context",
			Usage:  "kubernetes cluster context to interact with",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_NAMESPACE,CONFIG_NAMESPACE,KUBE_NAMESPACE",
			Name:   "config.namespace",
			Usage:  "kubernetes cluster namespace to interact with",
		},

		// Patch Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_CONTAINERS,PATCH_CONTAINERS",
			Name:   "patch.containers",
			Usage:  "containers from files to patch",
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
		"code": "https://github.com/go-vela/vela-kubernetes",
		"docs": "https://go-vela.github.io/docs/plugins/registry/kubernetes",
		"time": time.Now(),
	}).Info("Vela Kubernetes Plugin")

	// create the plugin
	p := &Plugin{
		// apply configuration
		Apply: &Apply{
			Files: c.StringSlice("apply.files"),
		},
		// config configuration
		Config: &Config{
			File:      c.String("config.file"),
			Context:   c.String("config.context"),
			Namespace: c.String("config.namespace"),
		},
		// patch configuration
		Patch: &Patch{
			RawContainers: c.String("patch.containers"),
		},
		// status configuration
		Status: &Status{
			Resources: c.StringSlice("status.resources"),
			Timeout:   c.Duration("status.timeout"),
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
