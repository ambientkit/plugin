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
	go mod tidy

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
	go mod tidy

# Update dependencies of other repos using this repo.
.PHONY: update-children
update-children:
	cd ../ambient-template && go get -u github.com/ambientkit/plugin@$(shell git rev-parse HEAD) && go mod tidy
	cd ../amb && go get -u github.com/ambientkit/plugin@$(shell git rev-parse HEAD) && go mod tidy