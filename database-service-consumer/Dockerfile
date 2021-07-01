FROM golang:alpine
WORKDIR /go/src/github.com/golang-rabbit-sample/database-service-consumer
ENV GO111MODULE=off
RUN apk update && apk add bash && apk add git 
RUN go get -u github.com/lib/pq 
RUN go get github.com/streadway/amqp
COPY . .