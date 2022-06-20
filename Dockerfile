FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o consumer-app ./cmd/consumer/main.go
RUN go build -o producer-app ./cmd/producer/main.go

FROM alpine:latest
WORKDIR /usr/bin
COPY .env .
COPY --from=builder /app/consumer-app ./
COPY --from=builder /app/producer-app ./