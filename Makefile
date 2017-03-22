all: test

test:
	@./scripts/test unit

test-race:
	@./scripts/test unit-race

.PHONY: all test
