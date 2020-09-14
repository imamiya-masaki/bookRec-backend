FROM golang:1.14.2-alpine

ENV GO111MODULE=on

RUN apk update && apk add git bash

RUN apk add build-base

RUN go get github.com/pilu/fresh && \
    go get -u github.com/gin-gonic/gin && \
    go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/mysql

# CMD [ "go", "run", "main.go" ]