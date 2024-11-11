# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies first (if any)
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o kube-connectivity-probe

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/kube-connectivity-probe .

# Expose the health check port
EXPOSE 8080

# Run the application
CMD ["./kube-connectivity-probe"]
