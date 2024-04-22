#Comandos
#Docker
docker build -t cliente_grcp .
docker run cliente_grcp

docker build -t servidor_grcp .
docker run servidor_grcp

#GRPC
protoc --go_out=. --go-grpc_out=. cliente.proto

#Locust
locust -f traffic.py