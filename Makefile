
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin


BRA := $(GOBIN)/bra-v0.0.0-20200517080246-1e3013ecaff8

.PHONY: run
run: $(BRA) ## Build and run web server on filesystem changes. See /.bra.toml for configuration.
	$(BRA) run

.PHONY: build
build: 
	@echo "building project..."
	go build