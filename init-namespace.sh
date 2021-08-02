#!/bin/bash

kubectl create namespace istio-routing-issue
kubectl label namespace istio-routing-issue istio.io/rev=1-6-14
