cert-manager:
	kustomize build common/cert-manager/cert-manager/base | kubectl apply -f -

kubeflow-issuer:
	kustomize build common/cert-manager/kubeflow-issuer/base | kubectl apply -f -

istio-crd:
	kustomize build common/istio-1-9/istio-crds/base | kubectl apply -f -

istio-namespace:
	kustomize build common/istio-1-9/istio-namespace/base | kubectl apply -f -

istio-install:
	kustomize build common/istio-1-9/istio-install/base | kubectl apply -f -

dex:
	kustomize build common/dex/overlays/istio | kubectl apply -f -

oidc:
	kustomize build common/oidc-authservice/base | kubectl apply -f -

kubeflow-namespace:
	kustomize build common/kubeflow-namespace/base | kubectl apply -f -

kubeflow-roles:
	kustomize build common/kubeflow-roles/base | kubectl apply -f -

kubeflow-istio-resources:
	kustomize build common/istio-1-9/kubeflow-istio-resources/base | kubectl apply -f -

profiles-kfam:
	kustomize build apps/profiles/upstream/overlays/kubeflow | kubectl apply -f -

volumes:
	kustomize build apps/volumes-web-app/upstream/overlays/istio | kubectl apply -f -

notebook-controller:
	kustomize build apps/jupyter/notebook-controller/upstream/overlays/kubeflow | kubectl apply -f -

notebooks:
	kustomize build apps/jupyter/jupyter-web-app/upstream/overlays/istio | kubectl apply -f -

tensorboard:
	kustomize build apps/tensorboard/tensorboards-web-app/upstream/overlays/istio | kubectl apply -f -

tensorboard-controller:
	kustomize build apps/tensorboard/tensorboard-controller/upstream/overlays/kubeflow | kubectl apply -f -

training-operator:
	kustomize build apps/training-operator/upstream/overlays/kubeflow | kubectl apply -f -

mpi-operator:
	kustomize build apps/mpi-job/upstream/overlays/kubeflow | kubectl apply -f -

user-namespace:
	kustomize build common/user-namespace/base | kubectl apply -f -

central-dashboard:
	kustomize build apps/centraldashboard/upstream/overlays/istio | kubectl apply -f -

local:
	kubectl port-forward svc/istio-ingressgateway -n istio-system 8080:80

ingress-do-controller:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.1.0/deploy/static/provider/do/deploy.yaml