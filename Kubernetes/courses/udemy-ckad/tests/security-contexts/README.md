# Udemy CKAD

## Chatper 3: Configuration

### Security Context

1.) Edit the pod ubuntu-sleeper to run the sleep process with user ID 1010.

- `kubectl get pod ubuntu-sleeper -o yaml > ubuntu-sleeper.yaml`
- Add `runAsUser: 1010` to `securizyContext`
- `kubectl delete pod ubuntu-sleeper`
- `kubectl apply -f ubuntu-sleeper.yaml`

2.) Update pod ubuntu-sleeper to run as Root user and with the SYS_TIME capability.

- Add `capabilities:` to `securityContext:`
- Add `add: ["SYS_TIME"]` to `capabilities`.
- Run `kubectl delete pod ubuntu-sleeper`
- Run `kubectl apply -f ubuntu-sleeper.yaml`

Make sure you add the capabilities to the container, not the entire pod!
