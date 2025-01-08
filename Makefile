.PHONY: build
build:
	CGO_ENABLED=0 go build -tags=viper_bind_struct -o ./bin/server ./cmd/server

.PHONY: run
run: build
	./bin/server

.PHONY: docker-run
docker-run:
	docker build -t calculate-api-server:latest .
	docker run --rm -p 8080:8080 calculate-api-server:latest

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -v ./internal/... ./pkg/...

.PHONY: test-cover
test-cover:
	mkdir -p coverage \
	&& go test ./internal/... ./pkg/... -coverprofile=coverage/cover \
	&& go tool cover -html=coverage/cover -o coverage/cover.html
