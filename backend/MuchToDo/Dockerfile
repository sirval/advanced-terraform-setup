# Build stage
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o muchtodo-api ./cmd/api

# Runtime stage
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates && adduser -D appuser

COPY --from=builder /app/muchtodo-api .

USER appuser

EXPOSE 8080

CMD ["./muchtodo-api"]