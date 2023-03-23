## Description

This plugin enables the ability to manage resources in [Kubernetes](https://kubernetes.io/) in a Vela pipeline.

Source Code: https://github.com/go-vela/vela-kubernetes

Registry: https://hub.docker.com/r/target/vela-kubernetes

## Usage

> **NOTE:**
>
> Users should refrain from using latest as the tag for the Docker image.
>
> It is recommended to use a semantically versioned tag instead.

Sample of applying Kubernetes files:

```yaml
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: apply
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
```

Sample of pretending to apply Kubernetes files:

```diff
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: apply
+     dry_run: true
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
```

Sample of patching containers in Kubernetes files:

```yaml
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: patch
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
      containers:
        - name: sample
          image: alpine:latest
```

Sample of pretending to patch containers in Kubernetes files:

```diff
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: patch
+     dry_run: true
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
      containers:
        - name: sample
          image: alpine:latest
```

Sample of watching the status of resources:

```yaml
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: status
      statuses: [ sample ]
```

### GKE clusters authentication

Google Kubernetes Engine (GKE) clusters version 1.26 and up require use of new "gke-gcloud-auth-plugin", see [blog](https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke).

1. Download Google service account credentials:

```bash
gcloud beta secrets versions access 1 --secret [Secret Name] --project [Google Project Name] --out-file [Secret Name]-gsa-key.json
```

2. Upload Google service account secret to vela native secrets store:

```bash
vela add secret --secret.engine native --secret.type org --org MYORGNAME --name k8s-gsa-key --value @k8s-gsa-key.json -event deployment --event pull_request --event push --event tag --event comment
```
3. Create kubeconfig file and make sure to use user  `gke-gcloud-auth-user` for your cluster:


```yaml
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: XXXXXXXXXXX
    server: https://ip of K8s IP Server
  name: clustername
contexts:
- context:
    cluster: clustername
    user: gke-gcloud-auth-user
  name: clustername

current-context: ""
kind: Config
preferences: {}
users:
- name: gke-gcloud-auth-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
      - --use_application_default_credentials
      command: gke-gcloud-auth-plugin
      env: null
      installHint: Install gke-gcloud-auth-plugin for use with kubectl by following
        https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke
      interactiveMode: IfAvailable
      provideClusterInfo: true

```
4. Upload kubeconfig file to vela native secretrs store: 

```bash
vela add secret --secret.engine native --secret.type org --org MYORGNAME --name kube_config_secret --value @kube_config_secret.file -event deployment --event pull_request --event push --event tag --event comment
```

5. Configure Vela file:

```yaml
secrets:
  - name: kube_config_secret
    key: ORGNAME/kube-config
    type: org
  - name: gsa_key
    key: ORGNAME/k8s-gsa-key
    type: org

  k8s-gsa-setup:
    steps:
    - name: setup-gke-access
      image: alpine:latest
      ruleset:
        branch: master
        event: [pull_request, push, deployment]
      pull: always
      environment:
        GOOGLE_APPLICATION_CREDENTIALS: "/vela/secrets/.kube/gsa-key.json" # <== change name as needed
      secrets:
        - source: gsa_key
          target: GSA_KEY_FILE
      commands:
        - mkdir -p /vela/secrets/.kube/
        - echo "$GSA_KEY_FILE" > $GOOGLE_APPLICATION_CREDENTIALS

  k8s-plugin-dry-run:
    needs: [ k8s-gsa-setup ]
    steps:
    - name: dry-run
      image: target/vela-kubernetes:latest
      ruleset:
        branch: master
        event: [ pull_request ]
      environment:
        GOOGLE_APPLICATION_CREDENTIALS: "/vela/secrets/.kube/gsa-key.json" # <== change name as needed
        USE_GKE_GCLOUD_AUTH_PLUGIN: True # <== required
      secrets:
        - source: kube_config_secret
          target: kube_config
      parameters:
        action: apply
        dry_run: true
        context: context-name   # <== change as needed
        files: [ Kubernetes/manifests/service.yaml ]

  k8s-plugin-apply:
    needs: [ k8s-gsa-setup ]
    steps:
    - name: run-apply
      image: target/vela-kubernetes:latest
      ruleset:
        branch: master
        event: [ push ]
      environment:
        GOOGLE_APPLICATION_CREDENTIALS: "/vela/secrets/.kube/gsa-key.json" # <== change name as needed
        USE_GKE_GCLOUD_AUTH_PLUGIN: True # <== required
      secrets:
        - source: kube_config_secret
          target: kube_config
      parameters:
        action: apply
        context: context-name   # <== change as needed
        files: [ Kubernetes/manifests/service.yaml ]
```

## Secrets

> **NOTE:** Users should refrain from configuring sensitive information in your pipeline in plain text.

### Internal

Users can use [Vela internal secrets](https://go-vela.github.io/docs/tour/secrets/) to substitute these sensitive values at runtime:

```diff
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
+   secrets: [ kube_config ]
    parameters:
      action: apply
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
-     config: |
-     ---
-     apiVersion: v1
-     kind: Config
```

> This example will add the secrets to the `kubernetes` step as environment variables:
>
> * `KUBE_CONFIG=<value>`

### External

The plugin accepts the following files for authentication:

| Parameter  | Volume Configuration                                                    |
| ---------- | ----------------------------------------------------------------------- |
| `config`   | `/vela/parameters/kubernetes/config`, `/vela/secrets/kubernetes/config` |

Users can use [Vela external secrets](https://go-vela.github.io/docs/concepts/pipeline/secrets/origin/) to substitute these sensitive values at runtime:

```diff
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
+   secrets: [ kube_config ]
    parameters:
      action: apply
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
-     config: |
-     ---
-     apiVersion: v1
-     kind: Config
```

> This example will read the secret values in the volume stored at `/vela/secrets/`

## Parameters

> **NOTE:**
>
> The plugin supports reading all parameters via environment variables or files.
>
> Any values set from a file take precedence over values set from the environment.

The following parameters are used to configure the image:

| Name          | Description                                                     | Required | Default           | Environment Variables                                      |
| ------------- | --------------------------------------------------------------- | -------- | ----------------- | ---------------------------------------------------------- |
| `action`      | action to perform against Kubernetes                            | `true`   | `N/A`             | `PARAMETER_ACTION`<br>`KUBERNETES_ACTION`                  |
| `cluster`     | Kubernetes cluster from the configuration file                  | `false`  | `N/A`             | `PARAMETER_CLUSTER`<br>`KUBERNETES_CLUSTER`                |
| `context`     | Kubernetes context from the configuration file                  | `false`  | `N/A`             | `PARAMETER_CONTEXT`<br>`KUBERNETES_CONTEXT`                |
| `config`      | content of configuration file for communication with Kubernetes | `true`   | `N/A`             | `PARAMETER_CONFIG`<br>`KUBERNETES_CONFIG`<br>`KUBE_CONFIG` |
| `log_level`   | set the log level for the plugin                                | `true`   | `info`            | `PARAMETER_LOG_LEVEL`<br>`KUBERNETES_LOG_LEVEL`            |
| `namespace`   | Kubernetes namespace from the configuration file                | `false`  | `N/A`             | `PARAMETER_NAMESPACE`<br>`KUBERNETES_NAMESPACE`            |
| `path`        | path to configuration file for communication with Kubernetes    | `false`  | `N/A`             | `PARAMETER_PATH`<br>`KUBERNETES_PATH`                      |
| `version`     | version of the `kubectl` CLI to install                         | `false`  | `v1.24.0`         | `PARAMETER_VERSION`<br>`KUBERNETES_VERSION`                |

#### Apply

The following parameters are used to configure the `apply` action:

| Name      | Description                                      | Required | Default | Environment Variables                       |
| --------- | ------------------------------------------------ | -------- | ------- | ------------------------------------------- |
| `dry_run` | enables pretending to perform the apply          | `false`  | `false` | `PARAMETER_DRY_RUN`<br>`KUBERNETES_DRY_RUN` |
| `files`   | list of Kubernetes files or directories to apply | `true`   | `N/A`   | `PARAMETER_FILES`<br>`KUBERNETES_FILES`     |
| `output`  | set the output for the apply                     | `false`  | `N/A`   | `PARAMETER_OUTPUT`<br>`KUBERNETES_OUTPUT`   |

#### Patch

The following parameters are used to configure the `patch` action:

| Name         | Description                                      | Required | Default | Environment Variables                             |
| ------------ | ------------------------------------------------ | -------- | ------- | ------------------------------------------------- |
| `containers` | containers from the files to patch               | `true`   | `N/A`   | `PARAMETER_CONTAINERS`<br>`KUBERNETES_CONTAINERS` |
| `dry_run`    | enables pretending to perform the patch          | `false`  | `false` | `PARAMETER_DRY_RUN`<br>`KUBERNETES_DRY_RUN`       |
| `files`      | list of Kubernetes files or directories to patch | `true`   | `N/A`   | `PARAMETER_FILES`<br>`KUBERNETES_FILES`           |
| `output`     | set the output for the patch                     | `false`  | `N/A`   | `PARAMETER_OUTPUT`<br>`KUBERNETES_OUTPUT`         |

#### Status

The following parameters are used to configure the `status` action:

| Name       | Description                                      | Required | Default | Environment Variables                         |
| ---------- | ------------------------------------------------ | -------- | ------- | --------------------------------------------- |
| `statuses` | list of Kubernetes resources to watch status on  | `true`   | `N/A`   | `PARAMETER_STATUSES`<br>`KUBERNETES_STATUSES` |
| `timeout`  | total time allowed to watch Kubernetes resources | `true`   | `5m`    | `PARAMETER_TIMEOUT`<br>`KUBERNETES_TIMEOUT`   |
| `watch`    | enables watching until the resource completes    | `false`  | `true`  | `PARAMETER_WATCH`<br>`KUBERNETES_WATCH`       |

## Template

COMING SOON!

## Troubleshooting

You can start troubleshooting this plugin by tuning the level of logs being displayed:

```diff
steps:
  - name: kubernetes
    image: target/vela-kubernetes:latest
    pull: always
    parameters:
      action: apply
      files: [ kubernetes/common, kubernetes/dev/deploy.yml ]
+     log_level: trace
```

Below are a list of common problems and how to solve them:
