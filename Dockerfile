# Based off https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/marauders-list
COPY . .

RUN go get -d -v

RUN go build .

# Container
FROM alpine
ARG run_mode="prod"
ENV RUN_MODE=$run_mode
ARG templates_dir="/go/templates"
ENV TEMPLATES_DIR=$templates_dir

WORKDIR /go

COPY --from=builder "/go/src/marauders-list/marauders-list" .
COPY --from=builder "/go/src/marauders-list/templates/*" "/go/templates/"

ENTRYPOINT ["./marauders-list"]