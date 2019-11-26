# Deploying onos micro-services with HELM

One of the goals of the micro-onos project is to provide simple deployment options
that integrate with modern technologies. Deployment configurations can be found in
the `/deployments/helm` folder in every repository that posses the Helm charts. 
For example see the `onos-config/deplyments/helm` folder.

## Deploying on Kubernetes with Helm

[Helm] is a package manager for [Kubernetes] that allows projects to provide a
collection of templates for all the resources needed to deploy on k8s. ONOS Config
provides a Helm chart for deploying a cluster for development and testing. In the
future, this chart will be extended for production use.

### Resources

The Helm chart provides resources for deploying the config service and accessing
it over the network, both inside and outside the k8s cluster:
* `Deployment` - Provides a template for ONOS Config pods
* `ConfigMap` - Provides test configurations for the application
* `Service` - Exposes ONOS Config to other applications on the network
* `Secret` - Provides TLS certificates for end-to-end encryption
* `Ingress` - Optionally provides support for external load balancing

### Local Deployment Setup

To deploy the Helm chart locally first install [Kind] following the [instructions](https://kind.sigs.k8s.io).  
[Helm] is also required. On OSX, this Helm can be installed using [Brew]:
```bash
> go get sigs.k8s.io/kind
> brew install kubernetes-helm
```

You will also need Docker to build and deploy an image locally.
* Docker [installation instructions](https://docs.docker.com/v17.12/install/)


Once Kind has been installed, start it with `kind create cluster`. 

Once Kind has started, set your  environment to the Kubernetes cluster:

```bash
> export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
```
### Deploy Atomix Controller

The various `onos` services leverage Atomix as the distributed store for HA, scale and redundancy.
The first thing that needs to be deployed in any `onos` deployment is the Atomix go controller.
Fetch the `atomix-controller.yaml` file:
```bash
curl -o atomix-controller.yaml https://raw.githubusercontent.com/atomix/atomix-k8s-controller/master/deploy/atomix-controller.yaml
```
if you want to deploy the cluster in the non default namespace please edit `atomix-controller.yaml`, replacing `default`
with your namespace.

To then deploy the atomix controller issue:
```bash
kubectl create -f atomix-controller.yaml
```
with the correct return output being 
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

NAMESPACE     NAME                                         READY   STATUS    RESTARTS   AGE
default       atomix-controller-b579b9f48-lgvxf            1/1     Running   0          152m
```

### Deploy services

Once you have exported the `KUBECONFIG` flag you can start deploy `onos` services through helm charts.
Please refer to each service's `deployment` file to get the exact command for each helm chart.
Example for [onos-config](https://docs.onosproject.org/onos-config/docs/deployment/).

[Kind]: https://kind.sigs.k8s.io
[Brew]: https://brew.sh/
[Helm]: https://helm.sh/
[Kubernetes]: https://kubernetes.io/
[ingress]: https://kubernetes.io/docs/concepts/services-networking/ingress/