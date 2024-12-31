# Udemy CKAD

## Chapter 3: Configuration

### Taints and Toleration

1.) How many nodes exist on the system?

- `kubectl get nodes`

2.) Do any taints exist on node01 node?

- `kubectl describe nodes node01 | grep Taint`

3.) Create a taint on node01 with key of spray, value of mortein and effect of NoSchedule

- `kubectl taint nodes node01 spray=mortein:NoSchedule`

4.) Create a new pod with the nginx image and pod name as mosquito.

- `kubectl run mosquito --image nginx`

5.) What is the state of the POD?

- `kubectl get pods`

6.) Create another pod named bee with the nginx image, which has a toleration set to the taint mortein.

- `kubectl run bee --image nginx --dry-run=client -o yaml > pod.yaml`
- Add to spec (not containers): 
```
tolerations:
  - key: spray
    value: mortein
    effect: NoSchedule
    operator: Equal
```
- `kubectl apply -f pod.yaml`

7.) Notice the bee pod was scheduled on node node01 despite the taint.

- `kubectl get pods -o wide`

8.) Do you see any taints on controlplane node?

- `kubectl describe nodes controlplane | grep Taint`

9.) Remove the taint on controlplane, which currently has the taint effect of NoSchedule.

- `kubectl taint nodes controlplane node-role.kubernetes.io/control-plane:NoSchedule-`

10.) What is the state of the pod mosquito now?

- `kubectl get pods`

11.) Which node is the POD mosquito on now?

- `kubectl get pods -o wide`
