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

* First, you will need to install [Docker] to build and deploy an image locally.

* Second, install [Kind].
> **Kind v0.7.0 at least** is required, which provides the **K8S API v1.17**

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
helm repo add onos https://charts.onosproject.org
```

### Update the local cache with charts from these repos:
```bash
helm repo update
```

### Inspect the chart versions and app versions
To see the list of the latest chart versions and app versions, use the "search" command.
```bash
helm search repo onosproject
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
NAME                                                         READY   STATUS    RESTARTS   AGE
atomix-controller-kubernetes-controller-66965c74b-v6d9m      1/1     Running   0          18m
cache-controller-cache-storage-controller-5666d67c4d-kv2hh   1/1     Running   0          30s
raft-controller-raft-storage-controller-78c7999cf5-zwn8b     1/1     Running   0          99s
```

## Deploy the SD-RAN set of onos services
A complete set of onos services can be deployed with just the over-arching (umbrella)
[`sd-ran` chart](https://github.com/onosproject/onos-helm-charts/tree/master/sd-ran). 

Run the install:
```bash
helm -n micro-onos install sd-ran onosproject/sd-ran \
--set global.store.controller=atomix-controller-kubernetes-controller:5679
```

this will deploy `onos-ric`, `onos-ric-ho`, `onos-ric-mlb`, `ran-simulator`,
`onos-topo`, `onos-cli`, `onos-gui`, and `onos-config` (but not `onos-classic` as
it's not needed for SD-RAN - see [Deploying ONOS classic with HELM](./deploy_onos_classic.md)).

To monitor the startup of the pods use `kubectl` like:
```bash
kubectl -n micro-onos get pods -w
```

giving a list like:
```
NAME                                                         READY   STATUS    RESTARTS   AGE
atomix-controller-kubernetes-controller-66965c74b-v6d9m      1/1     Running   0          28m
cache-controller-cache-storage-controller-5666d67c4d-kv2hh   1/1     Running   0          10m
onos-cache-1-57484c95df-th9tg                                1/1     Running   0          6m25s
onos-cli-677dcfccb9-tn8jh                                    1/1     Running   0          6m26s
onos-consensus-1-0                                           1/1     Running   0          6m26s
onos-gui-6f498dcdd4-jbnvq                                    2/2     Running   0          6m26s
onos-ric-bcc5f9668-6844p                                     1/1     Running   1          6m26s
onos-ric-ho-8566597f4b-8vl7v                                 1/1     Running   0          6m26s
onos-ric-mlb-64d6ccd64-ppb2m                                 1/1     Running   0          6m25s
onos-topo-545497b866-ccpvf                                   1/1     Running   1          6m26s
raft-controller-raft-storage-controller-78c7999cf5-zwn8b     1/1     Running   0          12m
ran-simulator-647d587c6f-vgwzr                               1/1     Running   0          6m26s
sd-ran-grafana-7dfdb45bc5-xm9xf                              2/2     Running   0          6m26s
sd-ran-prometheus-alertmanager-dcc498556-p29pp               2/2     Running   0          6m26s
sd-ran-prometheus-kube-state-metrics-5566cb77b6-5snnb        1/1     Running   0          6m26s
sd-ran-prometheus-node-exporter-pz77d                        1/1     Running   0          6m26s
sd-ran-prometheus-pushgateway-d468f4798-mj9lv                1/1     Running   0          6m26s
sd-ran-prometheus-server-75f99b45f8-g27rs                    2/2     Running   0          6m26s
``` 

## Deploy only **onos-config** and related services
**Alternatively** to install a cluster where you are not interested in SD-RAN and only want onos-config, you could run
```bash
helm -n micro-onos install sd-ran onosproject/sd-ran \
     --set global.store.controller=atomix-controller-kubernetes-controller:5679 \
     --set import.onos-config.enabled=true \
     --set import.onos-ric.enabled=false \
     --set import.onos-ric-ho.enabled=false \
     --set import.onos-ric-mlb.enabled=false \
     --set import.ran-simulator.enabled=false \
     --set import.nem-monitoring.enabled=false
```

### Maintenance
To see the list of installed charts:
```bash
helm -n micro-onos ls
```

```
NAME             	NAMESPACE 	REVISION	UPDATED                                	STATUS  	CHART                         	APP VERSION  
atomix-controller	micro-onos	1       	2020-04-27 08:34:11.571681508 +0100 IST	deployed	kubernetes-controller-0.4.2   	v0.3.0-beta.1
cache-controller 	micro-onos	1       	2020-04-27 08:52:14.329737405 +0100 IST	deployed	cache-storage-controller-0.3.2	v0.2.0       
raft-controller  	micro-onos	1       	2020-04-27 08:51:05.369054085 +0100 IST	deployed	raft-storage-controller-0.3.2 	v0.2.0       
sd-ran           	micro-onos	1       	2020-04-27 08:56:38.538404922 +0100 IST	deployed	sd-ran-0.0.2                  	v0.6.0
```

To delete the deployment issue:
```bash
helm delete -n micro-onos sd-ran
```

## Deploy single services services
**Alternatively** can deploy each service by itself. Please refer to each service's `deployment` file to get the exact command for each helm chart.
Example for [onos-topo](https://docs.onosproject.org/onos-topo/docs/deployment/).

```bash
helm -n micro-onos install onos-topo onosproject/onos-topo \
--set store.controller=atomix-controller-kubernetes-controller:5679
```

> For individual services it is necessary to install Atomix first, as above:

## Developer workflow
Developers may want to run and deploy charts that have not yet been released.

### Check out the Helm charts
The helm charts need to be present on your PC. Run:
```bash
git clone https://github.com/onosproject/onos-helm-charts && cd onos-helm-charts
```

### Over-arching (umbrella) chart 
Run the build of dependent charts to use the local `sd-ran` over-arching chart:
```bash
helm dep build sd-ran
```

If you make changes to one of the charts and want to re-deploy, please first issue:
```bash
helm dependency update sd-ran
```

### Individual charts
To deploy charts individually (from the `onos-helm-charts` directory) for example:
```bash
helm -n micro-onos install onos-topo onos-topo --set store.controller=atomix-controller-kubernetes-controller:5679
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
