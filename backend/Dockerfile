# https://code-yeongyu.tistory.com/36

FROM golang:1.20-alpine AS builder
LABEL authors="cheesecat47 <cheesecat47@gmail.com>"

RUN apk update

WORKDIR /build

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .
RUN go clean -cache && go build -o bin/main .

FROM alpine:3.18

RUN apk update && apk add --no-cache bash
COPY --from=builder /build/bin/main /main
