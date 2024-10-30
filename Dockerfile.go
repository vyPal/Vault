# Dockerfile.go
# Stage 1: Build the Go application
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code and build it
COPY . .
RUN go build -o server ./path/to/your/go/server # Adjust to the server directory

# Stage 2: Run the Go application
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .

# Copy .env file for environment variables
COPY .env .env

# Run the server
CMD ["sh", "-c", "source .env && ./server"]

