# Wordpress on k8s

## Create the wordpress deployment:
- `kubectl apply -f wp-app-deploy.yml`

## Create the mysql deployment:
- `kubectl apply -f wp-db-deploy.yml`

## Create the ClusterIP service:
When a pod inside a deployment dies, it gets recreated. This gives it a new ID and IP. That's horrible for us. If we add the pod's IP into the wordpress config, it can't connect after a restart. A restart deletes the pod and creates a new one. 

To get around this, we create a ClusterIP. The clusterIP will look for all pods with the selector label. In this case, "wp-db-deploy". From now on, we only have to add the clusterIP name to wordpress and it works. Kubernetes will take care of the rest.

- `kubectl apply -f wp-db-clusterip.yml`

## Create the Load Balancer service:
A ClusterIP service only works internally. If we want to access the wordpress pod, we have to make it available to everyone. Internally is not enough. To do this, we create a load balancer. This will expose one port for us to access. It will share all traffic with pods based on a label. In this case, "wp-app-deploy". 

- `kubectl apply -f wp-lb.yml`

## Disadvantages:
- When you recreate a pod, everything gets deleted. Including database entries. You restart the database, you end up with the installation window.
- The connection is not secure. Doesn't matter when local, but matters when no longer local.
