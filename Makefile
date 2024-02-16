build: del_app test_app build_app

build_app:
	go build -v ./cmd/shortener

del_app:
	rm -f shortener

test_app:
	go clean -testcache && go test ./...

start:
	go run ./cmd/shortener

fmt_app:
	go fmt ./... && staticcheck ./...

docker_build:
	docker-compose build

docker_run_sh:
	docker-compose run --service-ports back sh || exit 0

clean:
	docker-compose down --remove-orphans || exit 0

develop: docker_build docker_run_sh

