#!/bin/bash
export AWS_SHARED_CREDENTIALS_FILE=$PWD/terraform/secrets/aws.credentials
export AWS_DEFAULT_REGION=us-east-1
export API_GATEWAY_ID=$(aws apigateway get-rest-apis --output text --query 'items[?name == `dev-uploadservice`].id')
export REACT_APP_URL='https://'${API_GATEWAY_ID}'.execute-api.us-east-1.amazonaws.com/dev'
(cd web && npm install && npm run build)

STATIC_WEBSITE_ID=$(cd terraform && terraform output static_website_id)
aws s3 sync web/build s3://$STATIC_WEBSITE_ID
