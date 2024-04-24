#Comandos
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

docker build -t cliente-grafana .
docker tag cliente-grafana angegt3/cliente-grafana
docker push angegt3/cliente-grafana
angegt3/cliente-grafana


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