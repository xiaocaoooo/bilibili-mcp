# Build stage
FROM golang:1.26-alpine AS builder

# Install git for fetching dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Run stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose port
EXPOSE 8080

# Default environment variables
ENV MCP_PORT=8080
ENV BILI_COOKIE=

# Run the server
CMD ["./server"]
