FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest
# Install tzdata and ca-certificates เพื่อใช้ time zone asia/bangkok
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime && \
    echo "Asia/Bangkok" > /etc/timezone
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env.example /root/.env
EXPOSE 8080
CMD ["./main"]
