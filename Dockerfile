FROM golang:1.17-alpine

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}

RUN apk add make

COPY go.mod .

WORKDIR $GOPATH/src/github.com/mrauer/gifme
COPY . .
