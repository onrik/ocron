PROJECT_NAME = ocron
VERSION = 1.0
BUILD = 0
BUILD_VERSION = $(VERSION).$(BUILD)

build:
	docker build . -t $(PROJECT_NAME):$(BUILD_VERSION)

push:
	docker tag $(PROJECT_NAME):$(BUILD_VERSION) onrik/$(PROJECT_NAME):$(BUILD_VERSION)
	docker tag $(PROJECT_NAME):$(BUILD_VERSION) onrik/$(PROJECT_NAME):$(VERSION)
	docker tag $(PROJECT_NAME):$(BUILD_VERSION) onrik/$(PROJECT_NAME):latest
	docker push onrik/$(PROJECT_NAME):$(BUILD_VERSION)
	docker push onrik/$(PROJECT_NAME):$(VERSION)
	docker push onrik/$(PROJECT_NAME):latest
