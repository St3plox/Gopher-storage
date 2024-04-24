GOLANG          	:= golang:1.22
ALPINE 				:= alpine:3.19
BASE_IMAGE_NAME 	:= localhost/tveu/storage
SERVICE_NAME    	:= node-api
GATEWAY_NAME    	:= gateway-api
CONTAINER_NAME  	:= node
EXPOSE_PORT			:= 3000
INTERNAL_PORT		:= 3000
VERSION				:= 0.0.1
SERVICE_IMAGE		:= $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
GATEWAY_IMAGE		:= $(BASE_IMAGE_NAME)/$(GATEWAY_NAME):$(VERSION)
DOCKER_COMPOSE_FILE := zarf/docker/service/docker-compose.yml
NAME 		:= service_storage-api_1





tools:
	go install github.com/divan/expvarmon@latest

dev-docker:
	docker pull $(ALPINE)
	docker pull $(GOLANG)

tidy:
	go mod tidy
	go mod vendor

service-run-local:
	sudo go run app/services/node-api/main.go

run-foundation-tests:
	sudo go test ./foundation/storage -v

run-tests: run-foundation-tests

dev-test-curl:
	curl localhost:$(EXPOSE_PORT)/storage/1

service-build-image:
	docker build -t $(SERVICE_IMAGE) -f zarf/docker/service/Dockerfile .

docker-compose-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-compose-down:
	docker-compose f $(DOCKER_COMPOSE_FILE) down

service-logs:
	docker logs service_storage-api_1 -f

docker-compose-logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs

metrics:
	expvarmon -ports="localhost:4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

service-run:
	docker run -d -p $(EXPOSE_PORT):$(INTERNAL_PORT) --name $(CONTAINER_NAME) $(SERVICE_IMAGE)

service-stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

gateway-build-image:
	docker build -t $(GATEWAY_IMAGE) -f zarf/docker/gateway/Dockerfile .

gateway-run:
	docker run -d -p 8081:8081 --name $(GATEWAY_NAME) $(GATEWAY_IMAGE)

gateway-run-local:
	 go run app/services/gateway-api/main.go

gateway-stop:
	docker stop $(GATEWAY_NAME)
	docker rm $(GATEWAY_NAME)

