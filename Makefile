PROJECT_NAME := "task-service"
APP_NAME := "service"
PKG := "./"
CMD := "$(PKG)/cmd/$(APP_NAME)"

build: ## Build the binary file
	@go build -o $(APP_NAME).o -i -v $(CMD)

run: build
	@./$(APP_NAME).o