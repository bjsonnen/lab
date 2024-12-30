# Udemy CKAD

## Chapter 3: Configuration

### Service Accounts:

1.) How many Service Accounts exist in the default namespace?

- `kubectl get serviceaccounts`

2.) What is the secret token used by the default service account?

- `kubectl describe serviceaccounts default`

3.) We just deployed the Dashboard application. Inspect the deployment. What is the image used by the deployment?

- `kubectl get deployments.apps`
- `kubectl describe deployments.apps web-dashboard`

4.) Inspect the Dashboard Application POD and identify the Service Account mounted on it.

- `kubectl get pods`
- `kubectl describe pod web-dashboard-66d58fc7b8-5t4wt`

5.) The application needs a ServiceAccount with the Right permissions to be created to authenticate to Kubernetes. The default ServiceAccount has limited access. Create a new ServiceAccount named dashboard-sa.

- `kubectl create serviceaccount dashboard-sa`

6.) Create an authorization token for the newly created service account, copy the generated token and paste it into the token field of the UI. 

- `kubectl create token dashboard-sa`

7.) Edit the deployment to change ServiceAccount from default to dashboard-sa.

- Add a new field to the pod's spec: `serviceAccountName: dashboard-sa`
