FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD data.json /app/data.json
ADD grpc-testing-cli /app/grpc-testing-cli

EXPOSE 8080

CMD ["./grpc-testing-cli"]
