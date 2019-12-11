build:
	GOOS=linux GOARCH=amd64 go build -o rest-testing-cli
	docker build -t rest-testing-cli .
	docker tag rest-testing-cli:latest makubit/grpc-vs-rest-service_rest-testing-cli:latest
	docker push makubit/grpc-vs-rest-service_rest-testing-cli:latest