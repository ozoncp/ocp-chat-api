SRC_PATH := /go/src/github.com/ozoncp/ocp-chat-api
DOCKER_IMAGE := repo_addr.com/ocp-development-group/ocp-chat-api

ifeq ($(ENV_FILE),)
ENV_FILE=cmd/ocp-chat-api/conf.env
endif

ifeq ($(TAG),)
TAG=$(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
endif

build:
	go mod tidy
	go mod vendor
	go build -o bin/ocp-chat-api ./cmd/ocp-chat-api

run: build
	@./bin/ocp-chat-api

test:
	go test -test.v -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out

lint:
	golangci-lint run -v

generate-mocks:
	go generate ./...

grpc-proto:
	protoc --proto_path=pkg/chat_api --go_out=pkg/chat_api  --go_opt=paths=source_relative --go-grpc_out=pkg/chat_api --go_opt=paths=source_relative ocp-chat-api.proto

docker-build:
	docker build --build-arg SRC_PATH=${SRC_PATH} -t ${DOCKER_IMAGE}:${TAG} .

docker-run:
	docker run -p 8080:80 --env-file ${ENV_FILE} ${DOCKER_IMAGE}:${TAG}

docker-compose-up: generate-mocks grpc-proto docker-build
	TAG=${TAG} docker-compose up --remove-orphans

all: generate-mocks grpc-proto lint test docker-build docker-run
