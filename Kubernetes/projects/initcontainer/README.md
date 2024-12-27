# Kubernetes Initcontainer

For a webserver for my homelab, I was looking for a solution to pull data into a volume before it starts. I found on the k8s docs that I should use an initcontainer.

Link: https://kubernetes.io/docs/concepts/storage/volumes/#gitrepo

## Usage

Build thd docker image:
- `docker build -t init-git-curl:1.0`

Apply the storage.yaml:
- `kubectl apply -f storage.yaml`

Apply the deployment.yaml:
- `kubectl apply -f storage.yaml`

Use port-forwarding for this example.

First, list out the containers:
- `kubectl get pods`

Then, pick the running php pod and run:
- `kubectl port-forward php-54b66559c-xzlbs 8080:80`

Now open your browser and point to `localhost:8080`

## Info

The php container will only start if the init container finished. If it puts out an error, k8s will restart the pod.
