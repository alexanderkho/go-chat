# Build stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy dependencies first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build server
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server

# Minimal final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
# ENTRYPOINT ["tail", "-f", "/dev/null"]
CMD ["./server"]
