# Базовый образ
FROM golang:1.24.2-alpine

# Создаём рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем бинарник
RUN go build -o ascue .

# Указываем порт
EXPOSE 8080

# Команда запуска
CMD ["./ascue"]