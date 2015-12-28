TEST = $$(GO15VENDOREXPERIMENT=1 go list ./... | grep -v '/vendor/')

all: test

deps:
	@export GO15VENDOREXPERIMENT=1
	go get github.com/FiloSottile/gvt
	gvt rebuild
	go install -v

test:
	@./test unit

test-race:
	@./test unit-race

.PHONY: all deps test
