build:
	go build -o bin/ocp-chat-api cmd/ocp-chat-api/main.go

run: build
	@./bin/ocp-chat-api

test:
	go test -test.v ./...

cover: test
	go tool cover -html=coverage.out

lint:
	golangci-lint run -v

generate-mocks:
	go generate ./...

grpc-proto:
	protoc --proto_path=api/ocp-chat-api --go_out=vendor.protogen --go_opt=paths=import --go-grpc_out=vendor.protogen --go_opt=paths=import ocp-chat-api.proto