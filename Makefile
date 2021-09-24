centraldashboard:
	kustomize build components/dashboard/manifests/overlays/istio | kubectl apply -f -

user:
	kustomize build user | kubectl apply -f -
