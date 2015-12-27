TEST = $$(GO15VENDOREXPERIMENT=1 go list ./... | grep -v '/vendor/')

all: test

deps:
	@export GO15VENDOREXPERIMENT=1
	go get github.com/FiloSottile/gvt
	gvt rebuild
	go install -v

test:
	@export GO15VENDOREXPERIMENT=1
	go vet $(TEST)
	go test $(TEST) $(TESTARGS) -cover -timeout=30s -parallel=4

test-race:
	go test $(TEST) -race

.PHONY: all deps test
