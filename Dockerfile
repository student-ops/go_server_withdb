# Builder stage
FROM golang:1.16 AS builder

WORKDIR /app
COPY app/ .

RUN go mod init app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .

CMD ["./app"]
