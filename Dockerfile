FROM gcr.io/cloud-builders/docker

LABEL maintainer "Brodie <brodiep21@hotmail.com>"

WORKDIR /weatherapp

COPY go.mod .

COPY notice.sh /usr/bin

COPY homepage.html .

COPY weather.html .

COPY main.go .

ENV PORT 8080

RUN go build

ENTRYPOINT go run main.go


# ENTRYPOINT ["/usr/bin/notice.sh"]
