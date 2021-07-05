API_IMAGE 		 = luizgustavoob/golang-rabbit-api:latest
DBCONSUMER_IMAGE = luizgustavoob/golang-rabbit-dbconsumer:latest

.PHONY: install
install:
	GO111MODULE=off go get github.com/stretchr/testify
	GO111MODULE=off go get github.com/gin-gonic/gin
	GO111MODULE=off go get github.com/streadway/amqp
	GO111MODULE=off go get github.com/lib/pq

.PHONY: build
build:
	docker image build \
		--tag $(API_IMAGE) \
		--target=build \
		--file ./api-producer/Dockerfile \
		./api-producer

	docker image build \
		--tag $(DBCONSUMER_IMAGE) \
		--target=build \
		--file ./database-service-consumer/Dockerfile \
		./database-service-consumer

.PHONY: image
image:
	docker image build \
		--tag $(API_IMAGE) \
		--target=image \
		--file ./api-producer/Dockerfile \
		./api-producer

	docker image build \
		--tag $(DBCONSUMER_IMAGE) \
		--target=image \
		--file ./database-service-consumer/Dockerfile \
		./database-service-consumer

.PHONY: up
up: image
	docker-compose up -d

.PHONY: down
down:
	docker-compose down