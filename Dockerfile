# Step 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV GO111MODULE=on

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go project using the Makefile
RUN go build -o /app/bin/zabbixprometheusexporter ./cmd/zabbixprometheusexporter

# Step 2: Create a lightweight image to run the binary
FROM alpine:latest

# Install certificates for HTTPS communication (if needed)
RUN apk add --no-cache ca-certificates

# Set working directory in the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bin/zabbixprometheusexporter .

# Copy the .env file
COPY .env .

# Expose the metrics port
EXPOSE 9100

# Run the Go binary
CMD ["./zabbixprometheusexporter"]
