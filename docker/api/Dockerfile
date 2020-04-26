FROM golang:1.14.2

ENV GO111MODULE=on
ENV APP_DIR=/go/src/github.com/konchanxxx/sample-go-mysql-docker

WORKDIR $APP_DIR

# APIサーバーからDBを参照する際のデバッグ等に使えるのでclientを入れておく
RUN apt-get update && apt-get install -y default-mysql-client

COPY . .

RUN go mod download
