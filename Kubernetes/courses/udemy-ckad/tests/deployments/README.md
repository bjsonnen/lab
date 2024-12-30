# Udemy CKAD

## Chapter 2: Core Concepts

### Deployments

1.) How many PODs exist on the system?

0

2.) How many ReplicaSets exist on the system?

0

3.) How many Deployments exist on the system?

0

4.) How many Deployments exist on the system now?

1

5.) How many ReplicaSets exist on the system now?

1

6.) How many PODs exist on the system now? 

4

7.) Out of all the existing PODs, how many are ready?

0

8.) What is the image used to create the pods in the new deployment?

busybox888

9.) Why do you think the deployment is not ready?

The image BUSYBOX888 doesn't exist.

10.) Create a new Deployment using the deployment-definition-1.yaml file located at /root/.

- `kubectl apply -f deployment-definition-1.yaml`
- Change `kind: deployment` to `kind: Deployment`
- `kubectl apply -f deployment-definition-1.yaml`

11.) Create a new Deployment with the below attributes using your own deployment definition file.

- `kubectl create deployment httpd-frontend --replicas 3 --image httpd:2.4-alpine`
