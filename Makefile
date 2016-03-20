TEST = $$(go list ./... | grep -v '/vendor/')

all: test

deps:
	go get github.com/kardianos/govendor
	govendor init
	govendor add +external

test:
	@./test unit

test-race:
	@./test unit-race

.PHONY: all deps test
