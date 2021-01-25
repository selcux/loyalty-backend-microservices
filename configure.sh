#!/usr/bin/env bash

echo "Configuring Minikube..."

echo "Enabling ingress..."
minikube addons enable ingress

echo "...DONE!..."