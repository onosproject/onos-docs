# Deploying ONOS classic with HELM

Like the µONOS services, [ONOS](https://wiki.onosproject.org/) classic (Java-Karaf
based SDN controller) can also be deployed to a Kubernetes cluster.

## Deployment options

### Local Deployment Setup with Kind
See [deploy with helm](./deploy_with_helm.md#Local Deployment Setup with Kind)
for a general descriptions for Kind and Helm installation.

## Bare metal deployment
For High Availability features the deployment should have at least 3 nodes in the
cluster. Because "PodAffinity" is enabled by default - the onos-classic service
will not start without 3 nodes with the default settings.

## Prerequisites
Follow the same [Prerequisites](./deploy_with_helm.md#Prerequisites) as in the
µONOS deployment (except the CORD repo is not necessary).

## Deploy onos-classic
> onos-classic can be deployed alongside the µONOS
> [umbrella chart](./deploy_with_helm.md#Deploy the SD-RAN set of onos services)
> or any combination of individual charts.

### Deploy the full HA scenario
> This requires at least 3 nodes in the K8S cluster.
```bash
helm -n micro-onos install onos-classic onosproject/onos-classic
```

### Deploy on a single node Kind installation
This is the simplest possible deployment with a local Atomix installation and
a single replica.

```bash
helm -n micro-onos install onos-classic onosproject/onos-classic --set atomix.replicas=0 --set replicas=1
```

> Kind will pull several images from Docker in to its own repo before it starts
> the cluster. This can be time consuming over a limited connection. It may help
> performance by doing `docker pull` on each of the images listed below just once
> in to your local registry, and then after Kind starts up loading these in to Kind
> with:
```bash
kind load docker-image ubuntu:16.04
kind load docker-image onosproject/onos:2.2.2
kind load docker-image tutum/dnsutils:latest
kind load docker-image atomix/atomix:3.1.0
```

## Accessing the ONOS CLI
The CLI can be accessed through SSH (or if it is available the classic `onos` tool).

On a bare metal cluster:
```bash
ssh -p 8101 onos@<cluster ip>
```
> the password as usual is `rocks`. Access can also be got by `karaf/karaf`.

Alternatively using the `onos` script from `~/onos/tools/test/bin/onos`
```bash
onos onos@<cluster ip>
```

On **Kind**, because NodePorts are not exposed, it is necessary to set up a `port-forward`
to the `onos-classic` pod in the cluster, before accessing the CLI.
```bash
kubectl -n micro-onos port-forward $(kubectl -n micro-onos get pods -l app=onos-classic-onos-classic -o name) 8101
```

The CLI is then available at `localhost`.
```bash
ssh -p 8101 onos@localhost
```

## Accessing the ONOS GUI
The GUI can be accessed at port 8181 of the cluster like
**http://<cluster ip>:8181/onos/ui**

> With Kind it is necessary to set up port-forwarding to port 8181 and the GUI
> will be available at **http://localhost:8181/onos/ui**

## Controlling the default set of applications
By default only the following apps are installed:
```
*  28 org.onosproject.optical-model        2.2.2    Optical Network Model
*  29 org.onosproject.openflow-base        2.2.2    OpenFlow Base Provider
*  34 org.onosproject.drivers              2.2.2    Default Drivers
*  83 org.onosproject.gui2                 2.2.2    ONOS GUI2
```

To load additional apps on startup, add them like:
```bash
helm -n micro-onos install onos-classic onosproject/onos-classic \
--set atomix.replicas=0 --set replicas=1 \
--set apps[1]=org.onosproject.drivers.bmv2 \
--set apps[2]=org.onosproject.lldpprovider \
--set apps[3]=org.onosproject.hostprovider
```
> i.e. the set used in the [NG-SDN Tutorial](https://github.com/opennetworkinglab/ngsdn-tutorial/blob/master/EXERCISE-3.md#2-start-onos)
