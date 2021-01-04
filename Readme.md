# KongHQ API Gateway with JWT

## Docker Compose

### Steps to Run

1. Docker Compose Up

```
docker-compose -f "docker-compose.yml" up -d --build
```

2. Create User by sending a POST Request to http://localhost:8000/auth/signup

```json
{
  "username": "shivam",
  "password": "shivam",
  "details": "secret"
}
```

3. Login by sending a POST Request to http://localhost:8000/auth/signin

```json
{
  "username": "shivam",
  "password": "shivam"
}
```

4. Retrieve the token

5. Access the secured route using token at http://localhost:8000/secure/restricted

## Kubernetes

### Steps to Run

1. Ensure your kubernetes cluster is up and running.(I have assumed you have a Kops k8s.local cluster)

2. Apply the configuration files

```
kubectl create -f kubernetes/
```

3. Wait for the pods to be up and running. All pods and deployment are in namespace kong.

4. Configure KongHQ using decK.(Run this in folder containing kong.yaml)

```
deck sync
```

5. Get external ip of KongHQ API Gateway(kong-proxy) using

```
kubectl get svc -n kong
```

6. Access all addresses similarly

```
http://localhost:8000 -> http://amazon-ip:80
```
