FROM golang:1.19.1-alpine

RUN apk update && apk add git

ADD . /go/src/github.com/go-geo-server
WORKDIR /go/src/github.com/go-geo-server/src

CMD ["go", "run", "main.go"]