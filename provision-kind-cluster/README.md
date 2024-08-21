# Cluster Setup

## kind

```bash
./create_cluster.sh
./destroy_cluster.sh
kubectl cluster-info --context kind-read4me-local
kubectl config view --minify --flatten --context=kind-read4me-local
```
