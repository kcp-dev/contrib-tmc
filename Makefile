# Copyright 2021 The KCP Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# We need bash for some conditional logic below.
SHELL := /usr/bin/env bash -e

REPO ?= quay.io/kcp-dev/contrib-tmc
KO_DOCKER_REPO ?= ${REPO}

GO_INSTALL = ./hack/go-install.sh

TOOLS_DIR=hack/tools
TOOLS_GOBIN_DIR := $(abspath $(TOOLS_DIR))
GOBIN_DIR=$(abspath ./bin)
PATH := $(GOBIN_DIR):$(TOOLS_GOBIN_DIR):$(PATH)

ARCH := $(shell go env GOARCH)
OS := $(shell go env GOOS)

LDFLAGS := \
	-extldflags '-static'

ldflags:
	@echo $(LDFLAGS)

.PHONY: require-%
require-%:
	@if ! command -v $* 1> /dev/null 2>&1; then echo "$* not found in ${PATH}"; exit 1; fi

all: build
.PHONY: all

# Detect the path used for the install target
ifeq (,$(shell go env GOBIN))
INSTALL_GOBIN=$(shell go env GOPATH)/bin
else
INSTALL_GOBIN=$(shell go env GOBIN)
endif

CONTROLLER_GEN_VER := v0.10.0
CONTROLLER_GEN_BIN := controller-gen
CONTROLLER_GEN := $(TOOLS_DIR)/$(CONTROLLER_GEN_BIN)-$(CONTROLLER_GEN_VER)
export CONTROLLER_GEN # so hack scripts can use it

GOLANGCI_LINT_VER := v1.50.1
GOLANGCI_LINT_BIN := golangci-lint
GOLANGCI_LINT := $(TOOLS_GOBIN_DIR)/$(GOLANGCI_LINT_BIN)-$(GOLANGCI_LINT_VER)

OPENSHIFT_GOIMPORTS_VER := c72f1dc2e3aacfa00aece3391d938c9bc734e791
OPENSHIFT_GOIMPORTS_BIN := openshift-goimports
OPENSHIFT_GOIMPORTS := $(TOOLS_DIR)/$(OPENSHIFT_GOIMPORTS_BIN)-$(OPENSHIFT_GOIMPORTS_VER)
export OPENSHIFT_GOIMPORTS # so hack scripts can use it

CODE_GENERATOR_VER := v2.1.0
CODE_GENERATOR_BIN := code-generator
CODE_GENERATOR := $(TOOLS_GOBIN_DIR)/$(CODE_GENERATOR_BIN)-$(CODE_GENERATOR_VER)
export CODE_GENERATOR # so hack scripts can use it

STATICCHECK_VER := 2022.1
STATICCHECK_BIN := staticcheck
STATICCHECK := $(TOOLS_GOBIN_DIR)/$(STATICCHECK_BIN)-$(STATICCHECK_VER)

KCP_APIGEN_VER := v0.20.0
KCP_APIGEN_BIN := apigen
KCP_APIGEN_GEN := $(TOOLS_DIR)/$(KCP_APIGEN_BIN)-$(KCP_APIGEN_VER)
export KCP_APIGEN_GEN # so hack scripts can use it

LOGCHECK_VER := v0.4.0
LOGCHECK_BIN := logcheck
LOGCHECK := $(TOOLS_GOBIN_DIR)/$(LOGCHECK_BIN)-$(LOGCHECK_VER)
export LOGCHECK # so hack scripts can use it

ifdef ARTIFACT_DIR
GOTESTSUM_ARGS += --junitfile=$(ARTIFACT_DIR)/junit.xml
endif

GO_TEST = go test
ifdef USE_GOTESTSUM
GO_TEST = $(GOTESTSUM) $(GOTESTSUM_ARGS) --
endif

tools: $(GOLANGCI_LINT) $(CONTROLLER_GEN) $(KCP_APIGEN_GEN) $(OPENSHIFT_GOIMPORTS)
.PHONY: tools

$(CONTROLLER_GEN):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/controller-gen $(CONTROLLER_GEN_BIN) $(CONTROLLER_GEN_VER)

$(GOLANGCI_LINT):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) github.com/golangci/golangci-lint/cmd/golangci-lint $(GOLANGCI_LINT_BIN) $(GOLANGCI_LINT_VER)

$(CODE_GENERATOR):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) github.com/kcp-dev/code-generator/v2 $(CODE_GENERATOR_BIN) $(CODE_GENERATOR_VER)

$(KCP_APIGEN_GEN):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) github.com/kcp-dev/kcp/sdk/cmd/apigen $(KCP_APIGEN_BIN) $(KCP_APIGEN_VER)

$(OPENSHIFT_GOIMPORTS):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) github.com/openshift-eng/openshift-goimports $(OPENSHIFT_GOIMPORTS_BIN) $(OPENSHIFT_GOIMPORTS_VER)

$(STATICCHECK):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) honnef.co/go/tools/cmd/staticcheck $(STATICCHECK_BIN) $(STATICCHECK_VER)

$(LOGCHECK):
	GOBIN=$(TOOLS_GOBIN_DIR) $(GO_INSTALL) sigs.k8s.io/logtools/logcheck $(LOGCHECK_BIN) $(LOGCHECK_VER)

crds: $(CONTROLLER_GEN) $(YAML_PATCH)
	./hack/update-codegen-crds.sh
.PHONY: crds

lint: $(GOLANGCI_LINT) $(STATICCHECK) $(LOGCHECK)
	$(GOLANGCI_LINT) -v run ./...
	$(STATICCHECK) -checks ST1019,ST1005 ./...
	./hack/verify-contextual-logging.sh
.PHONY: lint


codegen: crds $(CODE_GENERATOR)
	go mod download
	./hack/update-codegen-clients.sh
	$(MAKE) imports
.PHONY: codegen

# Note, running this locally if you have any modified files, even those that are not generated,
# will result in an error. This target is mostly for CI jobs.
.PHONY: verify-codegen
verify-codegen:
	if [[ -n "${GITHUB_WORKSPACE}" ]]; then \
		mkdir -p $$(go env GOPATH)/src/github.com/kcp-dev; \
		ln -s ${GITHUB_WORKSPACE} $$(go env GOPATH)/src/github.com/kcp-dev/kcp; \
	fi

	$(MAKE) codegen

	if ! git diff --quiet HEAD; then \
		git diff; \
		echo "You need to run 'make codegen' to update generated files and commit them"; \
		exit 1; \
	fi

.PHONY: imports
imports: $(OPENSHIFT_GOIMPORTS)
	$(OPENSHIFT_GOIMPORTS) -m github.com/kcp-dev/contrib-tmc

all: build
.PHONY: all

build: require-jq require-go require-git
	GOOS=$(OS) GOARCH=$(ARCH) CGO_ENABLED=0 go build $(BUILDFLAGS) -ldflags="$(LDFLAGS)" -o bin ./cmd/...
	echo "To use the binaries, add them to your PATH."
	export PATH=$$PATH:$(PWD)/bin
.PHONY: build

test:
	go test -v -failfast `go list ./... | egrep -v /test/` -coverprofile=profile.cov
.PHONY: test

generate: codegen crds
	go generate ./...
.PHONY: generate


E2E_PARALLELISM ?=
ifdef E2E_PARALLELISM
PARALLELISM_ARG = -p $(E2E_PARALLELISM) -parallel $(E2E_PARALLELISM)
endif

COMPLETE_SUITES_ARG =
ifdef SUITES
SUITES_ARG = --suites $(SUITES)
COMPLETE_SUITES_ARG = -args $(SUITES_ARG)
endif

.PHONY: test-e2e
ifdef USE_GOTESTSUM
test-e2e: $(GOTESTSUM)
endif
test-e2e: TEST_ARGS ?=
test-e2e: WHAT ?= ./test/e2e...
test-e2e: build
	UNSAFE_E2E_HACK_DISABLE_ETCD_FSYNC=true NO_GORUN=1 GOOS=$(OS) GOARCH=$(ARCH) \
		$(GO_TEST) -race $(COUNT_ARG) $(PARALLELISM_ARG) $(WHAT) $(TEST_ARGS) $(COMPLETE_SUITES_ARG)

images:
	KO_DOCKER_REPO=${KO_DOCKER_REPO} ko build --sbom=none -B --platform=linux/amd64 -t latest ./cmd/*
