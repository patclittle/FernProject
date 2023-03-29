#!/bin/bash

imgName="fern-test-app"
imgTag="v4"
acrName="fernprojectacr"

echo "Building docker file ${imgName}:${imgTag} and pushing to ${acrName}"
docker build -t ${imgName}:${imgTag} -f ./Dockerfile ../src/
az acr login --name ${acrName}
docker tag ${imgName}:${imgTag} ${acrName}.azurecr.io/${imgName}:${imgTag}
docker push ${acrName}.azurecr.io/${imgName}:${imgTag}
