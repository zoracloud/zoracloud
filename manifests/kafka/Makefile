install:
	kubectl create -f install.yaml -n kafka

cluster:
	kubectl apply -f cluster.yaml -n kafka

status:
	kubectl wait kafka/zora --for=condition=Ready --timeout=300s -n kafka