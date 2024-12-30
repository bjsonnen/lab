# Chapter 2: Core Concepts

## Practice Test - Pods

1.) How many pods exist inside the default namespace?

0

2.) Create a new pod with the nginx image.

`kubectl run nginx --image=nginx`

3.) How many pods are created now?

4

4.) What images are they running?

Busybox

5.) Which nodes are these pods placed on?

Control plane

6.) How many containers are part of the pod webapp?

2

7.) What images are used in the new webapp pod?

nginx & agentx

8.) What is the state of the container agentx in the pod webapp?

Waiting

9.) Why do you think the container agentx in pod webapp is in error?

Image does not exist.

10.) What does the READY column in the output of the kubectl get pods command indicate?

Total containers / Running containers

11.) Delete the pod webapp.

`kubectl delete pod webapp`

12.) Create a new pod with the name redis and the image redis123.

`kubectl run redis --image=redis123`

13.) Now change the image on this pod to redis.

`kubectl edit pod redis`

## Notes

To edit a pod, make sure you export the existing config to yaml with:
- `kubectl get pod <name> -o yaml > pod.yaml`

Then make your changes inside the .yaml file. Once done, apply it:
- `kubectl apply -f pod.yaml`

This way you use the declarative way of k8s. 
Don't use the imperative!
