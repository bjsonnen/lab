# Udemy CKAD Course

## Chapter 3: Configuration

### ConfigMaps

1.) How many PODs exist on the system?

- `kubectl get pods`

2.) What is the environment variable name set on the container in the pod?

- `kubectl describe pod webapp-color`

3.) Update the environment variable on the POD to display a green background. 

- `kubectl get pod webapp-color -o yaml > webapp-color.yaml`
- Change `value` for `APP_COLOR` to `green`
- `kubectl apply -f webapp-color.yaml`

4.) Identify the database host from the config map db-config.

- `kubectl describe configmaps db-config`
