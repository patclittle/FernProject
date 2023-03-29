helm uninstall main
helm install main .

# sleep 1s to let the pod startup
usleep 1000000

echo "Go to http://localhost/ideas to try it out!"
kubectl port-forward fern-project-main-574ccc7bdc-zh48l-ltmng 8080:8080