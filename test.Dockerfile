#!/bin/bash

FROM golang:latest

WORKDIR /go/src/github.com/yatabis/Jehanne/TaskBoard
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .

WORKDIR ./cmd
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

WORKDIR ../
RUN CGO_ENABLED=0 GOOS=linux go test -v ./*/
