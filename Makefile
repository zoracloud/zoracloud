centraldashboard:
	kustomize build components/centraldashboard/manifests/overlays/istio | kubectl apply -f -

user:
	kustomize build user | kubectl apply -f -

