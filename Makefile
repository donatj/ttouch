.PHONY: build
build: generate
	go build ./cmd/ttouch

.PHONY: install
install: generate
	go install ./cmd/ttouch

.PHONY: generate
generate:
	npm ci && npx tsc
	go generate ./...

.PHONY: test
test:
	go test ./...