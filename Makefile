build:
	@go build -o ./bin/ims ./cmd/ims/main.go

run: build
	@./bin/ims

test:
	@go test -v ./...
