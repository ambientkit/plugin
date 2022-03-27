# This Makefile is an easy way to run common operations.
# Execute commands like this:
# * make

# Load the environment variables.
-include .env

.PHONY: default
default: amb

################################################################################
# Dependency management
################################################################################

# Update Go dependencies.
.PHONY: update
update:
	go get -u -f -d ./...
	go mod tidy -compat=1.17

# Run go mod tidy.
.PHONY: tidy
tidy:
	go mod tidy -compat=1.17

# Pass in ARGS.
# https://stackoverflow.com/a/14061796
ifeq (update-ambient,$(firstword $(MAKECMDGOALS)))
  ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(ARGS):;@:)
endif

# Update Ambient dependency.
.PHONY: update-ambient
update-ambient:
	go get -u github.com/ambientkit/ambient@${ARGS}
	go mod tidy -compat=1.17

# Update dependencies of other repos using this repo.
.PHONY: update-children
update-children:
	cd ../ambient-template && go get github.com/ambientkit/plugin@$(shell git rev-parse HEAD) && go mod tidy -compat=1.17
	cd ../amb && go get github.com/ambientkit/plugin@$(shell git rev-parse HEAD) && go mod tidy -compat=1.17

################################################################################
# gRPC
################################################################################

# Build the plugins.
.PHONY: build-plugins
build-plugins:
	go build -o pluginmain pkg/grpctestutil/testingdata/cmd/server/main.go
	@cd ./pkg/grpctestutil/testingdata/plugin/hello/cmd/plugin && go build -o ambplugin
	@cd ./generic/bearblog/cmd/plugin && go build -o ambplugin
	@cd ./generic/bearcss/cmd/plugin && go build -o ambplugin
	@cd ./generic/pluginmanager/cmd/plugin && go build -o ambplugin

# Start the build and run process for grpc.
.PHONY: start
start: build-plugins
	./pluginmain

# Test the gRPC code.
.PHONY: test
test: build-plugins
	go test -race pkg/grpctestutil/grpcp_test.go

# Test all the code.
.PHONY: test-all
test-all: build-plugins
	go test -race ./...