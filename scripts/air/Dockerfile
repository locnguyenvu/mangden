FROM golang:1.20-alpine 

ENV GOPATH /go
ENV GO111MODULE on
ENV OOS linux
ENV CGO_ENABLED 0


RUN apk add gcc \
    && go install github.com/cosmtrek/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

ENTRYPOINT ["/go/bin/air"]
