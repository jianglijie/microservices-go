#!/bin/bash

mode=prod

# docker
make -f services/Makefile

# k8s
kubectl apply -f build/mysql.yaml
kubectl apply -f build/volume.yaml
kubectl apply -f services/demo-services-cm-${mode}.yaml
kubectl apply -f services/k8s-services.yaml
kubectl apply -f services/demo-services-hpa.yaml
