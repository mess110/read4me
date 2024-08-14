# Cluster Setup

## kind

```bash
make create_local_cluster
make destroy_local_cluster
kubectl cluster-info --context kind-read4me-local
kubectl config view --minify --flatten --context=kind-read4me-local
```
