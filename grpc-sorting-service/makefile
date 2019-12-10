build:
	protoc -I. --go_out=plugins=micro:. \
      proto/sortingService/sortingService.proto
	docker build -t grpc-sorting-service .

run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 grpc-sorting-service