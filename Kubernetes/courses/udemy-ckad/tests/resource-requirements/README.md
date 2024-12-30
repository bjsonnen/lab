# Udemy CKAD

## Chapter 3: Configuration

### Resource Requirements

1.) A pod called rabbit is deployed. Identify the CPU requirements set on the Pod

- `kubectl get pods`
- `kubectl get pod rabit -o yaml`

2.) Delete the rabbit Pod.

- `kubectl delete pod rabbit`

3.) Another pod called elephant has been deployed in the default namespace. It fails to get to a running state. Inspect this pod and identify the Reason why it is not running.

- `kubectl get pods`
- `kubectl describe pod elephant`

4.) The status OOMKilled indicates that it is failing because the pod ran out of memory. Identify the memory limit set on the POD.

- `kubectl get pod elephant -o yaml`
