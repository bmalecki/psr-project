#!/bin/bash
export AWS_SHARED_CREDENTIALS_FILE=$PWD/terraform/secrets/aws.credentials
export AWS_DEFAULT_REGION=us-east-1
export API_GATEWAY_ID=$(aws apigateway get-rest-apis --output text --query 'items[?name == `dev-uploadservice`].id')
export REACT_APP_URL='https://'${API_GATEWAY_ID}'.execute-api.us-east-1.amazonaws.com/dev'

IMAGES_BUCKET_NAME=$(aws cloudformation describe-stack-resources --stack-name uploadservice-dev \
    --query 'StackResources[?LogicalResourceId == `UploadImageStorage`].[PhysicalResourceId]' --output text \
    | xargs echo)

export REACT_APP_IMAGES_URL=http://$IMAGES_BUCKET_NAME.s3-website-us-east-1.amazonaws.com

(cd web && npm install && npm run build)

STATIC_WEBSITE_ID=$(cd terraform && terraform output static_website_id)
aws s3 sync web/build s3://$STATIC_WEBSITE_ID
