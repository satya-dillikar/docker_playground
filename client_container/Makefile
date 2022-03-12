
NAME:=client-container
DOCKER_REPOSITORY:=dsatya6
DOCKER_IMAGE_NAME:=$(DOCKER_REPOSITORY)/$(NAME)
VERSION:=0.1.0

.PHONY: build-container push-container test-container

build-container :
	@docker build -t dsatya6/$(NAME):$(VERSION) .

push-container :
	@docker push dsatya6/$(NAME):$(VERSION)

test-container :
	@docker rm -f $(NAME) || true
	@docker run -it --name=$(NAME) $(DOCKER_IMAGE_NAME):$(VERSION) /bin/bash
