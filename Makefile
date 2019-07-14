build:
	go build .

format:
	gofmt -s -w .

test:
	go test -v ./...