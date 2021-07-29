# Copyright (c) 2021 Mobigen JBLIM. All Rights Reserved.

################################################################################
##                             Docker PARAMS                                 ##
################################################################################

## Docker Build Versions
DOCKER_BUILD_IMAGE = golang:1.16.2
DOCKER_BASE_IMAGE = alpine:3.13.2

################################################################################
##                             PROGRAM PARAMS                                 ##
################################################################################

# program name and version info 
TARGET := test
VERSION := v1.0.0
IMAGE ?= repo.iris.tools/iris/web-server:$(VERSION)

################################################################################

GO ?= $(shell command -v go 2> /dev/null)
MACHINE = $(shell uname -m)
GOFLAGS ?= $(GOFLAGS:)
BUILD_TIME := $(shell date -u +%Y%m%d.%H%M%S)
BUILD_HASH := $(shell git rev-parse --short HEAD)

################################################################################

MODULE_NAME := $(shell head -1 go.mod | awk '{print $$2}')
LDFLAGS += -X '$(MODULE_NAME)/common/appdata.Name=$(TARGET)'
LDFLAGS += -X '$(MODULE_NAME)/common/appdata.Version=$(VERSION)'
LDFLAGS += -X '$(MODULE_NAME)/common/appdata.BuildHash=$(BUILD_HASH)'

# Binaries.
TOOLS_BIN_DIR := $(abspath bin)
GO_INSTALL = ./scripts/go_install.sh

MOCKGEN_VER := v1.4.3
MOCKGEN_BIN := mockgen
MOCKGEN := $(TOOLS_BIN_DIR)/$(MOCKGEN_BIN)-$(MOCKGEN_VER)

OUTDATED_VER := master
OUTDATED_BIN := go-mod-outdated
OUTDATED_GEN := $(TOOLS_BIN_DIR)/$(OUTDATED_BIN)

GOVERALLS_VER := master
GOVERALLS_BIN := goveralls
GOVERALLS_GEN := $(TOOLS_BIN_DIR)/$(GOVERALLS_BIN)

GOLINT_VER := master
GOLINT_BIN := golint
GOLINT_GEN := $(TOOLS_BIN_DIR)/$(GOLINT_BIN)

export GO111MODULE=on

## Checks the code style, tests, builds and bundles.
all: check-style dist

## Runs govet and gofmt against all packages.
.PHONY: check-style
check-style: govet lint
	@echo Checking for style guide compliance

## Runs lint against all packages.
.PHONY: lint
lint: $(GOLINT_GEN)
	@echo Running lint
	$(GOLINT_GEN) -set_exit_status ./...
	@echo lint success

## Runs govet against all packages.
.PHONY: vet
govet:
	@echo Running govet
	$(GO) vet ./...
	@echo Govet success

## Builds and thats all :)
.PHONY: dist
dist:	build

.PHONY: build
build: ## Build binary
	@echo Building $(TARGET)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 $(GO) build -ldflags "$(LDFLAGS)" -gcflags all=-trimpath=$(PWD) -asmflags all=-trimpath=$(PWD) \
	     -a -installsuffix cgo -o build/bin/$(TARGET) main.go

build-image:  ## Build the docker image 
	@echo Building $(TARGET) Docker Image
	docker build \
	--build-arg DOCKER_BUILD_IMAGE=$(DOCKER_BUILD_IMAGE) \
	--build-arg DOCKER_BASE_IMAGE=$(DOCKER_BASE_IMAGE) \
	. -f build/Dockerfile -t $(IMAGE) \
	--no-cache


HOME = $(shell pwd)
PROFILE = "prod"
.PHONY: run
run: 
	mkdir -p db &&  \
	HOME=$(HOME) PROFILE=$(PROFILE) ./build/bin/test

# Generate mocks from the interfaces.
.PHONY: mocks
mocks:  $(MOCKGEN)
	go generate ./...

.PHONY: check-modules
check-modules: $(OUTDATED_GEN) ## Check outdated modules
	@echo Checking outdated modules
	$(GO) list -u -m -json all | $(OUTDATED_GEN) -update -direct

.PHONY: goverall
goverall: $(GOVERALLS_GEN) ## Runs goveralls
	$(GOVERALLS_GEN) -coverprofile=coverage.out -service=circle-ci -repotoken ${COVERALLS_REPO_TOKEN} || true

.PHONY: unittest
unittest:
	$(GO) test ./... -v -covermode=count -coverprofile=coverage.out

.PHONY: verify-mocks
verify-mocks:  $(MOCKGEN) mocks
	@if !(git diff --quiet HEAD); then \
		echo "generated files are out of date, run make mocks"; exit 1; \
	fi

## --------------------------------------
## Tooling Binaries
## --------------------------------------

$(MOCKGEN): ## Build mockgen.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/golang/mock/mockgen $(MOCKGEN_BIN) $(MOCKGEN_VER)

$(OUTDATED_GEN): ## Build go-mod-outdated.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/psampaz/go-mod-outdated $(OUTDATED_BIN) $(OUTDATED_VER)

$(GOVERALLS_GEN): ## Build goveralls.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) github.com/mattn/goveralls $(GOVERALLS_BIN) $(GOVERALLS_VER)

$(GOLINT_GEN): ## Build golint.
	GOBIN=$(TOOLS_BIN_DIR) $(GO_INSTALL) golang.org/x/lint/golint $(GOLINT_BIN) $(GOLINT_VER)
