FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY main.go .

RUN go build -o server main.go

# Final stage - minimal runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY static ./static

EXPOSE 8080

CMD ["./server"]
