FROM golang:1.13-alpine

LABEL author = buksi.pajti@gmail.com

WORKDIR /opt

COPY fibonacci.go /opt

RUN ls -lsa && cat fibonacci.go

ENV CGO_ENABLED = 0 go build -o /usr/myapp && ls -lsa