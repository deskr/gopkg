TEST = $$(go list ./... | grep -v '/vendor/')

all: test

deps:
	go get github.com/kardianos/govendor
	-govendor init
	-govendor add +outside
	-govendor sync +outside
	-govendor fetch +outside

test:
	@./scripts/test unit

test-race:
	@./scripts/test unit-race

.PHONY: all deps test
