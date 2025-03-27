# Первый Dockerfile
FROM golang:1.23.6 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальной код проекта
COPY . .

# Собираем бинарный файл
RUN go build -o main main.go

# Используем минимальный образ для запуска
FROM ubuntu:latest

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /root/

# Создаем директорию для хранения конфигов
RUN mkdir "configs"

# Копируем бинарный файл из стадии сборки
COPY --from=builder /app/main .

# Копируем конфиги
COPY --from=builder /app/configs/docker/configs.json ./configs
COPY --from=builder /app/configs/docker/example.json ./configs

# Копируем переменные окружения
COPY --from=builder /app/.env .
COPY --from=builder /app/example.env .

# Открываем порт
EXPOSE 8585

# Команда для запуска
CMD ["./main"]
