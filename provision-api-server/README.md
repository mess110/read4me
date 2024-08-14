# README

```sh
/etc/hosts
127.0.0.1 books-api.local

curl http://books-api.local/api/users-service/users/1
curl http://books-api.local/api/books-service/books
```

## Logs

```sh
kubectl logs -l app=books-api-server-deployment --all-containers=true --prefix=true -f

kubectl get pods -l app=books-api-server-deployment
kubectl logs books-api-server-deployment-abc123
```
