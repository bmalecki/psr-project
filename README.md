# psr-project

# Terraform 

terraform apply -var-file=secret.tfvars

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
export GOOGLE_CREDENTIALS="$HOME/service/google.json"

