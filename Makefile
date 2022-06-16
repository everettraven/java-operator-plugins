lint:
	@./hack/check-license.sh
	@go fmt ./...

test:
	@go test -coverprofile=coverage.out -covermode=count -short ./...

.PHONY: test lint

.PHONY: generate
generate:
	go run ./hack/generate/samples/generate.go
