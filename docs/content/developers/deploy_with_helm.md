# Deploying onos micro-services with HELM

One of the goals of the micro-onos project is to provide simple deployment options
that integrate with modern technologies. Deployment configurations can be found in
the `onos-helm-charts` [repository](https://github.com/onosproject/onos-helm-charts). Each `onos` service has a directory containing its chart.
As an example the `onos-config` chart is in `onos-helm-charts/onos-config`. 

## Deploying on Kubernetes with Helm

[Helm] is a package manager for [Kubernetes] that allows projects to provide a
collection of templates for all the resources needed to deploy on k8s. ONOS Config
provides a Helm chart for deploying a cluster for development and testing. In the
future, this chart will be extended for production use.

### Resources

The Helm chart provides resources for deploying the config service and accessing
it over the network, both inside and outside the k8s cluster:

*  `Deployment` - Provides a template for ONOS Config pods
*  `ConfigMap` - Provides test configurations for the application
*  `Service` - Exposes ONOS Config to other applications on the network
*  `Secret` - Provides TLS certificates for end-to-end encryption
*  `Ingress` - Optionally provides support for external load balancing

### Local Deployment Setup

To deploy the Helm chart locally: 

* First, you will need Docker to build and deploy an image locally. Install Docker following the 
[Docker installation instructions](https://docs.docker.com/v17.12/install/).

* Second, install [Kind] following the [instructions](https://kind.sigs.k8s.io).  

* Third, install [Helm]. On OSX, this Helm can be installed using [Brew]:

```bash
brew install kubernetes-helm
```
* Once Kind has been installed, start it with 

```bash
kind create cluster. 
```

* Once Kind has started, set your environment to the Kubernetes cluster:
   
```bash
export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
```
   
### Deploy Atomix Controller

The various `onos` services leverage Atomix as the distributed store for HA, scale and redundancy.
The first thing that needs to be deployed in any `onos` deployment is the Atomix go controller.
To deploy the Atomix controller do:
```bash
kubectl create -f https://raw.githubusercontent.com/atomix/atomix-k8s-controller/master/deploy/atomix-controller.yaml
```
The correct return output looks like this: 
```bash
customresourcedefinition.apiextensions.k8s.io/partitionsets.k8s.atomix.io created
customresourcedefinition.apiextensions.k8s.io/partitions.k8s.atomix.io created
clusterrole.rbac.authorization.k8s.io/atomix-controller created
clusterrolebinding.rbac.authorization.k8s.io/atomix-controller created
serviceaccount/atomix-controller created
deployment.apps/atomix-controller created
service/atomix-controller created
```
If you watch the `pods` you should now see:
```bash
kubectl get pods --all-namespaces

NAMESPACE         NAME                                         READY   STATUS    RESTARTS   AGE
kube-system       atomix-controller-b579b9f48-lgvxf            1/1     Running   0          152m
```
### Configure the micro-onos namespace
The various `onos` services can be deployed to any namespace. 
For consistency between documentation we are going to use the `micro-onos` one.
To create the `micro-onos` namespace do the following:
```bash
kubectl create namespace micro-onos
```

### Deploy the whole set of onos service
Once you have exported the `KUBECONFIG` flag you can start deploy `onos` services through helm charts
A complete set of onos services can be deployed with just the [`onos` chart](https://github.com/onosproject/onos-helm-charts/tree/master/onos). 
In the root directory of the `onos-helm-chart` repository issue
```bash
helm install -n micro-onos onos onos
```

this will deploy `onos-config`, `onos-topo`, `onos-cli` and `onos-gui`. 

To delete the deployment issue:
```bash
helm delete -n micro-onos onos
```

If you make changes to one of the charts and want to re-deploy, please first issue
```bash
helm dependency update onos
```
### Deploy single services services

You can also deploy each service by itself. Please refer to each service's `deployment` file to get the exact command for each helm chart.
Example for [onos-config](https://docs.onosproject.org/onos-config/docs/deployment/).

[Kind]: https://kind.sigs.k8s.io
[Brew]: https://brew.sh/
[Helm]: https://helm.sh/
[Kubernetes]: https://kubernetes.io/
[ingress]: https://kubernetes.io/docs/concepts/services-networking/ingress/
