OS   := $(shell uname | awk '{print tolower($$0)}')
ARCH := $(shell case $$(uname -m) in (x86_64) echo amd64 ;; (aarch64) echo arm64 ;; (*) echo $$(uname -m) ;; esac)

BUF_VERSION                     := 1.4.0
PROTOC_GEN_GO_VERSION           := 1.28.0
PROTOC_GEN_GO_GRPC_VERSION      := 1.2.0
PROTOC_GEN_GRPC_GATEWAY_VERSION := 2.10.0

BIN_DIR := $(shell pwd)/bin

BUF                     := $(abspath $(BIN_DIR)/buf)
PROTOC_GEN_GO           := $(abspath $(BIN_DIR)/protoc-gen-go)
PROTOC_GEN_GO_GRPC      := $(abspath $(BIN_DIR)/protoc-gen-go-grpc)
PROTOC_GEN_GRPC_GATEWAY := $(abspath $(BIN_DIR)/protoc-gen-grpc-gateway)

buf: $(BUF)
$(BUF):
	@curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" -o $(BUF) && chmod +x $(BUF)

protoc-gen-go: $(PROTOC_GEN_GO)
$(PROTOC_GEN_GO):
	@curl -sSL https://github.com/protocolbuffers/protobuf-go/releases/download/v$(PROTOC_GEN_GO_VERSION)/protoc-gen-go.v$(PROTOC_GEN_GO_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv protoc-gen-go

protoc-gen-go-grpc: $(PROTOC_GEN_GO_GRPC)
$(PROTOC_GEN_GO_GRPC):
	@curl -sSL https://github.com/grpc/grpc-go/releases/download/cmd%2Fprotoc-gen-go-grpc%2Fv$(PROTOC_GEN_GO_GRPC_VERSION)/protoc-gen-go-grpc.v$(PROTOC_GEN_GO_GRPC_VERSION).$(OS).$(ARCH).tar.gz | tar -C $(BIN_DIR) -xzv ./protoc-gen-go-grpc

protoc-gen-grpc-gateway: $(PROTOC_GEN_GRPC_GATEWAY)
$(PROTOC_GEN_GRPC_GATEWAY):
	@curl -sSL "https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${PROTOC_GEN_GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-v${PROTOC_GEN_GRPC_GATEWAY_VERSION}-${OS}-$(shell uname -m)" -o $(PROTOC_GEN_GRPC_GATEWAY) && chmod +x $(PROTOC_GEN_GRPC_GATEWAY)

.PHONY: gen-proto
gen-proto: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY)
	@$(BUF) generate \
		--path ./echoserver/ \
		--path ./echocaller/

.PHONY: echoserver
echoserver:
	@CGO_ENABLED=0 go build \
		-a \
		-trimpath \
		-o ./bin/echoserver \
		-ldflags "-s -w -extldflags -static" \
		./echoserver

.PHONY: echocaller
echocaller:
	@CGO_ENABLED=0 go build \
		-a \
		-trimpath \
		-o ./bin/echocaller \
		-ldflags "-s -w -extldflags -static" \
		./echocaller
