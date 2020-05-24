# psr-project

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

Create Terraform state bucket
```
export BACKEND_BUCKET_TF="your_bucket_name"
aws s3api create-bucket --bucket $BACKEND_BUCKET_TF --region us-east-1 --acl private --object-lock-enabled-for-bucket
```

# Terraform 

Create and init AWS bucket for terraform state
```
BACKEND_BUCKET_TF="your_bucket_name" \
envsubst < backend.tf.template > backend.tf
```

Create `secret.auto.tfvars` and set your GCP project id:
```
WEB_DOMAIN='your_domain' \
echo static_website = \"$WEB_DOMAIN\" >> secret.auto.tfvars
```

Init and apply terraform
```
terraform init
terraform apply
```

# Serverless framework

```
npm install -g serverless
. ./scripts/login-aws.sh
sls create --template aws-go-mod --path myService
```

Run unit test
```
go test ./uploadS3 -v -run TestS3Upload
```

# Frontend

```
npm run deploy
```