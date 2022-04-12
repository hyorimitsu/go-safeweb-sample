FROM golang:1.17.8-alpine3.15

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go install github.com/cosmtrek/air@v1.27.8
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.0
