#!/bin/bash -e

source vars.sh

helm diff upgrade --install istio-routing-issue ./chart/istio-routing-issue \
    -f chart/istio-routing-issue/values.yaml \
    -n istio-routing-issue \
    --set "gateway.hosts={$ENDPOINT}" \
    --set serviceA.image=$DOCKERHUB_USERNAME/istio-routing-issue-service-a:$A_VERSION \
    --set serviceB.image=$DOCKERHUB_USERNAME/istio-routing-issue-service-b:$B_VERSION
