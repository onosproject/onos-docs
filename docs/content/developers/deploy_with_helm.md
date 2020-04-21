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
* deployment of Atomix controller in the namespace.

The individual components in the [umbrella chart](https://github.com/onosproject/onos-helm-charts/tree/master/sd-ran) are:

*   onos-topo:
*   onos-config:
*   onos-cli:
*   onos-gui:
*   onos-ric:
*   onos-ric-ho:
*   onos-ric-mlb:
*   ran-simulator:
*   nem-monitoring:

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

* First, you will need Docker to build and deploy an image locally. Install Docker following the 
[Docker installation instructions](https://docs.docker.com/v17.12/install/).

* Second, install [Kind] following the [instructions](https://kind.sigs.k8s.io).
> **Kind v0.7.0 at least** is required, which provides the **K8S API v1.17**

* Third, install [Helm]. On OSX, this Helm can be installed using [Brew]:

```bash
brew install kubernetes-helm
```
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

### Update the local cache with charts from these repos:
```bash
helm repo update
```

### Check out the Helm charts
The helm charts need to be present on your PC. Run:
```bash
git clone https://github.com/onosproject/onos-helm-charts && cd onos-helm-charts
```

### Configure the micro-onos namespace
The various `onos` services can be deployed to any `namespace`.
> A "namespace" partitions the cluster in to independent islands.
> For consistency between documentation we are going to use the `micro-onos` here.

To create the `micro-onos` namespace run:
```bash
kubectl create namespace micro-onos
```

## Deploy the SD-RAN set of onos services
A complete set of onos services can be deployed with just the
[`sd-ran` chart](https://github.com/onosproject/onos-helm-charts/tree/master/sd-ran). 

Run the build of dependent charts:
```bash
helm dep build sd-ran
```

Finally, run the install:
```bash
helm -n micro-onos install sd-ran sd-ran \
    --set onos-ric.store.controller=atomix-controller.micro-onos.svc.cluster.local:5679
```

this will deploy `onos-ric`, `onos-ric-ho`, `onos-ric-mlb`, `ran-simulator`, `onos-topo`, `onos-cli` and `onos-gui`,
but not `onos-config` (as it's not currently needed for SD-RAN).

It should give an output like:
```bash
NAME: sd-ran
LAST DEPLOYED: Tue Apr  7 15:56:45 2020
NAMESPACE: micro-onos
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Thank you for installing sd-ran Helm chart.

Your release is named sd-ran in namespace micro-onos.
See https://docs.onosproject.org/developers/deploy_with_helm/

To learn more about the release, try:
  $ helm -n micro-onos status sd-ran
  $ helm -n micro-onos get all sd-ran
  $ watch kubectl -n micro-onos get pods

You can attach to:
* Onos CLI pod with
$ kubectl -n micro-onos exec -it $(kubectl -n micro-onos get pods -l type=cli -o name) -- /bin/sh
* SD-RAN GUI at http://<server_IP>:31180
* Prometheus at http://<server_IP>:31301/targets
* Grafana at http://<server_IP>:31300 (admin/strongpassword)

If you are using Kind as a Kubernetes server, you will have to use a "port-forward" to access the GUI, Grafana and Prometheus e.g.
$ kubectl -n micro-onos port-forward $(kubectl -n micro-onos get pods -l type=gui -o name) 8182:80
and then access the GUI at
* http://localhost:8182
``` 

## Deploy only **onos-config** and related services
Alternatively to install a cluster where you are not interested in SD-RAN and only want onos-config, you could run
```bash
helm -n micro-onos install sd-ran sd-ran \
     --set onos-ric.store.controller=atomix-controller.micro-onos.svc.cluster.local:5679 \
     --set onos-ric.store.raft.backend.image=atomix/local-replica:latest \
     --set import.onos-config.enabled=true \
     --set import.onos-ric.enabled=false \
     --set import.onos-ric-ho.enabled=false \
     --set import.onos-ric-mlb.enabled=false \
     --set import.ran-simulator.enabled=false \
     --set import.nem-monitoring.enabled=false
```

### Maintenance
To delete the deployment issue:
```bash
helm delete -n micro-onos sd-ran
```

If you make changes to one of the charts and want to re-deploy, please first issue
```bash
helm dependency update sd-ran
```
## Deploy single services services

You can also deploy each service by itself. Please refer to each service's `deployment` file to get the exact command for each helm chart.
Example for [onos-topo](https://docs.onosproject.org/onos-topo/docs/deployment/).
> For individual services it is necessary to install Atomix first, as below:

### Deploy Atomix Controller

The various `onos` services leverage Atomix as the distributed store for HA, scale and redundancy.
The first thing that needs to be deployed in any `onos` deployment is the Atomix go controller.
To deploy the Atomix controller do:

```bash
helm -n micro-onos install atomix-controller atomix/kubernetes-controller --set scope=Namespace
helm -n micro-onos install cache-controller atomix/cache-storage-controller --set scope=Namespace
helm -n micro-onos install raft-controller atomix/raft-storage-controller --set scope=Namespace
```

If you watch the `pods` you should now see:
```bash
$ kubectl -n micro-onos get pods
NAME                                 READY   STATUS    RESTARTS   AGE
atomix-controller-68dc7d7c79-dxztq   1/1     Running   0          5m35s
cache-controller-964794d57-zw4cq     1/1     Running   0          4m7s
raft-controller-6dd86cfd54-nlzzt     1/1     Running   0          2m50s
```


[Kind]: https://kind.sigs.k8s.io
[Brew]: https://brew.sh/
[Helm]: https://helm.sh/
[Kubernetes]: https://kubernetes.io/
[ingress]: https://kubernetes.io/docs/concepts/services-networking/ingress/
[Rancher]: https://rancher.com/quick-start/
[Kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
