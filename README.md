# psr-project

# Terraform 

Create `secret.auto.tfvars` and set your GCP project id:
`gcp_project="your_project_id"`

terraform apply

# GCP 

## Add iam

```
export GCP_USER=''
export GCP_PROJECT=''
gcloud projects add-iam-policy-binding $GCP_PROJECT \
  --member serviceAccount:$GCP_USER \
  --role roles/editor
```

## Terraform prerequisites 

### Service User Credentials
export GOOGLE_CREDENTIALS="$HOME/service/google.json"

