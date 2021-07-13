FROM        golang:alpine AS base
WORKDIR     $GOPATH/src/github.com/golang-rabbit-sample/database-service-consumer

FROM        base AS dependencies
ENV         GO111MODULE=on
COPY        go.mod .
COPY        go.sum .
RUN         go mod download

FROM        dependencies AS build
COPY        . .
RUN         GOOS=linux GOARCH=amd64 go build -o /go/bin/database-service-consumer ./cmd/database-service-consumer

FROM        alpine:latest AS image
WORKDIR     /root/
RUN         apk update && apk add bash && rm -rf /var/cache/apk/*
COPY        --from=build /go/bin/database-service-consumer .
ENTRYPOINT  [ "./database-service-consumer" ]