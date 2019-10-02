include secrets.env
include server.env

BINARY_NAME := server

.PHONY: deps	
deps:
	GO111MODULE=on go get -insecure ./...

.PHONY: build
build:
	cd cmd/server && CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ${BINARY_NAME}

.PHONY: test
test:
	CGO_ENABLED=0 go test -count=1 -v ./...

secrets.env:
	gpg -d -o $@ secrets.env.gpg

secrets.env.gpg: secrets.env
	gpg -c $<