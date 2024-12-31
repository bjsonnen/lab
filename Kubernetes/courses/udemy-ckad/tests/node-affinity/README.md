# Udemy CKAD

## Chapter 3: Configuration

### Node Affinity

1.) How many Labels exist on node node01?

- `kubectl describe nodes node01`

2.) Which nodes can the pods for the blue deployment be placed on?

- `kubectl describe nodes controlplane`
- `kubectl describe nodes node01`

Look for the `Taints:` configuration.
