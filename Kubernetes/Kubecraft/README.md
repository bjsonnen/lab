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

## Recommendations:
- Get used to using the CLI
- Generate files for Kubernetes and deploy these. This makes it easy for you to get used to GitOps. Don't log into servers and configure them on your own. Write scripts and make them auto deploy. You can also run automated tests against them. This makes it easy to check if everything works.
- Get used to using the Kubernetes help (man) pages. During Kubernetes exams, you can quickly review the built-in docs. This saves you a lot of time.

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

Expose a port of a deployment:
- kubectl expose deployment httpd-deploy --port 8080

## Networking

- Networking happens on a pod level.
- All containers inside a pod can communicate with each other via localhost.
- When using default settings, all pods can communicate with each other via ips.
- Each pod has an own ip address. To see the IP, use `kubectl get pods -o wide`.
