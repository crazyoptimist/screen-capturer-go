APP_NAME=screen

.PHONY: build-linux build-windows build-mac all

build-linux:
	@echo "Building for Linux"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/linux/$(APP_NAME)-server-linux -tags netgo -ldflags "-w -s" ./cmd/server/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/linux/$(APP_NAME)-client-linux -tags netgo -ldflags "-w -s" ./cmd/client/
build-windows:
	@echo "Building for Windows"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/windows/$(APP_NAME)-server.exe -tags netgo -ldflags "-w -s" ./cmd/server/
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./dist/windows/$(APP_NAME)-client.exe -tags netgo -ldflags "-w -s" ./cmd/client/
build-mac:
	@echo "Building for Mac"
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/$(APP_NAME)-server-darwin -tags netgo -ldflags "-w -s" ./cmd/server/
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./dist/mac/$(APP_NAME)-client-darwin -tags netgo -ldflags "-w -s" ./cmd/client/
all: build-linux build-windows build-mac
