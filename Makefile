default: test build run

test:
	go test ./...

build:
	go build -o app ./cmd/api/main.go

run:
	./app ${ARGS}