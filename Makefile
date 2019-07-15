build:
	go build -o ./marauders-list cmd/marauders.go

clean:
	go clean

dbuild:
	docker build -t marauders-list .

drun: 
	docker run -p 8080:8080 -it marauders-list

format:
	gofmt -s -w .

run:
	go run cmd/marauders.go

test:
	go test -v ./...