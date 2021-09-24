# run:
# 	docker run -d -p 19091:19091 \
# 		-p 19092:19092 \
# 		mailgun/kafka-pixy --config /etc/kafka-pixy.yaml

pixy:
	kubectl port-forward svc/pixy -n kafka 9000:80

create-secret:
	kubectl create -n default secret tls ingressgateway-certs --key=privkey.pem --cert=cert.pem
