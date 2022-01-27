FROM golang:1.12 as builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

FROM scratch as runtime
COPY --from=builder /src/app .
EXPOSE 8080
USER 1000
ENTRYPOINT ["/app"]
