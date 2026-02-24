# Taming the Multi-Cluster Sprawl: Add-on Automation with Project Sveltos

[Presentation slides](https://docs.google.com/presentation/d/1kbYUEIAVpcOxA5c2zZFBe6pqxSrXOWVK/edit?usp=drive_link&ouid=113873886986304892682&rtpof=true&sd=true)
## Setup sveltos & sveltos dashboard
Install the helm charts
```
helm install projectsveltos projectsveltos/projectsveltos -n projectsveltos --create-namespace --debug
```
Install dashboard
```
helm install projectsveltos-dashboard projectsveltos/sveltos-dashboard -n projectsveltos --debug
```
Create required RBACs
```
kubectl apply -f ./sveltos-dashboard-rbac.yaml
```
Get the token
```
kubectl -n projectsveltos create token sveltos-dashboard --duration=24h
```
Port-forward the dashboard svc & visit https://localhost:8000 on your browser
```
k port-forward svc/dashboard -n projectsveltos 8000:80
```
## clusterprofile creation, adding a new cluster to the fleet
Apply clusterprofiles
```
kubectl apply -f ./clusterprofiles/
```
Register a new cluster with env=uat
```
sveltosctl register cluster --namespace=kcd \
  --cluster=pf9-uat \
  --kubeconfig=<path/to/kubeconfigs/sveltos-uat-kubeconfig.yaml> \
  --labels \
  env=uat
```
Check cluster profile status
```
kubectl get clusterprofiles demo-httpbin-app -o yaml | yq .status
```
## classifier example: automate label addition based on k8s version
Create classifiers
```
kubectl apply -f ./classifiers/gittea-classifiers.yaml
```
Check labels on clusters
```
kubectl get sveltoscluster -n kcd --show-labels
```
## event source & event trigger example
Change context to the managed cluster that's labeled with env=dev
```
export KUBECONFIG=<path/to/kubeconfigs/sveltos-dev-kubeconfig.yaml>
```
Checkout event trigger definition
```
kubectl get eventtrigger service-network-policy -o yaml | yq
```
Check the event source
```
kubectl get eventsource sveltos-service -o yaml | yq
```
Apply the demo svc
```
kubectl apply -f ./event-driven-deployment/demo-svc.yaml
```
Wait for the networkpolicy get created by sveltos
```
kubectl get networkpolicy -w
```

