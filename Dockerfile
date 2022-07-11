# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine3.16

RUN apk update \
  && apk add --no-cache \
    build-base=0.5-r3 \
    git=2.36.1-r0 \
    protoc=3.18.1-r2 \
    bash=5.1.16-r2 \
    sqlite=3.38.5-r0

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
  && go install github.com/cosmtrek/air@v1.40.3

WORKDIR /go/src/app

ENTRYPOINT [ "bash" ]
