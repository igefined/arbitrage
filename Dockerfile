FROM golang:1.20-alpine AS builder

WORKDIR /app

ENV PORT=8080
ENV APP_NAME=arbitrage

COPY . .

RUN go mod tidy
RUN go build -o ./bin/app main.go

EXPOSE 8080

ENTRYPOINT ["/app/bin/app"]