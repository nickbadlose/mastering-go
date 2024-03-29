# docker build -t goPost .
# docker run -it goPost

FROM golang:1.22.1-alpine3.19 AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /pro
ADD . /pro/
WORKDIR /pro
RUN go mod download
RUN go build -o server ./cmd/usePostgres.go

FROM alpine:latest

RUN mkdir /pro
COPY --from=builder /pro/server /pro/server
WORKDIR /pro
CMD ["/pro/server"]