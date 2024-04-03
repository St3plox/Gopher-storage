GOLANG          	:= golang:1.22
ALPINE 				:= alpine:3.19
BASE_IMAGE_NAME 	:= localhost/tveu/storage
SERVICE_NAME    	:= storage-api
CONTAINER_NAME  	:= storage-container
EXPOSE_PORT			:= 3000
INTERNAL_PORT		:= 3000
VERSION				:= 0.0.1
SERVICE_IMAGE		:= $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
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
	sudo go run app/services/main-storage/main.go

dev-test-curl:
	curl localhost:$(EXPOSE_PORT)/storage/1

build-service-image:
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
	docker stop $(NAME)
	docker rm $(NAME)