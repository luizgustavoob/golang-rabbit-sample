API_IMAGE = luizgustavoob/golang-rabbit-api:latest
DBCONSUMER_IMAGE = luizgustavoob/golang-rabbit-dbconsumer:latest

.PHONY: image
image:
	docker image build \
		-t $(API_IMAGE) \
		--target=final \
		-f api-producer/Dockerfile \
		api-producer/

	docker image build \
		-t $(DBCONSUMER_IMAGE) \
		--target=final \
		-f database-service-consumer/Dockerfile \
		database-service-consumer/

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down