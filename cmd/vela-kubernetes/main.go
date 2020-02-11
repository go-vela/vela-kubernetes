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

		// Kubernetes Flags

		cli.StringFlag{
			EnvVar: "PARAMETER_CONFIG,KUBE_CONFIG,KUBERNETES_CONFIG",
			Name:   "kubernetes.config",
			Usage:  "kubernetes cluster configuration",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_CONTEXT,KUBE_CONTEXT,KUBERNETES_CONTEXT",
			Name:   "kubernetes.context",
			Usage:  "kubernetes cluster context to interact with",
		},
		cli.StringFlag{
			EnvVar: "PARAMETER_NAMESPACE,KUBE_NAMESPACE,KUBERNETES_NAMESPACE",
			Name:   "kubernetes.namespace",
			Usage:  "kubernetes cluster namespace to interact with",
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
		// Kubernetes configuration
		Kubernetes: &Kubernetes{
			Config:    c.String("kubernetes.config"),
			Context:   c.String("kubernetes.context"),
			Namespace: c.String("kubernetes.namespace"),
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
