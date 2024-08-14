# README

```sh
gcloud config get-value project

# terraform.tfvars
project_id = "REPLACE_ME"
region     = "us-central1"

terraform init
terraform apply

# configure kubectl
gcloud container clusters get-credentials $(terraform output -raw kubernetes_cluster_name) --region $(terraform output -raw region)
```
