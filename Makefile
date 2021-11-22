install-platform:
	while ! kustomize build manifests/platform | kubectl apply -f -; do echo "Retrying to apply resources"; sleep 10; done

dev:
	kubectl port-forward svc/istio-ingressgateway -n istio-system 8080:80