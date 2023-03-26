helm uninstall main
helm install main .

# sleep 1s to let the pod startup
usleep 1000000

echo "Go to http://localhost/ideas to try it out!"
kubectl port-forward fern-project-main-68df48ccb4-ltmng 8080:8080