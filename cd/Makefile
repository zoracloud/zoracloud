install:
	kubectl apply -n argocd -f install.yaml

access:
	kubectl port-forward svc/argocd-server -n argocd 8080:443

pass:
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
