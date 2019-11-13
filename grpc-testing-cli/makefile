build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t grpc-testing-cli .

run:
	docker run grpc-testing-cli