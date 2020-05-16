# psr-project

# Terraform 


Create and init GCP bucket for terraform state
```
GCP_BACKEND_BUCKET_TF="your_bucket_name" envsubst < backend.tf.template > backend.tf
```

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

