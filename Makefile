tidy:
	go mod tidy

build:
	go build ./cmd/web/main.go

proto:
	protoc -I . --go_out=plugins=grpc:./proto count_num.proto

run:
	@./main