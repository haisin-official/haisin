# Docker build command
# docker build -t haisin-official/backend:latest -t haisin-official/backend:v1 -f ./docker/production/app/Dockerfile .
# docker run --name haisin_backend --net=host --env-file "$(pwd)/.env.production" -p 8000:8000 haisin-official/backend:latest 
# build stage
FROM golang:1.20-alpine AS build
ENV TZ=Asia/Tokyo
ENV GOOS=linux
ENV GOARCH=amd64

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

RUN GO_ENV=production go build -ldflags "-w -s" -o main main.go

# Execute go binary
FROM alpine:latest AS execute

RUN apk update

COPY --from=build /go/app/main .

EXPOSE 8000

CMD ["./main"]
