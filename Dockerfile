# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the Go module file
COPY go.mod ./

# Copy the rest of the application source code
COPY . .

# Build the binary with CGO disabled for an Alpine environment
RUN CGO_ENABLED=0 GOOS=linux go build -o goffiti .

# Final minimal stage
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/goffiti .

# Copy the runtime assets
COPY --from=builder /app/web ./web
COPY --from=builder /app/server ./server

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./goffiti"]
