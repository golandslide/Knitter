OUT_DIR = _output

export GOFLAGS

default: help

help:
	@echo "Usage: make <target>"
	@echo
	@echo " 'build'          - Build all knitter related binaries(e.g. knitter-manager,knitter-agent,knitter-plugin)"
	@echo " 'test-ut'        - Test knitter with unit test"
	@echo " 'test-e2e'       - Test knitter with e2e test"
	@echo " 'clean'          - Clean artifacts"
	@echo " 'verify'         - Execute the source code verification tools(e.g. gofmt,lint,govet)"
	@echo " 'install-extra'  - Install tools used by verify(e.g. gometalinter)"



check-gopath:
ifndef GOPATH
        $(error GOPATH is not set)
endif
.PHONY: check-gopath


# Test travis ci with test code
#
# Example:
#     make test-build
#     make test-verify
#     make test-test    
test-build:
	go build ./cmd/test.go
.PHONY: test-build

test-verify:verify
.PHONY: test-verify

test-ut-test:
	go test -v ./pkg/common
.PHONY: test-ut-test

# Build code.
#
# Args:
#   GOFLAGS: Extra flags to pass to 'go' when building.
#
# Example:
#         make build
#         make all
all build: knitter-manager knitter-agent knitter-plugin
.PHONY: all build

# Build knitter-plugin
#
# Example:
#        make knitter-plugin
knitter-plugin:
	script/knitter.sh build
.PHONY: knitter-plugin

# Build knitter-manager
#
# Example:
#         make knitter-manager
knitter-manager:
	script/manager.sh build
.PHONY: knitter-manager
# Build knitter-agent
#
# Example:
#         make knitter-agent
knitter-agent:
	script/agent.sh build
.PHONY: knitter-agent

# Lint knitter code files. note that this lint process handled by gometalinter tools.
# link here (github.com/alecthomas/gometalinter)
# If users only need simple lint process, please run command 'make golint'.
# Example:
#         make lint
lint: check-gopath
	@echo "checking lint"
	hack/verify-lint.sh

# Lint knitter code files with golint tool.
#
# Example:
#         make golint
golint: check-gopath
	@echo "checking golint"
	hack/verify-golint.sh

# Format knitter code files with gofmt tool.
# 
# Example:
#         make gofmt
gofmt:
	@echo "checking gofmt"
	hack/verify-gofmt.sh

# Static check knitter code files 
#
# Example:
#         make govet
govet:
	@echo "checking govet"
	hack/verify-govet.sh


# verify whether code is properly organized.
#
# Example:
#         make verify
verify: gofmt golint govet
.PHONY: verify

# strict-verify verify code is properly organized.
#
# Example:
#         make strict-verify
strict-verify:gofmt lint govet
.PHONY:strict-verify


check: verify test	
.PHONY: check

# Run unit tests
# Example:
#   make test
#   make test WHAT=pkg/docker TESTFLAGS=-v 
test-ut:
	go test -timeout=20m -race ./pkg/... ./knitter-agent/... ./knitter-manager/... ./knitter-plugin/... $(BUILD_TAGS) $(GO_LDFLAGS) $(GO_GCFLAGS) 
.PHONY: test-ut

install-tools:install-gometalinter
.PHONY: install-tools

# install gometailinter tool
# Example:
# make install-gometalinter
install-gometalinter:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
# Deploy a dind kubernetes cluser
# note: Leverage Mirantis kubeadm-dind-cluster
# Example:
# make deploy-dind-k8s
deploy-dind-k8s:
	./hack/dind-ci.sh up
.PHONY: deploy-dind-k8s

# Run knitter e2e test case
# Example:
# make test-e2e
test-e2e:
	./hack/run-robotframe-e2e.sh
.PHONY: test-e2e

# Just for test
probe-test:
	./hack/probe-test.sh
.PHONY: probe-test

# Remove all build artifacts.
#
# Example:
#   make clean
clean:
	rm -rf $(OUT_DIR)
.PHONY: clean

# Build the release.
#
# Example:
#   make release
release: clean
	hack/build-release.sh
	hack/extract-release.sh
.PHONY: release
