.PHONY: emlog

emlog:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emlog/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
	@go build -buildmode=plugin -o emlog.so emlog/*
