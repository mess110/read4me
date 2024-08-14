#!/bin/bash

kubectl apply -f books-api-server-deployment.yaml
kubectl apply -f users-api-server-deployment.yaml
kubectl apply -f ingress-kind.yaml

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
