FROM golang:1.21-alpine

ENV GOOS linux
ENV CGO_ENABLED=0


WORKDIR /go-yandex

COPY go.mod .
COPY go.sum .

RUN go mod download && apk update && apk add make

COPY . .



