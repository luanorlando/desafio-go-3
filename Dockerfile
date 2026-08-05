# Etapa de build
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .

# Etapa de execução
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/configs ./configs

EXPOSE 8000 8080 50051

CMD ["./app"]