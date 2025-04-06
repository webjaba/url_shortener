BINARY_NAME = bin/url_shortener
MAIN_FILE = cmd/main.go
BUILD_FLAGS = -v
STRG = postgres # postgres, inmemory

build:
	@echo "Building..."
    STORAGE=$(STRG) go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_FILE)

run:
	@echo "Running..."
	STORAGE=$(STRG) go run $(MAIN_FILE)