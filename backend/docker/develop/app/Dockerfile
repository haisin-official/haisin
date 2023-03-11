FROM golang:1.20-alpine

RUN set -eux; \
    apk update && apk add git

# install air for hot reload
RUN set -eux; \
    go install github.com/cosmtrek/air@latest

RUN mkdir -p /go/app

WORKDIR /go/app

# COPY go.mod ./
# COPY go.sum ./

COPY . .

CMD ["air", "main.go"]