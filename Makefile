OUT_DIR = _output

export GOFLAGS

default: help

help:
	@echo "Usage: make <target>"
	@echo
	@echo " 'build'          - Build all knitter related binaries(e.g. knitter-manager,knitter-agent,knitter-plugin)"
	@echo " 'test'           - Test knitter with unit test"
	@echo " 'test-e2e'       - Test knitter with e2e test"
	@echo " 'clean'          - Clean artifacts"
	@echo " 'verify'         - Excute static code checking tools(e.g. gofmt,lint,govet)"
	@echo " 'install-extra'  - Install extra tools used by verify(e.g. gometalinter)"


check-gopath:
ifndef GOPATH
        $(error GOPATH is not set)
endif
.PHONY: check-gopath


# Build all knitter components.
#
# Example:
#         make build
#         make all
all build: knitter-manager knitter-agent knitter-plugin
.PHONY: all build

# Build knitter-plugin component
#
# Example:
#        make knitter-plugin
knitter-plugin:
	script/knitter.sh build
.PHONY: knitter-plugin

# Build knitter-manager component
#
# Example:
#         make knitter-manager
knitter-manager:
	script/manager.sh build
.PHONY: knitter-manager

# Build knitter-agent component
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
	@./hack/verify-lint.sh

# Lint knitter code files with golint tool.
#
# Example:
#         make golint
golint: check-gopath
	@echo "checking golint"
	@./hack/verify-golint.sh

# Format knitter code files with gofmt tool.
# 
# Example:
#         make gofmt
gofmt:
	@echo "checking gofmt"
	@./hack/verify-gofmt.sh

# Static check knitter code files 
#
# Example:
#         make govet
govet:
	@echo "checking govet"
	@./hack/verify-govet.sh


# Verify if code is properly organized.
#
# Example:
#         make verify
verify: gofmt lint govet 
.PHONY: verify

# Install travis dependencies
#
# Example:
#   make install-travis
install-travis:
	hack/install-tools.sh
.PHONY: install-travis

# Verify and Test knitter code
#
# Example:
#   make check
check: verify test	
.PHONY: check

# Run unit tests
# Example:
#   make test
#   make test WHAT=pkg/* TESTFLAGS=-v 
test:
	go test -timeout=20m -race ./pkg/... ./knitter-agent/... ./knitter-manager/... ./knitter-plugin/... $(BUILD_TAGS) $(GO_LDFLAGS) $(GO_GCFLAGS) 
.PHONY: test

install-extra: install-gometalinter
.PHONY: install-extra

# install gometailinter tool
# Example:
# make install-gometalinter

install-gometalinter:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install


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
