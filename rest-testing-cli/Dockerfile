FROM debian:latest

RUN mkdir -p /app
WORKDIR /app

ADD rest-testing-cli /app/rest-testing-cli
ADD data.json /app/data.json

EXPOSE 8080

CMD ["./rest-testing-cli"]