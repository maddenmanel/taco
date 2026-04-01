APP_NAME := taco
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
VERSION_PKG := github.com/maddenmanel/taco/cmd
LDFLAGS := -s -w -X $(VERSION_PKG).Version=$(VERSION)
BUILD_DIR := dist

.PHONY: build clean all

build:
	go build -ldflags "$(LDFLAGS)" -o $(APP_NAME) .

all: clean
	@mkdir -p $(BUILD_DIR)
	GOOS=linux   GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 .
	GOOS=linux   GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-linux-arm64 .
	GOOS=darwin  GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-amd64 .
	GOOS=darwin  GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe .
	GOOS=windows GOARCH=arm64 go build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(APP_NAME)-windows-arm64.exe .
	@echo "Build complete. Binaries in $(BUILD_DIR)/"

clean:
	rm -rf $(BUILD_DIR) $(APP_NAME) $(APP_NAME).exe
