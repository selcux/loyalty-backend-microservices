#!/usr/bin/env sh

DOMAIN="loyalty"
DNS_LABEL="$DOMAIN.northeurope.cloudapp.azure.com"
#STATIC_IP=$(az network public-ip create --resource-group loyalty-rg --name loyalty-ip --sku Standard --allocation-method static --dns-name "$DOMAIN" --query publicIp.ipAddress -o tsv)
STATIC_IP="52.169.43.16"

#helm install ingress-nginx ingress-nginx/ingress-nginx \
#  --set controller.service.loadBalancerIP="$STATIC_IP" \
#  --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-dns-label-name"="$DOMAIN.northeurope.cloudapp.azure.com"

SCRIPTS_DIR=$(dirname "$0")
sh "$SCRIPTS_DIR"/create-k8s-ns.sh ingress-basic

helm install ingress-nginx ingress-nginx/ingress-nginx \
    --namespace ingress-basic \
    --set controller.replicaCount=1 \
    --set controller.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set defaultBackend.nodeSelector."beta\.kubernetes\.io/os"=linux \
    --set controller.admissionWebhooks.patch.nodeSelector."beta\.kubernetes\.io/os"=linux \
#    --set controller.service.loadBalancerIP="$STATIC_IP" \
#    --set controller.service.annotations."service\.beta\.kubernetes\.io/azure-dns-label-name"="$DNS_LABEL"