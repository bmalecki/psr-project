#!/bin/bash
export AWS_SHARED_CREDENTIALS_FILE=$PWD/terraform/secrets/aws.credentials
(cd service && make deploy)