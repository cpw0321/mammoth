FROM golang:1.17.11 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 1
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/

COPY . .

RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o /app/app /build/cmd/main.go

FROM ubuntu:18.04

RUN apt-get update && apt-get install -y locales ca-certificates tzdata && mkdir -p /app/etc/
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY etc/config.toml /app/etc/config.toml
EXPOSE 8080
CMD ["./app", "-f", "/app/etc/config.toml"]