FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
ADD grpc-testing-cli /app/grpc-testing-cli

CMD ["./grpc-testing-cli"]