# Stage 1: Build
FROM golang:1.25 AS builder

# Install git and make (if needed)
RUN apt-get update && apt-get install -y git make

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 creates a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

# Stage 2: Run
FROM alpine:latest

# Install ca-certificates for HTTPS calls (e.g. Google OAuth)
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

# Copy migrations if needed (optional, if you run migrations inside app)
COPY --from=builder /app/migrations ./migrations

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
