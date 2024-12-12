# Kubernetes Foundations Course Notes
https://www.skool.com/kubecraft

## Introduction
Install using Rancher Desktop:
- https://rancherdesktop.io/ 

Add 
```
alias k='kubectl'
```
to `.zshrc` on MacOS.

Install 
```
brew install kubectl
brew install k9s
```

## Fundamentals:

### Pods vs Containers:
A pod can have multiple containers inside it.

### Commands

#### Pods
Get all pods:
- kubectl get pods

Run a pod with one container:
- kubectl run httpd-pods image=httpd

Create a pod config file:
- kubectl run httpd-pods --image=httpd --dry-run=client -o yaml > httpd-pods.yaml

Create a pod based on config file:
- kubectl create -f httpd-pods.yaml

Apply a pod config file:
- kubectl apply -f httpd-pods.yaml

Delete a pod config file:
- kubectl delete pods -f httpd-pods.yaml

#### Deployments
Show all deployments:
- kubectl get deployments.apps

Create a deployment config file:
- kubectl create deployment httpd-deploy --image=httpd --dry-run=client -o yaml > httpd-deploy.yaml

Create a deployment config files with replications:
- kubectl create deployment http-deploy-replicas --image=httpd --replicas=20 --dry-run=client -o yaml > httpd-deploy-replicas.yaml

Apply a deployment config file:
- kubectl apply -f httpd-deploy.yaml

Delete a deployment via config file:
- kubectl delete deployment -f httpd-deploy.yaml
