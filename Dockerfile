FROM gcr.io/cloud-builders/docker

LABEL maintainer "Brodie <brodiep21@hotmail.com>"

WORKDIR /weatherapp

COPY go.mod .

COPY homepage.html .

COPY weather.html .

COPY main.go .

ENV PORT 8080

RUN apk add --no-cache go

RUN go version

RUN go build

ENTRYPOINT go run main.go

