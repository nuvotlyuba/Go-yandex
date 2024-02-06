build:
		go build -v ./cmd/shortener
t:
	go mod tidy

test:
	go test ./...

test-clean:
	go clean -testcache

client:
	go run ./cmd/client
