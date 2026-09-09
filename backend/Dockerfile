# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init app || true
RUN go build -o server main.go

# Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]