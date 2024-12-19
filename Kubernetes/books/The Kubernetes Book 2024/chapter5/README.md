# The Kubernetes Book 2024

Add a namespace to all kubectl commands:
- kubectl config set-context --current --namespace shield

To unset a namespace:
- kubectl config set-context --current --namespace default
- kubectl config set-context --current --namespace=""

You can execute namespace command declaratively and imperatively:

### Imperatively

Add the `-n` or `--namespace` flag to all commands.

### Declaratively

Add the namespace in objects to your YAML file.
