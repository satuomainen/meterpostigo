# Makefile for ¡Meterpostigo!

GG=go
BIN_DIR=./bin
OUT=$(BIN_DIR)/meterpostigo

all: test build

test:
	$(GG) test ./...

build:
	$(GG) build -a -o $(OUT) ./cmd/meterpostigo

clean:
	$(GG) clean
	rm -f $(OUT)
