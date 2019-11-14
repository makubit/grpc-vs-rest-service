build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:. \
	  proto/grpcService/grpcService.proto
	#GOOS=linux go build
	docker build -t grpc-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 grpc-service