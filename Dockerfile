FROM golang:latest

RUN mkdir -p /go/src/github.com/msyamsula/grpc_mongo

WORKDIR /go/src/github.com/msyamsula/grpc_mongo

COPY ./go.mod go.mod
COPY ./connection connection
COPY ./model model
COPY ./proto proto
COPY ./server server
COPY ./service service
COPY ./Makefile Makefile

RUN go mod tidy