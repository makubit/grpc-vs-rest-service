build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t grpc-testing-cli .
	docker tag grpc-testing-cli:latest makubit/grpc-vs-rest-service_grpc-testing-cli:latest
	docker push makubit/grpc-vs-rest-service_grpc-testing-cli:latest

run:
	docker run grpc-testing-cli