include go.mk


.PHONY: gen-version
gen-version: ## Update version
	cd pkg/version/scripts && go run gen.go

.PHONY: clean
clean:  ## Clean build bundles
	-rm -rf ./build

.PHONY: build-all
build-all: build-darwin build-linux build-windows ## Build all platforms

.PHONY: build-darwin
build-darwin: gen-version ## Build for MacOS
	-rm -rf ./build/darwin
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/darwin/$(APPROOT) \
		./main.go

.PHONY: build-linux
build-linux: gen-version ## Build for Linux
	-rm -rf ./build/linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/linux/$(APPROOT) \
		./main.go

.PHONY: build-windows
build-windows: gen-version ## Build for Windows
	-rm -rf ./build/windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/windows/$(APPROOT).exe \
		./main.go
