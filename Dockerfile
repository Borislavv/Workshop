FROM golang:1.13-alpine

WORKDIR /srv/www/Go/Workshop
ADD ./ /srv/www/Go/Workshop

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/workshop cmd/workshop/main.go
ENTRYPOINT ["bin/workshop"]