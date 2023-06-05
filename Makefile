GOOS :=
GOARCH :=

INPUT_DIR := cmd/main.go
OUTPUT_DIR := ./bin

GO_ENV = GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH)

ifeq ($(GOOS),)
UNAME_GOOS := $(shell uname -s)
ifeq ($(UNAME_GOOS),Linux)
	GOOS = linux
endif
ifeq ($(UNAME_GOOS),Darwin)
	GOOS = darwin
endif
endif

ifeq ($(GOARCH),)
UNAME_GOARCH := $(shell uname -m)
ifeq ($(UNAME_GOARCH),x86_64)
	GOARCH = amd64
endif
ifeq ($(UNAME_GOARCH),i386)
	GOARCH = amd64
endif
ifneq ($(filter arm%,$(UNAME_GOARCH)),)
	GOARCH = arm64
endif
endif

release:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/app $(INPUT_DIR)
clean:
	-rm -rf $(OUTPUT_DIR)
run:
	$(GO_ENV) go run $(INPUT_DIR) -cfg config/config.yaml
