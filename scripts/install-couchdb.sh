#!/usr/bin/env sh

SCRIPTS_DIR=$(dirname "$0")
sh "$SCRIPTS_DIR"/create-k8s-ns.sh hyperledger

kubectl create secret generic couchdb0-couchdb \
  --from-literal=adminUsername=admin \
  --from-literal=adminPassword=adminpw \
  --from-literal=cookieAuthSecret=baz

helm install couchdb0 couchdb/couchdb -n hyperledger \
  --set couchdbConfig.couchdb.uuid="$(curl https://www.uuidgenerator.net/api/version4 2>/dev/null | tr -d -)" \
#  --set allowAdminParty=true

