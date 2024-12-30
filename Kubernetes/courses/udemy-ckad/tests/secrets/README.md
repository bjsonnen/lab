# Udemy CKAD

## Chapter 3: Configuration

### Secrets

1.) How many Secrets exist on the system?

- `kubectl get secrets`

2.) How many secrets are defined in the dashboard-token secret?

- `kubectl get secrets`

3.) What is the type of the dashboard-token secret? 

- `kubectl get secrets`

4.) Which of the following is not a secret data defined in dashboard-token secret?

- `kubectl describe secrets dashboard-token`

5.) The reason the application is failed is because we have not created the secrets yet. Create a new secret named db-secret with the data given below.

- `kubectl create secret generic db-secret --from-literal=DB_Host=sql01 --from-literal=DB_User=root --from-literal=DB_Password=password123 --dry-run=client -o yaml > secrets.yaml`
- `kubectl apply -f secrets.yaml`
