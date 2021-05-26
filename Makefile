build:
	go build -o bin/ocp-chat-api -mod vendor cmd/ocp-chat-api/main.go

run:
	@./bin/ocp-chat-api

test:
	go test -test.v ./...

cover: test
	go tool cover -html=coverage.out

lint:
	golangci-lint run -v