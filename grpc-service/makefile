build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/grpcService/grpcService.proto
	#GOOS=linux go build
	docker build -t grpc-service .
	docker tag grpc-service:latest makubit/grpc-vs-rest-service_grpc-service:latest
	docker push makubit/grpc-vs-rest-service_grpc-service:latest

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 grpc-service