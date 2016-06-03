.PHONY: install test run prepare

export GOPATH := $(shell while [ ! -d src/github.com ]; do cd ..; done && pwd)

install: prepare
	go install ./cmd/fibonacci
	go install ./cmd/fibonacci_server
	@echo Installed to $(GOPATH)/bin

test: install
	go test ./fibonacci
	$(GOPATH)/bin/fibonacci 10

run: install
	$(GOPATH)/bin/fibonacci_server

prepare:
	@echo GOPATH=$(GOPATH)
	[ -d $(GOPATH)/src/github.com/julienschmidt/httprouter ] || \
		go get github.com/julienschmidt/httprouter

