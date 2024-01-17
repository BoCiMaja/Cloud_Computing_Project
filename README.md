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
kubectl label namespace test istio-injection=enabled
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

Create a deployment.yaml 
