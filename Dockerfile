FROM golang:1.17-alpine

ENV GOPATH /usr/src/app/go
ARG dir=$GOPATH/src/github.com/mrauer
WORKDIR ${dir}

RUN apk add make ffmpeg gifsicle

COPY go.mod .
COPY go.sum .
RUN go mod download

WORKDIR $GOPATH/src/github.com/mrauer/gifme
COPY . .
