
NAME:=golang-docker
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build clean test build-container push-container test-container

build :
	export GO111MODULE=off
	export GOBIN=~/go/bin
	export GOPATH=~/go
	go build -o bin/server_main ./src

clean :
	@rm bin/*

test:
	go test -v ./...  

build-container :
	@docker build -t dsatya6/golang-docker:$(VERSION) .

push-container :
	@docker push dsatya6/golang-docker:$(VERSION)	

test-container :
	@docker rm -f golang-docker || true
	# go src is running a 8081, map to docker 8080
	@docker run -dp 8080:8081 --name=golang-docker $(DOCKER_IMAGE_NAME):$(VERSION)
	@docker ps
	@curl http://localhost:8080/test
	@echo ""
	@curl http://localhost:8080/hi