#!/bin/bash
export AWS_SHARED_CREDENTIALS_FILE=$PWD/terraform/secrets/aws.credentials
# alias sls='npx sls'
(cd uploadService && make deploy)