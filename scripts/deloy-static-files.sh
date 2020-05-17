#!/bin/bash
export AWS_CONFIG_FILE=$PWD/terraform/secrets/aws.credentials
STATIC_WEBSITE_ID=$(cd terraform && terraform output static_website_id)
aws s3 sync web/build s3://$STATIC_WEBSITE_ID