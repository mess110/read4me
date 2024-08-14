# read4me

This is a test project which setups kind for local development and deploys to
GCP and AWS.

The goal is to have 1 control panel and 3 worker nodes.

The project should have 2 APIs, workers, logging and a frontend.

fontend should be served on port 80 and 443 (later with lets encrypt)
/api
  /users-api
    GET /users - returns a json with users
    GET /users/{id} - returns a json with user
  /books-api
    GET /books - returns a json with books

frontend calls /api routes

The whole setup should be as automatic as possible.
There should be a kube dashboard

## Install

* [ ] kubectl
* [ ] kind
* [ ] terraform

## Steps

1. [ ] provision-local-cluster
2. [ ] provision-k8-dashboard
