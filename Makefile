# -*- mode: Makefile-gmake -*-

SHELL := bash

TOP_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

BUILD_DIR := build

COVERAGE_DIR := $(BUILD_DIR)/coverage

all: test

$(BUILD_DIR):
	@mkdir -p $@

$(COVERAGE_DIR):
	@mkdir -p $@

clean:
	rm -rf $(BUILD_DIR)

fmt:
	go fmt ./...

vet:
	go vet -mod=vendor ./...

test: $(COVERAGE_DIR)
	go test -mod=vendor -v -coverprofile=$(TOP_DIR)/$(COVERAGE_DIR)/coverage.out -covermode=atomic github.com/percolate/retry

.PHONY: clean vet fmt test
