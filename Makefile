build:
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

all: test lint generate-mocks grpc-proto