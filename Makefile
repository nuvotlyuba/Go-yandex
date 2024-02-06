build:
		go build -v ./cmd/shortener
t:
	go mod tidy

test:
	go test ./...
