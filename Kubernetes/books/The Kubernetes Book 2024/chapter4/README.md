# The Kubernes Book 2024
Disclamer: The pods/ folder comes from the course materials of "The Kubernetes Book 2024". Downloaded from GitHub and split into chapter folders.

## Pods

Create the pod from a file:
- kubectl apply -f pods/pod.yml

Get all current pods:
- kubectl get pods

Get all current pods and watch for changes:
- kubectl get pods --watch

Get created pod and show desired and observed state:
- kubectl get pods hello-pod -o yaml

Get formatted info about a pod:
- kubectl describe pods hello-pod

Get logs for first container inside YAML:
- kubectl logs hello-pod

Get logs for one specific container inside YAML:
- kubectl logs hello-pod -- syncer

Execute command inside specific container:
- kubectl exec hello-pod -- ps
-> Pods and containers should be immutable and not be changed! Create a new pod if you want to change something.

Start a shell to use inside specific container:
- kubectl exec -it hello-pod -- sh
-> Pods and containers should be immutable and not be changed! Create a new pod if you want to change something.

You can try the immutability by trying to change any of these values:
-> Pod name
-> Container name
-> Container port
-> Resource requests and limits
- kubectl edit pod hello-pod

## Resources

The control plane node's controller is checking in with all worker nodes. It checks which worker node has enough resources available for your pod. To help your Kubernetes cluster do this, you can add requests. This is the minimum amout of cpu and memory the nodes has to provide.

```
resources:
  requests:
    cpu: 0.5
    memory: 256Mi
  limits:
    cpu: 1.0
    memory: 512Mi
```

Limits are the limits of resources for your pod. The worker node doesn't provide more resources than that to it.

## Init Pods

Run the init pod with:
- kubectl apply -f pods/initpod.yml

This pod will wait until a Kubernetes service pops up.

Create the Kubernetes service with:
- kubectl apply -f pods/initsvc.yml
