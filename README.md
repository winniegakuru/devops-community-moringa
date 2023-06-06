# Kubernetes 101 - Hands-on
This is a basic Kubernetes hands-on:
We will use an already setup Kubernetes cluster: https://killercoda.com/

## Scope
- Deploy a simple application using Kubernetes

### Namespace
`kubectl create namespace devops-community`

> list the namespace

`kubectl get namespace` or `k get ns`
> Delete the namespace

`kubectl delete ns devops-community`

### Pods
`kubectl run nginx --image=nginx -n devops-community`
`kubectl run nginx --image=nginx`

### Deployment
``
