<!--
SPDX-FileCopyrightText: 2022 2020-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->

# Deploying µONOS micro-services with HELM

One of the goals of the µONOS project is to provide simple deployment options
that integrate with modern technologies. Deployment configurations can be found in
the `onos-helm-charts` [repository](https://github.com/onosproject/onos-helm-charts).
Each `onos` service has a directory containing its chart.
As an example the `onos-config` chart is in `onos-helm-charts/onos-config`.

## Deploying on Kubernetes with Helm

[Helm] is a package manager for [Kubernetes] that allows projects to provide a
collection of templates for all the resources needed to deploy on k8s. ONOS Config
provides a Helm chart for deploying a cluster for development and testing. In the
future, this chart will be extended for production use.

## Individual charts or overarching (umbrella) chart
The individual components of µONOS may be deployed one at a time, or altogether through
an overarching (unbrella) Helm chart, or some combination of both.

In all cases the prerequisites must be satisfied:

* Creation of a namespace
* deployment of Atomix controller(s) in the namespace.

The individual components in the [umbrella chart](https://github.com/onosproject/onos-helm-charts/tree/master/sd-ran) are:

*   onos-topo:
*   onos-config:
*   onos-cli:
*   onos-gui:

> The choice of which of these is deployed can be chosen at deployment time with an option like:
>
> `--set import.<component>.enabled=true`
>
> In this way all, none or some of the components can be deployed together.

### Resources

The Helm chart provides resources for deploying the config service and accessing
it over the network, both inside and outside the k8s cluster:

*  `Deployment` - Provides a template for ONOS Config pods
*  `ConfigMap` - Provides test configurations for the application
*  `Service` - Exposes ONOS Config to other applications on the network
*  `Secret` - Provides TLS certificates for end-to-end encryption
*  `Ingress` - Optionally provides support for external load balancing

## Deployment options

### Local Deployment Setup with Kind

To deploy the Helm chart locally:

* First, you will need to install [Docker] to build and deploy an image locally.

* Second, install [Kind].
> **Kind v0.11.0 at least** is required, which provides the **K8S API v1.21**

* Third, install [Helm] version 3. On OSX, this Helm can be installed using [Brew]:

```bash
brew install helm
```
> For more information, please refer to [Installing Helm] page.

* Once Kind has been installed, start it with

```bash
kind create cluster
```

* Once Kind has started, export the configuration:
> This needs to be refreshed if you delete and recreate the Kind cluster
```bash
kind get kubeconfig > ~/.kube/kind
```

* Once Kind has started, export the new environment to access the Kubernetes cluster:
> This needs to be run in each terminal window that will access the cluster
```bash
export KUBECONFIG=~/.kube/kind
```

### Bare metal deployment
ONOS can also be deployed on a bare metal cluster provisioned with [Rancher] or equivalent.

[Kubectl] and [Helm] are can be run from your local PC to control the remote cluster.

## Prerequisites
For any deployment scenario a number of steps must be performed first.

The steps below assume the `KUBECONFIG` environment variable to point `kubectl` to your cluster.

### Add the "CORD" Helm chart repo
The Prometheus and Grafana installations are derived from the CORD Helm charts. Run:
```bash
helm repo add cord https://charts.opencord.org
```

### Add the "Atomix" Helm chart repo
```bash
helm repo add atomix https://charts.atomix.io
```

### Add the "onosproject" Helm chart repo
```bash
helm repo add onosproject https://charts.onosproject.org
```

### Update the local cache with charts from these repos:
```bash
helm repo update
```

### Inspect the chart versions and app versions
To see the list of the latest chart versions and app versions, use the "search" command.
```bash
helm search repo onos
```


### Configure the micro-onos namespace
The various `onos` services can be deployed to any `namespace`.
> A "namespace" partitions the cluster in to independent islands.
> For consistency between documentation we use `micro-onos` as the namespace here.

To create the `micro-onos` namespace run:
```bash
kubectl create namespace micro-onos
```

## Deploy Atomix Controller

The various `onos` services leverage Atomix as the distributed store for HA, scale and redundancy.
The first thing that needs to be deployed in any `onos` deployment is the Atomix
Custom Resource Definitions (CRDs) and Go controller.

The controllers should always be deployed in the `kube-system` namespace.

```bash
helm install -n kube-system atomix atomix/atomix
```

## Deploy ONOS Operator
`onos-operator` ensures that ONOS Custom Resource Defintions (CRD) and their
controllers for `onos-topo` and `onos-config` are deployed in to the cluster.

The controllers should always be deployed in the `kube-system` namespace.

```bash
helm install -n kube-system onos-operator onosproject/onos-operator
```

## Deploy the µONOS services
A complete set of µONOS services can be deployed with just the over-arching
[`onos-umbrella` chart](https://github.com/onosproject/onos-helm-charts/tree/master/onos-umbrella).

Run the install:
```bash
helm -n micro-onos install onos-umbrella onosproject/onos-umbrella
```

this will deploy `onos-topo`, `onos-cli`, `onos-gui`, and `onos-config` (but not `onos-classic` as
it's not needed for µONOS - see [Deploying ONOS classic with HELM](./deploy_onos_classic.md)).

To monitor the startup of the pods use `kubectl` like:
```bash
kubectl -n micro-onos get pods -w
```

giving a list like:
```
NAMESPACE            NAME                            READY   STATUS    RESTARTS   AGE
micro-onos           onos-cli-b6ff8cf9b-2phln        1/1     Running   0          77s
micro-onos           onos-config-85d988999-wjlrv     4/4     Running   0          77s
micro-onos           onos-consensus-store-1-0        1/1     Running   0          75s
micro-onos           onos-gui-847f99659-7bfhj        2/2     Running   0          77s
micro-onos           onos-topo-65978f8c7c-5c4hw      3/3     Running   0          77s
```

Additionally, the Controllers for Atomix and Onos-Operator can be seen in the `kube-system` namespace:
```
NAMESPACE            NAME                                               READY   STATUS    RESTARTS   AGE
kube-system          atomix-consensus-controller-86db6d654b-48xsl       1/1     Running   0          2m5s
kube-system          atomix-controller-755867f47-9526c                  1/1     Running   0          2m5s
kube-system          atomix-pod-memory-controller-589445f57b-v5s5q      1/1     Running   0          2m5s
kube-system          atomix-raft-controller-59874f47cf-ljjxh            1/1     Running   0          2m5s
kube-system          atomix-runtime-controller-7878645d58-tnwcg         1/1     Running   0          2m5s
kube-system          atomix-shared-memory-controller-5bd59cd49c-gsq9p   1/1     Running   0          2m5s
kube-system          atomix-sidecar-controller-8559bfd6c-rjjmg          1/1     Running   0          2m5s
kube-system          onos-operator-app-55997c5989-8nvjn                 1/1     Running   0          5m1s
kube-system          onos-operator-topo-7bcd58d88d-t6qx5                1/1     Running   0          5m1s
```

### Maintenance
To see the list of installed charts:
```bash
helm -n micro-onos ls
```

```
NAME          	NAMESPACE 	REVISION	UPDATED                                	STATUS  	CHART                  	APP VERSION
onos-umbrella	micro-onos	1       	2021-07-14 08:23:46.130263777 +0100 IST	deployed	onos-umbrella-1.1.11	v1.1.0     
```

To delete the deployment issue:
```bash
helm delete -n micro-onos onos-umbrella
```

## Deploy single services
**Alternatively** can deploy each service by itself. Please refer to each service's `deployment` file to get the exact command for each helm chart.
Example for [onos-topo](https://docs.onosproject.org/onos-topo/docs/deployment/).

```bash
helm -n micro-onos install onos-topo onosproject/onos-topo
```

> For individual services it is necessary to install CRDs first, as above:

## Developer workflow
Developers may want to run and deploy charts that have not yet been released. This
must be done from the checked out charts folder

> To use the latest version of an application without having to update the chart,
> an override like `--set image.tag=latest` can be used when individual charts.
> Alternatively when deploying the umbrella chart the override like
> `--set onos-topo.image.tag=latest`. The individual applications can be updated
> in to `kind` with the command `make kind`.

> Note that the source of the charts like `onosproject/onos-topo` will use the
> chart from the helm repository (cached locally), where as a source like `./onos-topo`
> will load the chart from the local folder. This is useful when editing charts.

### Check out the Helm charts
The helm charts need to be present on your PC. Run:
```bash
git clone https://github.com/onosproject/onos-helm-charts && cd onos-helm-charts
```

### Over-arching (umbrella) chart
Run the build of dependent charts to use the *local* `onos-umbrella` over-arching chart:
```bash
make deps
```

### Individual local charts
To deploy charts individually (from the `onos-helm-charts` directory) for example:
```bash
helm -n micro-onos install onos-topo ./onos-topo
```

[Docker]: https://docs.docker.com/get-docker/
[Kind]: https://kind.sigs.k8s.io
[Brew]: https://brew.sh/
[Helm]: https://helm.sh/
[Installing Helm]: https://helm.sh/docs/intro/install/
[Kubernetes]: https://kubernetes.io/
[ingress]: https://kubernetes.io/docs/concepts/services-networking/ingress/
[Rancher]: https://rancher.com/quick-start/
[Kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
