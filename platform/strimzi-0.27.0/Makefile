kafka-cluster-operator:
	kustomize build | kubectl apply -f -

sed: 
	sed -i 's/namespace: .*/namespace: kafka/' install/cluster-operator/*RoleBinding*.yaml
	sed -i '' 's/namespace: .*/namespace: kafka/' install/cluster-operator/*RoleBinding*.yaml

install:
	kubectl create -f install/cluster-operator/ -n kafka

permission:
	kubectl create -f install/cluster-operator/020-RoleBinding-strimzi-cluster-operator.yaml -n zorakafka
	kubectl create -f install/cluster-operator/031-RoleBinding-strimzi-cluster-operator-entity-operator-delegation.yaml -n zorakafka


