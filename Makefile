GOLANG          := golang:1.22
EXPOSE_PORT 	:= 8080
INTERNAL_PORT	:= 8080
ALPINE          := alpine:3.19
BASE_IMAGE_NAME := localhost/tveu/storage
SERVICE_NAME    := storage-api
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)

dev-docker:
	docker pull $(GOLANG)
	docker pull $(ALPINE)
	
tidy:
	go mod tidy
	go mod vendor
	
service-run-local:
	go run app/services/main-storage/main.go

dev-test-curl:
	curl localhost:$(EXPOSE_PORT)/storage/1