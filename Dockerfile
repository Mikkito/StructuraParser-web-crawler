# Stage 1: build
FROM golang:1.21 AS builder

WORKDIR /app

# Кэширование зависимостей
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o web-crawler ./cmd/app.go

# Stage 2: final
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/web-crawler .

EXPOSE 8080

CMD ["./web-crawler"]