#!/bin/sh

DEPLOYMENT_ARTIFACT=deployment.zip

rm main ${DEPLOYMENT_ARTIFACT}
GOOS=linux go build -o main
zip ${DEPLOYMENT_ARTIFACT} main config.yaml
