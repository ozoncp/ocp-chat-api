build:
	go build -o bin/ocp-chat-api -mod vendor cmd/ocp-chat-api/main.go

run:
	@./bin/ocp-chat-api