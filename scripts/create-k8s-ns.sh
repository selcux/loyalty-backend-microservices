#!/usr/bin/env sh

NS=$(kubectl get ns | grep "$1")

if [ -z "$NS" ]; then
  kubectl create ns "$1"
else
  echo "Namespace '$1' already exists, skipping..."
fi