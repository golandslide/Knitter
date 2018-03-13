#!/bin/bash
export PATH="$HOME/.kubeadm-dind-cluster:$PATH"
kubectl get nodes
ifconfig
pwd
kubectl create -f ./hack/create-testpod.yaml
kubectl get pod
kubectl get pod
kubectl get pod

kubectl get  pod/nginx -o json
date