API_IMAGE = luizgustavoob/golang-rabbit-api:latest
DBCONSUMER_IMAGE = luizgustavoob/golang-rabbit-dbconsumer:latest

.PHONY: install
install:
	export GO111MODULE=off
	go get github.com/stretchr/testify
	go get github.com/gin-gonic/gin
	go get github.com/streadway/amqp
	go get github.com/lib/pq

.PHONY: build
build:
	docker image build \
		-t $(API_IMAGE) \
		--target=build \
		-f api-producer/Dockerfile \
		api-producer/

	docker image build \
		-t $(DBCONSUMER_IMAGE) \
		--target=build \
		-f database-service-consumer/Dockerfile \
		database-service-consumer/

.PHONY: image
image:
	docker image build \
		-t $(API_IMAGE) \
		--target=image \
		-f api-producer/Dockerfile \
		api-producer/

	docker image build \
		-t $(DBCONSUMER_IMAGE) \
		--target=image \
		-f database-service-consumer/Dockerfile \
		database-service-consumer/

.PHONY: up
up: image
	docker-compose up -d

.PHONY: down
down:
	docker-compose down