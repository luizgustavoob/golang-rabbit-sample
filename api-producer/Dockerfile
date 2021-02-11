FROM golang:alpine
WORKDIR /go/src/github.com/golang-rabbit-sample/api-producer
COPY . .
RUN apk update 
RUN apk add bash 
RUN apk add --no-cache git 
RUN go get -u github.com/gin-gonic/gin          
RUN go get github.com/streadway/amqp
EXPOSE 8889