# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG KUBECTL_VERSION=v1.34.1

###############################################################################
##    docker build --no-cache --target binary -t vela-kubernetes:binary .    ##
###############################################################################

FROM alpine:3.22.2@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412 as binary

ARG KUBECTL_VERSION

ADD https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl /bin/kubectl

RUN chmod 0700 /bin/kubectl

#############################################################################
##    docker build --no-cache --target certs -t vela-kubernetes:certs .    ##
#############################################################################

FROM alpine:3.22.2@sha256:4b7ce07002c69e8f3d704a9c5d6fd3053be500b7f1c69fc0d80990c2ad8dd412 as certs

RUN apk add --update --no-cache ca-certificates

#############################################################################
##    docker build --no-cache --target gcloud -t vela-kubernetes:gcloud .    ##
#############################################################################

FROM gcr.io/google.com/cloudsdktool/google-cloud-cli:548.0.0-alpine@sha256:554691b9f70191a0403713a2cd4079b5f425fe27d607a476a42a4853aded84ea as gcloud

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
