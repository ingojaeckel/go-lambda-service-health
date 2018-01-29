#!/bin/sh

FN_NAME=$1
REGION=$2
AWS_ACCOUNT_ID=$3
ROLE=$4
DEPLOYMENT_ARTIFACT=$5

echo "Checking if function exists..."
aws lambda get-function --region ${REGION} --function-name $FN_NAME

if [[ $? -eq "0" ]]; then
    echo "Updating function ${FN_NAME} .."

    aws lambda update-function-code					\
        --region ${REGION}					\
        --function-name ${FN_NAME}				\
        --zip-file fileb://./${DEPLOYMENT_ARTIFACT}		\
        --runtime go1.x						\
        --role arn:aws:iam::${AWS_ACCOUNT_ID}:role/${ROLE} 	\
        --handler main
else
    echo "Creating function ${FN_NAME}.."

    aws lambda create-function				                \
    	--region ${REGION}					                \
    	--function-name ${FN_NAME}				            \
    	--zip-file fileb://./${DEPLOYMENT_ARTIFACT}		    \
    	--runtime go1.x						                \
    	--role arn:aws:iam::${AWS_ACCOUNT_ID}:role/${ROLE} 	\
    	--handler main
fi