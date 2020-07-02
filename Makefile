# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

.PHONY: build
build: binary-build

.PHONY: run
run: build docker-build docker-run

.PHONY: test
test: build docker-build docker-example

#################################
######      Go clean       ######
#################################

.PHONY: clean
clean:

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@echo "I'm kind of the only name in clean energy right now"

#################################
######    Build Binary     ######
#################################

.PHONY: binary-build
binary-build:

	GOOS=linux CGO_ENABLED=0 go build -o release/vela-kubernetes github.com/go-vela/vela-kubernetes/cmd/vela-kubernetes

#################################
######    Docker Build     ######
#################################

.PHONY: docker-build
docker-build:

	docker build --no-cache -t vela-kubernetes:local .

#################################
######     Docker Run      ######
#################################

.PHONY: docker-run
docker-run:

	docker run --rm \
		-e KUBE_CONFIG \
		-e PARAMETER_ACTION \
		-e PARAMETER_CLUSTER \
		-e PARAMETER_CONTAINERS \
		-e PARAMETER_CONTEXT \
		-e PARAMETER_DRY_RUN \
		-e PARAMETER_FILES \
		-e PARAMETER_LOG_LEVEL \
		-e PARAMETER_NAMESPACE \
		-e PARAMETER_OUTPUT \
		-e PARAMETER_PATH \
		-e PARAMETER_RESOURCES \
		-e PARAMETER_TIMEOUT \
		-e PARAMETER_WATCH \
		-e PARAMETER_VERSION \
		vela-kubernetes:local

.PHONY: docker-example
docker-example:

	docker run --rm \
		-e KUBE_CONFIG \
		-e PARAMETER_ACTION=apply \
		-e PARAMETER_CONTEXT=docker-desktop \
		-e PARAMETER_DRY_RUN=true \
		-e PARAMETER_LOG_LEVEL=trace \
		-e PARAMETER_FILES=/examples \
		-v $(shell pwd)/examples:/examples \
		vela-kubernetes:local
