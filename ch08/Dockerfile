# WITH Go Modules

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git

RUN mkdir $GOPATH/src/server
ADD ./prometheus.go $GOPATH/src/server

WORKDIR $GOPATH/src/server
RUN go mod init
RUN go mod tidy
RUN go mod download

RUN mkdir /pro
RUN go build -o /pro/server prometheus.go

FROM alpine:latest

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server

EXPOSE 1234
WORKDIR /pro
CMD ["/pro/server"]