include secrets.env
include server.env
export

BINARY_NAME := server
CMD_PATH := cmd/server

.PHONY: deps	
deps:
	GO111MODULE=on go get -insecure ./...

.PHONY: build
build:
	cd ${CMD_PATH} && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ${BINARY_NAME}

.PHONY: run
run:
	#./${CMD_PATH}/${BINARY_NAME}
	go run ${CMD_PATH}/${BINARY_NAME}.go

.PHONY: test
test:
	CGO_ENABLED=0 go test -count=1 -v ./...

secrets.env:
	gpg -d -o $@ secrets.env.gpg

secrets.env.gpg: secrets.env
	gpg -c $<