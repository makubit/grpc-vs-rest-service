FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/rest-sorting-service

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rest-sorting-service

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/rest-sorting-service .

CMD ["./rest-sorting-service"]