FROM        golang:alpine AS base
WORKDIR     /go/src/github.com/golang-rabbit-sample/api-producer/

FROM        base AS dependencies
ENV         GO111MODULE=off
RUN         apk update && apk add bash git && rm -rf /var/cache/apk/*
RUN         go get github.com/stretchr/testify && \
            go get github.com/gin-gonic/gin && \
            go get github.com/streadway/amqp

FROM        dependencies AS build
COPY        . .
RUN         GOOS=linux GOARCH=amd64 go build -o bin ./cmd/api-producer

FROM        alpine:latest AS image
WORKDIR     /root/
RUN         apk update && apk add bash && rm -rf /var/cache/apk/*
COPY        --from=build /go/src/github.com/golang-rabbit-sample/api-producer/bin/api-producer .
COPY        wait-for-it.sh .
ENTRYPOINT  [ "./api-producer" ]
EXPOSE      8889