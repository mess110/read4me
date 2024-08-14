#!/bin/bash

kubectl apply -f books-api-server-deployment.yaml
kubectl apply -f users-api-server-deployment.yaml
kubectl apply -f ingress-gcp.yaml

# Point your domain name to this IP address using your DNS provider.
# kubectl get ingress
