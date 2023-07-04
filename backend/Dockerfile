# https://code-yeongyu.tistory.com/36

FROM golang:1.20-alpine AS builder
LABEL authors="cheesecat47 <cheesecat47@gmail.com>"

RUN apk update

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod tidy

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux

COPY . .

RUN go build -a -ldflags="-s -w" -o bin/main .

FROM alpine:3.18

COPY --from=builder /build/bin/main /main

ENTRYPOINT ["/main"]
