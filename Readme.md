# KongHQ API Gateway with JWT

## Steps to run

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
