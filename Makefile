build:
	rm -f shortener && go test ./... && go fmt ./... && staticcheck ./... && go build -v ./cmd/shortener
t:
	go mod tidy

test:
	go test ./...

test-clean:
	go clean -testcache

client:
	go run ./cmd/client

start:
	go run ./cmd/shortener

fmt:
	go fmt ./... && staticcheck ./...
