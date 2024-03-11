build: del_app build_app

build_app:
	go build -v ./cmd/shortener

del_app:
	rm -f shortener

test_app:
	go clean -testcache && go test ./...

start:
	go run ./cmd/shortener

flags:
	go run ./cmd/shortener -a :3333 -b http://localhost:3333

fmt_app:
	go fmt ./... && staticcheck ./...

docker_build:
	docker-compose build

docker_run_sh:
	docker-compose run --service-ports back sh || exit 0

clean:
	docker-compose down --remove-orphans || exit 0

develop: docker_build docker_run_sh

lint: golangci-lint run
