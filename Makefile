.PHONY: clean
clean:
	go clean

.PHONY: build
build:
	go build -o build/ .\cmd\five9-scim\five9_scim.go