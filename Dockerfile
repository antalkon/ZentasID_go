# Укажите базовый образ
FROM golang:latest AS builder

# Установите рабочую директорию внутри контейнера
WORKDIR /app

# Скопируйте go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загрузите зависимости
RUN go mod download

# Скопируйте остальные файлы в рабочую директорию
COPY . .

# Выполните go mod tidy для обновления зависимостей
RUN go mod tidy

# Соберите приложение
RUN go build -o main .

# Используйте минимальный образ для финального контейнера
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Запустите приложение
CMD ["./cmd/app/main"]
