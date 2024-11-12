simple:
	@go run ./cmd/anki-duo/main.go

run: build
	@./bin/anki-duo

build:
	@go build -o bin/anki-duo main.go

