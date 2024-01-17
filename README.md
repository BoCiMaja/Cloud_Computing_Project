# Cloud_Computing_Project
## Prerequisites 
Install `flagger` using `helm`.
```
helm repo add flagger https://flagger.app
```
Install flagger's canary custom resource definitions.
```
kubectl apply -f https://raw.githubusercontent.com/fluxcd/flagger/main/artifacts/flagger/crd.yaml
```
Install `istio`
```
istioctl manifest install --set profile=default
```
## Deployment
Deploy an `ingress gateway`
```
kubectl apply -f ingress-gateway.yaml
```
Create a `namespace`
```
kubectl create namespace <namespace-name>
```
Enable istio injection
```
kubectl label namespace <namespace-name> istio-injection=enabled
```
Deploy a `horizontal pod autoscaler`
```
kubectl apply -k https://github.com/fluxcd/flagger//kustomize/podinfo?ref=main
```
Deploy flagger's `loadtesting` service
```
kubectl apply -k https://github.com/fluxcd/flagger//kustomize/tester?ref=main
```
Create and upload different versions of a service to dockerhub.
`<dockerhub-username>/<service>:0.0.1`
`<dockerhub-username>/<service>:0.0.2`
`<dockerhub-username>/<service>:0.0.3`

Create and deploy `deployment.yaml` that uses `<dockerhub-username>/<service>:0.0.1`.
```
kubectl apply -f deployment.yaml
```
Create and deploy a `canary.yaml` that is set to manage our previous deployment.
```
kubectl apply -f canary.yaml
```
Canary will create the resources that it needs.
Can check the status of our canary deployment
```
kubectl describe canary <canary-name>
```
The deployment is working but it is not doing anything.
Set the `image` of the canary deployment to 
`<dockerhub-username>/<service>:0.0.2`
```
kubectl -n <namespace-name> set image deployment/<canary-deployment-name> <container-name>:<dockerhub-username>/<service>:0.0.2
```
Check canary deployment again.
```
kubectl describe canary <canary-name>
```
Canary found a new revision, it will start the scaling up and deployment process which can be monitored here.
The strategy for `deployment` and `rollback` and their `thresholds` are contained in the `canary.yaml` file.


