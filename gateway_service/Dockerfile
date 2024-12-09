FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o gateway_service ./cmd/gateway

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/gateway_service .
EXPOSE 8080
CMD ["./gateway_service"]
