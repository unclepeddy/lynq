# Ensure GOPATH
GOPATH ?= $(HOME)/go

# List all our actual files, excluding vendor
GOFILES ?= $(shell go list $(TEST) | grep -v /vendor/)

# Tags specific for building
GOTAGS ?=

# List of ldflags
LD_FLAGS ?= \
	-s \
	-w

# Path to Terraform plugins
PLUGIN_PATH ?= "${HOME}/.terraform.d/plugins"

# deps updates all dependencies
deps:
	@dep ensure -update
	@dep prune

# dev builds and installs the plugin into ~/.terraform.d
dev:
	@mkdir -p "${PLUGIN_PATH}"
	@go build \
		-ldflags "${LD_FLAGS}" \
		-tags "${GOTAGS}" \
		-o "${PLUGIN_PATH}/terraform-provider-googlecalendar"
