FROM golang:alpine as build
RUN apk add build-base
RUN apk add git
WORKDIR /src
COPY . .
RUN g