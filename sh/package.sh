#!/bin/sh

DEPLOYMENT_ARTIFACT=deployment.zip

rm main ${DEPLOYMENT_ARTIFACT}
GOOS=linux go build -v -o main
zip -v -r ${DEPLOYMENT_ARTIFACT} main config.yaml
unzip -l ${DEPLOYMENT_ARTIFACT}

