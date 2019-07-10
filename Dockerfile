# Based off https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang:alpine AS builder

ARG run_mode="prod"
ENV RUN_MODE=$run_mode

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/marauders-list
COPY . .

RUN go get -d -v

RUN go build
ENTRYPOINT ["/go/src/marauders-list/marauders-list"]
# FROM scratch

# COPY --from=builder /go/src/marauders-list/marauders-list /go/bin/marauders-list

# ENTRYPOINT ["/go/bin/marauders-list"]