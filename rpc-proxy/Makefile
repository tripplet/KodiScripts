############
# Makefile #
############

# git', 'describe', '--long', '--always', '--tags


GO = go

.PHONY: generate all
generate:
	@$(GO) generate

rpc-proxy-rpi: generate
	@echo Building rpc-proxy for linux/arm
	env GOOS=linux GOARCH=arm $(GO) build -ldflags "-s -w" -trimpath -o rpc-proxy-rpi

rpc-proxy: generate
	@echo Building rpc-proxy
	$(GO) build -ldflags "-s -w" -o rpc-proxy

all: rpc-proxy rpc-proxy-rpi