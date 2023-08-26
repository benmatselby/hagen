NAME := hagen
DOCKER_PREFIX = benmatselby
DOCKER_RELEASE ?= latest
DOCKER_PLATFORM ?= --platform=amd64

.PHONY: explain
explain:
	### Welcome
	#
	#    _   _   _   _   _
	#   / \ / \ / \ / \ /
	#  ( h | a | g | e | n |
	#   \_/ \_/ \_/ \_/ \_/
	#
	#
	### Installation
	#
	# $$ make all
	#
	### Targets
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

GITCOMMIT := $(shell git rev-parse --short HEAD)

.PHONY: clean
clean: ## Clean the local dependencies
	rm -fr vendor

.PHONY: install
install: ## Install the local dependencies
	go install github.com/golang/mock/mockgen@master
	go install github.com/securego/gosec/cmd/gosec@master
	go install golang.org/x/lint/golint@master
	go get ./...

.PHONY: lint
lint: ## Lint the code
	golangci-lint run

.PHONY: security
security: ## Inspect the code
	gosec ./...

.PHONY: build
build: ## Build the application
	go build .

.PHONY: static
static: ## Build the application
	CGO_ENABLED=0 go build -ldflags "-extldflags -static -X github.com/benmatselby/$(NAME)/version.GITCOMMIT=$(GITCOMMIT)" -o $(NAME) .

.PHONY: mocks
mocks:
	go generate -x ./...

.PHONY: test
test: mocks ## Run the unit tests
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: test-cov
test-cov: test ## Run the unit tests with coverage
	go tool cover -html=coverage.out

.PHONY: all ## Run everything
all: clean install vet build test

.PHONY: static-all ## Run everything
static-all: clean install static test


###
# Docker
###

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_PREFIX)/$(NAME):$(DOCKER_RELEASE) $(DOCKER_PLATFORM) .

.PHONY: docker-push
docker-push:
	docker push $(DOCKER_PREFIX)/$(NAME):$(DOCKER_RELEASE)

.PHONY: docker-scan
docker-scan: ## Scan the docker image
	docker scout recommendations $(DOCKER_PREFIX)/$(NAME)
	docker scout quickview $(DOCKER_PREFIX)/$(NAME)
