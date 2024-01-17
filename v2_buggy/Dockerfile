FROM golang:1.13-alpine

LABEL author = buksi.pajti@gmail.com

WORKDIR /opt

COPY buggy.go /opt

RUN ls -lsa && cat buggy.go

ENV CGO_ENABLED = 0 go build -o /usr/myapp && ls -lsa