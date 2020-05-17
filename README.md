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
GCP_USER='terraform-service-user-psr' \
GCP_PROJECT='your_project' \
gcloud iam service-accounts keys create terraform/secrets/gcp.json \
  --iam-account ${GCP_USER}@${GCP_PROJECT}.iam.gserviceaccount.com
```

Create bucket for Terraform state

```
GCP_PROJECT='your_project' \
GCP_BACKEND_BUCKET_TF="your_bucket_name" \
gsutil mb -p ${GCP_PROJECT} -c STANDARD -l US-EAST1 -b on gs://${GCP_BACKEND_BUCKET_TF}/
```

# AWS
Generate service user set up terraform/secrets/aws.credentials

```
[default]
aws_access_key_id=
aws_secret_access_key=
aws_session_token=
```

```
export AWS_CONFIG_FILE=$PWD/terraform/secrets/aws.credentials
```

# Terraform 

Create and init GCP bucket for terraform state
```
GCP_BACKEND_BUCKET_TF="your_bucket_name" \
envsubst < backend.tf.template > backend.tf
```

Create `secret.auto.tfvars` and set your GCP project id:
```
GCP_PROJECT='your_project' \
echo gcp_project=\"$GCP_PROJECT\"> secret.auto.tfvars
```

Init and apply terraform
```
terraform init
terraform apply
```


# Serverless framework

```
npm install -g serverless
```