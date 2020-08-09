# Makefile for ¡Meterpostigo!

GG=go
GEN_OAPI=$(GG) run github.com/deepmap/oapi-codegen/cmd/oapi-codegen

BIN_DIR=./bin
API_SPEC=./api/server.yml
API=./internal/api/serverapi.gen.go
OUT=$(BIN_DIR)/meterpostigo

all: api test build

$(API): $(API_SPEC)

api: $(API)
	$(GEN_OAPI) \
	--package serverapi \
	-generate types,server \
	-o $(API) \
	$(API_SPEC)

test:
	$(GG) test ./...

build:
	$(GG) build -a -o $(OUT) ./cmd/meterpostigo

clean:
	$(GG) clean
	rm -f $(OUT)
	rm -f $(API)
