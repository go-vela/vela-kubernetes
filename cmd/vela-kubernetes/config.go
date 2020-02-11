// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import "time"

// Config represents the plugin configuration for Kubernetes config information.
type Config struct {
	// Kubernetes files or directories to apply
	Files []string
	// new container images from files to apply
	Images []string
	// Kubernetes resources to watch status of
	Statuses []string
	// total time allowed to watch Kubernetes resources
	Timeout time.Duration
}
