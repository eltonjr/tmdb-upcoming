include server.env
-include secrets.env
export

.DEFAULT_GOAL := help

BINARY_NAME := server
CMD_PATH := cmd/server

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "Usage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: deps	
deps: ## Install go dependencies based on go mod
	GO111MODULE=on go get -insecure ./...

.PHONY: build
build: ## Compiles the system's binary
	cd ${CMD_PATH} && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ${BINARY_NAME}

.PHONY: run
run: ## Run the system locally
	go run ${CMD_PATH}/${BINARY_NAME}.go

.PHONY: test
test: ## Run unit tests
	CGO_ENABLED=0 go test -count=1 -v ./...

.PHONY: secrets-unlock
secrets-unlock: ## Unlock the secret variables
	gpg -d -o secrets.env secrets.env.gpg

.PHONY: secrets-lock
secrets-lock: ## Lock the secret variables
	gpg -c secrets.env

.PHONY: ops-build
ops-build: ## Compiles the system using docker
	docker build -t tmdb-server .

.PHONY: ops-run
ops-run: ## Runs the system using docker
	# docker run -d --name tmdb-server --env-file server.env --env-file secrets.env -p 9000:9000 tmdb-server
	docker run -d --name tmdb-server --env-file server.env --env-file secrets.env --net host tmdb-server