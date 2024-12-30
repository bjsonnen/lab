# Udemy CKAD

## Chapter 3: Configuration

### Commands and Arguments

1.) How many PODs exist on the system?

- `kubectl get pods`

2.) What is the command used to run the pod ubuntu-sleeper?

- `kubectl describe pod ubuntu-sleeper`

3.) Create a pod with the ubuntu image to run a container to sleep for 5000 seconds. Modify the file ubuntu-sleeper-2.yaml.

- Add `command: ["sleep"]` and `args: ["5000"]` to the YAML file.

4.) Create a pod using the file named ubuntu-sleeper-3.yaml. There is something wrong with it. Try to fix it! 

- `kubectl apply -f ubuntu-sleeper-3.yaml`

File:
```
command:
  - "sleep"
  - 1200
```

Change into:
```
command:
  - "sleep"
  - "1200"
```

5.) Inspect the file Dockerfile given at /root/webapp-color directory. What command is run at container startup?

- `cd webapp-color/`
- `cat Dockerfile`

6.) Inspect the file Dockerfile2 given at /root/webapp-color directory. What command is run at container startup?

- `cat Dockerfile2`

7.) Inspect the two files under directory webapp-color-2. What command is run at container startup?

- `cd ../webapp-color-2`
- `cat Dockerfile`
