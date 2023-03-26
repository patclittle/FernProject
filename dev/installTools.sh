#!/bin/bash
# Install Kubectl
echo "*********************************************************************"
echo "************************ Installing Kubectl *************************"
echo "** (https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/) **"
echo "*********************************************************************"
echo ""
echo ""
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl"

# install K9s 
echo "*********************************************************************"
echo "************************** Installing k9s ***************************"
echo "**************** (https://k9scli.io/topics/install/) ****************"
echo "*********************************************************************"
echo ""
echo ""
brew install derailed/k9s/k9s

# install kubelogin
echo "*********************************************************************"
echo "*********************** Installing kubelogin ************************"
echo "*************** (https://github.com/int128/kubelogin) ***************"
echo "*********************************************************************"
echo ""
echo ""
brew install int128/kubelogin/kubelogin

# install az cli
echo "***********************************************************************"
echo "************************* Installing az cli **************************"
echo "*(https://learn.microsoft.com/en-us/cli/azure/install-azure-cli-macos)*"
echo "***********************************************************************"
echo ""
echo ""
brew install azure-cli

# install AKS tools
echo "*********************************************************************************"
echo "**************************** Installing AKS tools *******************************"
echo "*(https://learn.microsoft.com/en-us/azure/aks/learn/quick-kubernetes-deploy-cli)*"
echo "*********************************************************************************"
echo ""
echo ""
az aks install-cli

# install helm
echo "*****************************************************************"
echo "*********************** Installing helm *************************"
echo "****(https://helm.sh/docs/intro/install/#from-homebrew-macos)****"
echo "*****************************************************************"
echo ""
echo ""
brew install helm

# install Docker
echo "******************************************************************"
echo "*********************** Installing Docker ************************"
echo "******(https://docs.docker.com/desktop/install/mac-install/)******"
echo "******************************************************************"
echo "******************************************************************"
echo "***************** BE READY TO ENTER PASSWORD!!!! *****************"
echo "******************************************************************"
echo ""
echo ""
curl --output Docker.dmg https://desktop.docker.com/mac/main/amd64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-amd64
chmod +x Docker.dmg
hdiutil attach Docker.dmg
sudo /Volumes/Docker/Docker.app/Contents/MacOS/install --accept-license
sudo hdiutil detach /Volumes/Docker