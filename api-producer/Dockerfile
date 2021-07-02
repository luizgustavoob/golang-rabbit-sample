FROM golang:alpine
WORKDIR /go/src/github.com/golang-rabbit-sample/api-producer
ENV GO111MODULE=off
RUN apk update && apk add bash git 
RUN go get github.com/gin-gonic/gin          
RUN go get github.com/streadway/amqp
COPY . .
EXPOSE 8889