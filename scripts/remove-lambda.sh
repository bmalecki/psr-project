#!/bin/bash
export AWS_SHARED_CREDENTIALS_FILE=$PWD/terraform/secrets/aws.credentials

aws cloudformation describe-stack-resources --stack-name uploadservice-dev \
    --query 'StackResources[?ResourceType == `AWS::S3::Bucket`].[PhysicalResourceId]' --output text \
    | xargs -d ' ' \
    | xargs -I% aws s3 rm s3://%/ --recursive > /dev/null

(cd service && npx serverless remove)