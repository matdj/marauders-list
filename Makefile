build:
	go build .

dbuild: | build
	docker build -t marauders-list .

drun: 
	docker run -p 8080:8080 -it marauders-list /bin/sh

format:
	gofmt -s -w .

test:
	go test -v ./...