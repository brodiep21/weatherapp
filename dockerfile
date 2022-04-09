FROM golang:1.17-alpine

LABEL maintainer "Brodie <brodiep21@hotmail.com>"

WORKDIR /weatherapp

COPY go.mod .

COPY homepage.html .

COPY weather.html .

COPY main.go .

ENV PORT 8080

RUN go build

ENTRYPOINT go run main.go