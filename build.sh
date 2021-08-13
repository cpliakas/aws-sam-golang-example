#!/bin/bash
# builds the project and deploys it

echo "compile"
GOARCH=amd64 GOOS=linux go build -o api ./service/api

echo "package"
sam package --template-file template.yaml --s3-bucket $S3_BUCKET --output-template-file packaged.yaml

echo "deploy"
sam deploy --stack-name $STACK_NAME --template-file packaged.yaml --capabilities CAPABILITY_IAM

echo "done"