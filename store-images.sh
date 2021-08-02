#!/bin/bash -e

source vars.sh
cd service_a && make && cd -
cd service_b && make && cd -
docker image build service_a -t istio-routing-issue-service-a
docker image build service_b -t istio-routing-issue-service-b
docker image tag istio-routing-issue-service-a:latest $DOCKERHUB_USERNAME/istio-routing-issue-service-a:$A_VERSION
docker image tag istio-routing-issue-service-b:latest $DOCKERHUB_USERNAME/istio-routing-issue-service-b:$B_VERSION
docker push $DOCKERHUB_USERNAME/istio-routing-issue-service-a:$A_VERSION
docker push $DOCKERHUB_USERNAME/istio-routing-issue-service-b:$B_VERSION
