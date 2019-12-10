build:
	#GOOS=linux GOARCH=amd64 go build -o rest-service
	docker build -t rest-service .
	docker tag rest-service:latest makubit/grpc-vs-rest-service_rest-service:latest
	docker push makubit/grpc-vs-rest-service_rest-service:latest