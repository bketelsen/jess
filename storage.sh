#!/usr/bin/env bash
# You MUST create a resource group first, or choose an existing one.

# Change these four parameters
ACI_PERS_STORAGE_ACCOUNT_NAME=homestorebketelsen
ACI_PERS_RESOURCE_GROUP=homeStoreRG
ACI_PERS_LOCATION=eastus
ACI_PERS_SHARE_NAME=acishare
ACI_PERS_STORAGE_KEY=azurefilestoragekey
ACI_PERS_VAULT_NAME=home-keyvault

# create the resource group
az group create --location $ACI_PERS_LOCATION --name $ACI_PERS_RESOURCE_GROUP

# Create the storage account with the parameters
az storage account create -n $ACI_PERS_STORAGE_ACCOUNT_NAME -g $ACI_PERS_RESOURCE_GROUP -l $ACI_PERS_LOCATION --sku Standard_LRS

# Export the connection string as an environment variable, this is used when creating the Azure file share
export AZURE_STORAGE_CONNECTION_STRING=`az storage account show-connection-string -n $ACI_PERS_STORAGE_ACCOUNT_NAME -g $ACI_PERS_RESOURCE_GROUP -o tsv`

# Create the share
az storage share create -n $ACI_PERS_SHARE_NAME


STORAGE_ACCOUNT=$(az storage account list --resource-group $ACI_PERS_RESOURCE_GROUP --query "[?contains(name,'homestore')].[name]" -o tsv)
echo "Storage Account: $STORAGE_ACCOUNT"

STORAGE_KEY=$(az storage account keys list --resource-group $ACI_PERS_RESOURCE_GROUP --account-name $STORAGE_ACCOUNT --query "[0].value" -o tsv)
echo "Storage Key: $STORAGE_KEY"

# create the keyvault
az keyvault create -n $ACI_PERS_VAULT_NAME --enabled-for-template-deployment -g $ACI_PERS_RESOURCE_GROUP

# set the secret
az keyvault secret set --vault-name $ACI_PERS_VAULT_NAME --name $ACI_PERS_STORAGE_KEY --value $STORAGE_KEY

# print the resource ID
az keyvault show --name $ACI_PERS_VAULT_NAME --query [id] -o tsv
