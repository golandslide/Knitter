OUT_DIR = _output

export GOFLAGS

default: help

help:
	@echo "Usage: make <target>"
	@echo
#	@echo " * 'install'        - Install binaries to system locations"
	@echo " * 'build'          - Build all knitter related binaries(e.g. knitter-manager,knitter-agent,knitter-plugin)"
	@echo " * 'test'           - Test knitter with unit test"
	@echo " * 'test-e2e'       - Test knitter with e2e test"
	@echo " * 'clean'          - Clean artifacts"
	@echo " * 'verify'         - Execute the source code verification tools(e.g. gofmt,lint,govet)"
	@echo " * 'install.tools'  - Install tools used by verify(e.g. gometalinter)"



check-gopath:
ifndef GOPATH
        $(error GOPATH is not set)
endif
.PHONY: check-gopath


# Build code.
#
# Args:
#   GOFLAGS: Extra flags to pass to 'go' when building.
#
# Example:
#   make
#   make all
all build: knitter-manager knitter-agent knitter-plugin
.PHONY: all build

knitter-plugin:
	script/knitter.sh build
.PHONY: knitter-plugin

knitter-manager:
	script/manager.sh build
.PHONY: knitter-manager

knitter-agent:
	script/agent.sh build
.PHONY: knitter-agent


lint: check-gopath
	@echo "checking lint"
	@./hack/verify-lint.sh

gofmt:
	@echo "checking gofmt"
	@./hack/verify-gofmt.sh

golint: check-gopath
	@echo "checking golint"
	@./hack/verify-golint.sh

govet:
	@echo "checking govet"
	@./hack/verify-govet.sh


# Verify if code is properly organized.
#
# Example:
#   make verify
verify: gofmt lint govet 
.PHONY: verify

# Install travis dependencies
#
# Example:
#   make install-travis
install-travis:
	hack/install-tools.sh
.PHONY: install-travis

# Build and run unit tests
#
# Args:
#   WHAT: Directory names to test.  All *_test.go files under these
#     directories will be run.  If not specified, "everything" will be tested.
#   TESTS: Same as WHAT.
#   GOFLAGS: Extra flags to pass to 'go' when building.
#   TESTFLAGS: Extra flags that should only be passed to hack/test-go.sh
#
# Example:
#   make check
#   make test
#   make check WHAT=pkg/docker TESTFLAGS=-v
check: verify test	
.PHONY: check

# Run unit tests
# Example:
#   make test
#   make test WHAT=pkg/docker TESTFLAGS=-v 
test:
	go test -timeout=20m -race ./pkg/... ./knitter-agent/... ./knitter-manager/... ./knitter-plugin/... $(BUILD_TAGS) $(GO_LDFLAGS) $(GO_GCFLAGS) 
.PHONY: test

install-tools: install-gometalinter
.PHONY: install-tools

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
