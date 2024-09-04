run: build
	@./bin/anki-duo

build:
	@go build -o bin/anki-duo main.go

