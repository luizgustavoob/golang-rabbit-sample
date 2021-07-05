FROM        golang:alpine AS base
WORKDIR     /go/src/github.com/golang-rabbit-sample/database-service-consumer/

FROM        base AS dependencies
ENV         GO111MODULE=off
RUN         apk update && apk add bash git && rm -rf /var/cache/apk/*
RUN         go get github.com/stretchr/testify && \
            go get github.com/lib/pq && \
            go get github.com/streadway/amqp

FROM        dependencies AS build
COPY        . .
RUN         GOOS=linux GOARCH=amd64 go build -o bin ./cmd/database-service-consumer

FROM        alpine:latest AS image
WORKDIR     /root/
RUN         apk update && apk add bash && rm -rf /var/cache/apk/*
COPY        --from=build /go/src/github.com/golang-rabbit-sample/database-service-consumer/bin/database-service-consumer .
COPY        wait-for-it.sh .
ENTRYPOINT  [ "./database-service-consumer" ]