# https://code-yeongyu.tistory.com/36

FROM golang:1.21-alpine AS builder
LABEL authors="cheesecat47 <cheesecat47@gmail.com>"

RUN apk update

WORKDIR /build

ENV GO111MODULE=auto \
    CGO_ENABLE=0 \
    GOOS=linux

COPY go.mod .
COPY go.sum .
COPY vendor/ .

COPY . .
RUN go build -v -o bin/main ./...

FROM alpine:3.18

RUN apk update && apk add --no-cache bash
COPY --from=builder /build/bin/main /main
