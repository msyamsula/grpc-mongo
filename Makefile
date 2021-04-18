up:
	go run server/server.go
startc:
	go run client/client.go
servers:
	docker-compose up -d --scale grpc-server=3
down:
	docker-compose down
build:
	docker build -t grpc-mongo .
	docker build -f Dockerfile.nginx -t grpc-mongo-nginx .