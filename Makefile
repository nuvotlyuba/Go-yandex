build: del_app fmt_app build_app

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
	go fmt ./...

docker_build:
	docker-compose build

docker_run_sh:
	docker-compose run --service-ports back sh || exit 0

clean:
	docker-compose down --remove-orphans || exit 0

develop: docker_build docker_run_sh

lint: golangci-lint run

migrate_up:
	migrate -path migrations -database "postgres://postgres:user@localhost:5432/shortener" up

migrate_down:
	migrate -path migrations -database "postgres://postgres:user@localhost:5432/shortener" down

migrate_create:
	migrate create -ext sql -dir migrations create_shortener

mockgen:
	mockgen -destination=mocks/mock_store.go -package=mocks project/apiserver Server
