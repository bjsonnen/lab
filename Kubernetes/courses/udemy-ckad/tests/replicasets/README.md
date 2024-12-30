# CKAD

# Chapter 2: Core Concepts

### ReplicaSets

1.) How many PODs exist on the system?

0

2.) How many ReplicaSets exist on the system? 

0

3.) How about now? How many ReplicaSets do you see?

1

4.) How many PODs are DESIRED in the new-replica-set?

4

5.) What is the image used to create the pods in the new-replica-set?

Busybox777

6.) How many PODs are READY in the new-replica-set?

0

7.) Delete any one of the 4 PODs. 

- `kubectl get pods`
- `kubectl delete pod new-replica-set-2878l`

8.) How many PODs exist now?

4

9.) Why are there still 4 PODs, even after you deleted one?

ReplicaSet ensured that desired number of PODs always run.

10.) Create a ReplicaSet using the replicaset-definition-1.yaml file located at /root/.

- `kubectl api-resources | grep replicaset`
- `kubectl explain replicaset | grep VERSION
- `Change `v1` to `apps/v1`
- `kubectl apply -f replicaset-definition-1.yaml`

11.) Fix the issue in the replicaset-definition-2.yaml file and create a ReplicaSet using it.

- `kubectl apply -f replicaset-definition-2.yaml`.
- Change under matchLabels `frontend` to `nginx`.
- Run `kubectl apply -f replicaset-definition-2.yaml`.

12.) Delete the two newly created ReplicaSets - replicaset-1 and replicaset-2

- `kubectl delete -f replicaset-definition-1.yaml`
- `kubectl delete -f replicaset-definition-2.yaml`

13.) Fix the original replica set new-replica-set to use the correct busybox image.

- `kubectl get replicaset.apps new-replica-set -o yaml > new-replica-set.yaml`
- Change image `busybox777` to `busybox`
- Run `kubectl apply -f new-replica-set.yaml`

14.) Scale the ReplicaSet to 5 PODs.

- `kubectl scale replicaset new-replica-set --replicas 5`

15.) Now scale the ReplicaSet down to 2 PODs.

- `kubectl scale replicaset new-replica-set --replicas 2`
