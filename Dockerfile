#!/bin/bash

FROM golang:latest as builder

WORKDIR /go/src/github.com/yatabis/Jehanne/TaskBoard
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

WORKDIR ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/yatabis/Jehanne/TaskBoard/cmd/main /main

ENV PORT 8080

CMD ["/main"]
