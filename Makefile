#Makefile

include ./include.mk

help:
	@echo "$(TAG) will build for $(CI_PROJECT_NAME)"
	@echo "make build    = build docker images"
	@echo "make push     = push image to repository"

build: build-testing
	@echo "done building images for $(CI_PROJECT_NAME)"

build-testing:
	@echo "build image for $(IMG)/build"
	docker build -t "$(IMG)/build:$(TAG)" \
		-f build/testing/Dockerfile .

push:
	docker push $(IMG)/build:$(TAG)

images: build

deploy:
	$(MAKE) -C docker deploy