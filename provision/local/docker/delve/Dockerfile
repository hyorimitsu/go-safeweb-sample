FROM golang:1.15.6-alpine3.12

ENV GO111MODULE=on
ENV AIR_VERSION=v1.12.4
ENV DELVE_VERSION=v1.5.0

RUN go get -u github.com/cosmtrek/air@"${AIR_VERSION}"
RUN go get -u github.com/go-delve/delve/cmd/dlv@"${DELVE_VERSION}"
