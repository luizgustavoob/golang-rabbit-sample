FROM golang:alpine
WORKDIR /go/src/github.com/golang-rabbit-sample/database-service-consumer
COPY . .
RUN chmod +x wait-for-it.sh
RUN apk update 
RUN apk add bash 
RUN apk add --no-cache git 
RUN go get -u github.com/lib/pq 
RUN go get github.com/streadway/amqp