.PHONY: install test run prepare

VENDOR := $(CURDIR)/_vendor

export GOPATH := $(VENDOR):$(PWD)

install: prepare
	go install ./src/cmd/fibonacci
	go install ./src/cmd/fibonacci_server

test: install
	go test ./src/fibonacci
	bin/fibonacci 10

run: install
	bin/fibonacci_server

prepare:
	@echo GOPATH=$(GOPATH)
	go get github.com/julienschmidt/httprouter
	# Blocked by GFW!
	go get golang.org/x/crypto/ssh/terminal

