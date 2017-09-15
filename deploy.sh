#!/usr/bin/env bash
az group deployment create --name bketelsen-aci-c9 --template-file azuredeploy.json --parameters @azuredeploy.parameters.json --resource-group homeStoreRG

