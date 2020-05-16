# psr-project

# GCP 

## Add iam

Create service account
```
gcloud iam service-accounts create terraform-service-user-psr
```

Add role to service your service user

```
GCP_USER='terraform-service-user-psr' \
GCP_PROJECT='your_project' \
gcloud projects add-iam-policy-binding $GCP_PROJECT \
  --member serviceAccount:${GCP_USER}@${GCP_PROJECT}.iam.gserviceaccount.com \
  --role roles/editor
```

Generate and save key
```
gcloud iam service-accounts keys create $HOME/service/google.json \
  --iam-account ${GCP_USER}@${GCP_PROJECT}.iam.gserviceaccount.com
```


Create bucket for Terraform state

```
GCP_PROJECT='your_project' \
GCP_BACKEND_BUCKET_TF="your_bucket_name" \
gsutil mb -p ${GCP_PROJECT} -c STANDARD -l US-EAST1 -b on gs://${GCP_BACKEND_BUCKET_TF}/
```

# AWS
Generate service user and set up.

```
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_SESSION_TOKEN=""
```

# Terraform 

Export GCP Service User Credentials
```
export GOOGLE_CREDENTIALS="$HOME/service/google.json"
```

Create and init GCP bucket for terraform state
```
GCP_BACKEND_BUCKET_TF="your_bucket_name" envsubst < backend.tf.template > backend.tf
```

Create `secret.auto.tfvars` and set your GCP project id:
```
echo gcp_project="your_project_id" > secret.auto.tfvars
```

Init and apply terraform
```
terraform init
terraform apply
```
