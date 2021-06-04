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