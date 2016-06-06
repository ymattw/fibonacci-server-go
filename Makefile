.PHONY: install test run prepare

VENDOR := $(CURDIR)/vendor

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
	# golang.org/x/crypto/ssh/terminal blocked by GFW!
	[ -d $(VENDOR)/src/golang.org/x/crypto/.git ] || \
		git clone --depth 1 https://github.com/golang/crypto.git \
			$(VENDOR)/src/golang.org/x/crypto
