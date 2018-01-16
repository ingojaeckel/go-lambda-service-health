#!/bin/sh

FN_NAME=ServiceHealth
REGION=us-east-1
AWS_ACCOUNT_ID=123
ROLE=lambda-role
DEPLOYMENT_ARTIFACT=deployment.zip

rm main ${DEPLOYMENT_ARTIFACT}
GOOS=linux go build -o main
zip ${DEPLOYMENT_ARTIFACT} main
aws lambda create-function \
	--region ${REGION} \
	--function-name ${FN_NAME} \
	--zip-file fileb://./${DEPLOYMENT_ARTIFACT} \
	--runtime go1.x \
	--tracing-config Mode=Active
	--role arn:aws:iam::${AWS_ACCOUNT_ID}:role/${ROLE} \
	--handler main
