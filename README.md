# read4me

This is a test project which setups kind for local development and deploys to
GCP and AWS.

The goal is to have 1 control panel and 3 worker nodes.

The project should have 2 APIs, workers, logging and a frontend.

```markdown
fontend should be served on port 80 and 443 (later with lets encrypt)
/api
  /users-api
    GET /users - returns a json with users
    GET /users/{id} - returns a json with user
  /books-api
    GET /books - returns a json with books
```

frontend calls /api routes

The whole setup should be as automatic as possible.

There should be a kube dashboard

## Install

* [ ] docker
* [ ] kubectl
* [ ] kind
* [ ] terraform

## Steps (in random order)

1. [x] provision-kind-cluster
1. [x] provision-k8-dashboard
1. [x] provision-api-server
1. [ ] lets encrypt
1. [ ] redirect http to https
1. [ ] connect to DB
1. [ ] show UI
1. [ ] create worker nodes
1. [x] build generic API server
1. [ ] provision-gke-cluster
1. [ ] handle logging
1. [ ] terraform deploy
1. [ ] domain name
