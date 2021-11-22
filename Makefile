install-platform:
	while ! kustomize build manifests/platform | kubectl apply -f -; do echo "Retrying to apply resources"; sleep 10; done