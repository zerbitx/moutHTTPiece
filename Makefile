# Default to linux, for build boxes/prod.
GOOS?=darwin
GOARCH?=amd64
DOCKER_REPO?=zerbitx
BUILD_TAG?=latest

build:
	mkdir -p bin
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/mouthttpiece main.go

build_linux: GOOS = linux
build_linux: build

docker: build_linux
	docker build -t $(DOCKER_REPO)/mouthttpiece:$(BUILD_TAG) .

docker_publish: docker
	docker push $(DOCKER_REPO)/mouthttpiece:$(BUILD_TAG)
