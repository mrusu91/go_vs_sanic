FROM ubuntu:17.04

RUN apt-get update && apt-get upgrade -y && apt-get install -y \
  python3 \
  python3-pip \
  golang-go

COPY app /app
WORKDIR /app

RUN pip3 install sanic

ENV GOPATH=/app/go
ENV PATH=$PATH:$GOPATH/bin
RUN go get go_vs_sanic
