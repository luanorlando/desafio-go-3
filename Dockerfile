# Etapa de build
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd/ordersystem

# Etapa de execução
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl libc6-compat mysql-client

WORKDIR /app

COPY --from=builder /app/app .
COPY .env.example .env
# COPY --from=builder /app/configs ./configs

EXPOSE 8000 8080 50051

CMD ["./app"]