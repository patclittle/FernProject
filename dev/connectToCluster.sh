#!/bin/bash

echo "About to launch a browser window to login. Please follow the prompt and return to this window."
az login

echo "Setting subscription to 9ed9130b-e74a-4aa3-8d56-2caf8a06b9d7"
az account set --subscription 9ed9130b-e74a-4aa3-8d56-2caf8a06b9d7

echo "Fetching creds for 'ferntestaks' cluster in 'FernProject' resource group"
az aks get-credentials --resource-group FernProject --name ferntestaks

echo "Using kubecontxt ferntestaks"
kubectl config use-context ferntestaks