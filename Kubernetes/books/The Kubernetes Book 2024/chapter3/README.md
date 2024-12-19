# The Kubernetes Book 2024

## Chapter 3

### Switching contexts inside kubectl

You can install Docker Desktop with Kubernetes or Rancher Desktop.
If you have many Kubernetes clusters installed, you may want to switch.

See current context:
- kubectl config current-context

List all available contexts:
- kubectl config all-contexts

Switch to different context:
- kubectl config use-context (docker-desktop / rancher desktop)
