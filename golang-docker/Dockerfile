## We specify the base image we need for our
## go application
FROM golang:1.16-alpine as builder


ENV GO111MODULE=off
WORKDIR /app
COPY . /app

RUN go build -o main ./server/

EXPOSE 8081
CMD ["/app/main"]
