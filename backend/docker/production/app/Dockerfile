# Docker build command


# build stage
FROM golang:1.20-alpine AS build

# Install necessary package
RUN set -eux; \
    apk update && apk add git 

# Set workdir
RUN mkdir -p /go/app
WORKDIR /go/app

# copy go.mod and go.sum
COPY go.mod ./
COPY go.sum ./

# Download go modules package
RUN go mod download

COPY . .

RUN go build -ldflags "-w -s" -o main

# Execute go binary
FROM alpine:latest AS execute

RUN apk update

COPY --from=build /go/app/main .

CMD ["./main"]
