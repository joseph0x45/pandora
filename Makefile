build:
	@go build -o pandora main.go

download-deps:
	@go mod download
