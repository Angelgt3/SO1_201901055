#Comandos
docker build -t backend-semi1 .
#Docker
docker build -t cliente_grcp .
docker tag cliente_grcp angegt3/cliente_grcp
docker push angegt3/cliente_grcp

docker build -t servidor_grcp .
docker tag servidor_grcp angegt3/servidor_grcp
docker push angegt3/servidor_grcp

docker build -t golang-consumer .
docker tag golang-consumer angegt3/golang-consumer
docker push angegt3/golang-consumer


#GRPC
protoc --go_out=. --go-grpc_out=. cliente.proto

#Locust
locust -f traffic.py

#Gcloud
gcloud container clusters get-credentials proyecto2 --location us-centra1l-c
kubectl create namespace so1

kubectl get pods -n so1
kubectl get deployments -n so1
kubectl get services -n so1

kubectl expose deployment grpc-deployment --type=LoadBalancer --port 3000 -n so1

#grafana
kubectl port-forward -n monitoring --address 0.0.0.0 svc/grafana 3000:3000

kubectl logs deploy/consumer-deployment -n so1 -f

#REDIS
kubectl get pods -n monitoring
kubectl exec -it redis-6fbbbc7b97-c5nc7 -n monitoring -- redis-cli -a YOUR_PASSWORD

#MONGO
kubectl get pods -n mongospace
kubectl exec -it <nombre-del-pod> -n mongospace -- /bin/bash
mongo -u admin -p password
