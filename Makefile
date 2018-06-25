build: generate
	go build ./cmd/ttouch

install: generate
	go install ./cmd/ttouch

.PHONY: generate
generate:
	tsc
	go generate ./...
