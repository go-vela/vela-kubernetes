# Copyright (c) 2022 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG KUBECTL_VERSION=v1.24.12

###############################################################################
##    docker build --no-cache --target binary -t vela-kubernetes:binary .    ##
###############################################################################

FROM alpine as binary

ARG KUBECTL_VERSION

ADD https://storage.googleapis.com/kubernetes-release/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl /bin/kubectl

RUN chmod 0700 /bin/kubectl

#############################################################################
##    docker build --no-cache --target certs -t vela-kubernetes:certs .    ##
#############################################################################

FROM alpine as certs

RUN apk add --update --no-cache ca-certificates

#############################################################################
##    docker build --no-cache --target gcloud -t vela-kubernetes:gcloud .    ##
#############################################################################

FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:alpine as gcloud

RUN gcloud components install gke-gcloud-auth-plugin

##############################################################
##    docker build --no-cache -t vela-kubernetes:local .    ##
##############################################################

FROM scratch

ARG KUBECTL_VERSION

ENV PLUGIN_KUBECTL_VERSION=${KUBECTL_VERSION}

COPY --from=binary /bin/kubectl /bin/kubectl

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=gcloud /google-cloud-sdk/bin/gke-gcloud-auth-plugin /bin/gke-gcloud-auth-plugin

COPY release/vela-kubernetes /bin/vela-kubernetes

ENTRYPOINT [ "/bin/vela-kubernetes" ]
