tidy:
	go mod tidy

build:
	go build ./cmd/web/main.go

run:
	@./main