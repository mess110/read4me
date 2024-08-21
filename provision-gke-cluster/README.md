# README

```sh
gcloud init
gcloud auth application-default login
gcloud container clusters get-credentials $(terraform output -raw kubernetes_cluster_name) --region $(terraform output -raw region)
gcloud config get-value project

# terraform.tfvars
project_id = "REPLACE_ME"
region     = "us-central1"

terraform init
terraform apply

# configure kubectl
```

When creating credentials, make sure to add the following roles:

* Kubernetes Engine Admin
* Editor
* Service Account User
* Service Account Admin
