FROM golang:1.20-alpine AS build
WORKDIR /app
COPY main.go .
RUN apk add --no-cache git && \
    go mod init prodgo || true && \
    go get github.com/streadway/amqp && \
    go build -o /prod ./main.go

FROM alpine:3.18
COPY --from=build /prod /prod
CMD ["/prod"]
