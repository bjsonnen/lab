# Udemy CKAD

## Chapter 2: Core Concepts

### Namespaces

1.) How many Namespaces exist on the system?

10

2.) How many pods exist in the research namespace?

2

3.) Create a POD in the finance namespace.

- `kubectl run redis --namespace finance --image redis`

4.) Which namespace has the blue pod in it?

- `kubectl get pods -A | grep blue`
