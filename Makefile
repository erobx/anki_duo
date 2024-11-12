run: build
	@./bin/anki-duo

build:
	@go build -o bin/anki-duo cmd/anki-duo/main.go

test:
	@go run ./cmd/anki-duo/main.go

