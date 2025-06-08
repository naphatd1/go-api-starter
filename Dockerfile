# ---------- Stage: Builder ----------
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# ติดตั้ง git สำหรับ go get ที่ใช้ private repo
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# build สำหรับ production
RUN go build -o server ./cmd/api

# ---------- Stage: Dev (Air Hot-Reload) ----------
FROM golang:1.24.3-alpine AS dev

WORKDIR /app

# ติดตั้ง Air
RUN apk add --no-cache curl git \
 && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

COPY . .

CMD ["air", "--config", "./tmp/.air.toml"]

# ---------- Stage: Prod ----------
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server ./server
COPY .env .env

EXPOSE 4000

ENTRYPOINT ["./server"]
