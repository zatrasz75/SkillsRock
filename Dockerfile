FROM golang:1.23.2-alpine AS builder
LABEL authors="https://t.me/Zatrasz"

RUN apk add --no-cache git

# Создание рабочий директории
RUN mkdir -p /app

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы проекта внутрь контейнера
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/app

# Второй этап: создание production образ
FROM ubuntu AS chemistry

WORKDIR /app

# Обновляем список пакетов
RUN apt-get update && apt-get install -y nginx

COPY --from=builder /app/bin/app /app/app

COPY --from=builder /app/docs /app/docs

COPY --from=builder /app/configs /app/configs

# Устанавливаем сервер для обслуживания Swagger
COPY ./nginx.conf /etc/nginx/sites-available/default

# Копируем остальное
COPY ./ ./

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]

CMD ["./app"]