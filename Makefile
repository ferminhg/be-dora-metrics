.PHONY: build run test clean dev

TARGET := metrics
BUILD_DIR := build
MAIN_DIR := cmd/api

# build the binary
build:
	@go build -o $(BUILD_DIR)/$(TARGET) $(MAIN_DIR)/main.go

# run the binary
run: build
	@./$(BUILD_DIR)/$(TARGET)

# run tests
test:
	@go test -v ./...

# clean up
clean:
	@go clean
	@rm -f $(TARGET)

# run dev server
dev:
	@go run cmd/api/main.go

docker-run:
	@docker build -t $(TARGET) .
	@docker run --publish 8080:8080 --name webmetrics --rm metrics -m http.server --bind 0.0.0.0

docker-stop:
	@docker stop webmetrics test