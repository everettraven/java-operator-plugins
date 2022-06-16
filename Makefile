lint:
	@./hack/check-license.sh
	@go fmt ./...

test:
	@go test -coverprofile=coverage.out -covermode=count -short ./...

.PHONY: test lint

.PHONY: generate
generate: # generate the testdata samples
	go run ./hack/generate/samples/generate.go
