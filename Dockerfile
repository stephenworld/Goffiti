# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application as a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o goffiti .

# Stage 2: Final lightweight image
FROM alpine:latest

WORKDIR /root/

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/goffiti .

# Since it's a CLI tool, we use ENTRYPOINT so users can pass flags
ENTRYPOINT ["./goffiti"]