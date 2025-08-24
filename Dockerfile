# Build stage
FROM golang:alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum for dependency caching
COPY go.mod .
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage
FROM alpine:latest
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
# Expose port if your application is a web server
EXPOSE 8080

# Command to run the application
CMD ["./main"]