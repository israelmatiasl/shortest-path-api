# ---------------------------
# Etapa 1: Build 
# ---------------------------
FROM golang:1.21-alpine AS builder
WORKDIR /src

# Install git (for go modules) and ca-certificates
RUN apk add --no-cache git ca-certificates

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the binary (adjust -o and package path if needed)
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /app ./cmd/api

# ---------------------------
# Etapa 2: Imagen final
# ---------------------------
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app /app

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app"]