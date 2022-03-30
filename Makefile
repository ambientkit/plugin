# This Makefile is an easy way to run common operations.
# Execute commands like this:
# * make
# * make update
# * make test

# Load the environment variables.
-include .env

.PHONY: default
default: start

################################################################################
# Dependency management
################################################################################

# Update Ambient dependencies.
.PHONY: update
update: update-ambient tidy

# Update Ambient dependency. Requires the repo to be local and in the same folder.
.PHONY: update-ambient
update-ambient:
	go get -u github.com/ambientkit/ambient@$(shell cd ../ambient && git rev-parse HEAD)

# Update all Go dependencies.
.PHONY: update-all
update-all: update-all-go tidy

# Update all Go dependencies.
.PHONY: update-all-go
update-all-go:
	go get -u -f -d ./...

# Run go mod tidy.
.PHONY: tidy
tidy:
	go mod tidy -compat=1.17

################################################################################
# gRPC
################################################################################

# Build the plugins.
.PHONY: build-plugins
build-plugins:
	go build -o pluginmain pkg/grpctestutil/testingdata/cmd/server/main.go
	cd ./pkg/grpctestutil/testingdata/plugin/hello/cmd/plugin && go build -o ambplugin
	cd ./generic/bearblog/cmd/plugin && go build -o ambplugin
	cd ./generic/bearcss/cmd/plugin && go build -o ambplugin
	cd ./generic/pluginmanager/cmd/plugin && go build -o ambplugin

# Start the build and run process for grpc.
.PHONY: start
start: build-plugins
	./pluginmain

# Test the all the code for short (no response time testing).
.PHONY: test
test: build-plugins
	go test -race -short ./...

# Test the gRPC code.
.PHONY: test-grpc
test-grpc: build-plugins
	go test -race pkg/grpctestutil/grpcp_test.go

# Test the response times.
.PHONY: test-response
test-response: #build-plugins
	-go clean -testcache && go test -v -run ^TestKPIStandard github.com/ambientkit/plugin/pkg/grpctestutil
	-go clean -testcache && go test -v -run ^TestKPIGRPC github.com/ambientkit/plugin/pkg/grpctestutil