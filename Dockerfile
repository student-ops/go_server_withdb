# Builder stage
FROM golang:1.16 AS builder

WORKDIR /go/src/app
COPY . .

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -v -o app

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app .

CMD ["./app"]
