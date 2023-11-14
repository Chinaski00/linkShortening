# Используем официальный образ Golang
FROM golang:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы приложения в контейнер
COPY . .

# Собираем приложение
RUN go build -o shorturlservice .

# Экспортируем порт, который будет использоваться приложением
EXPOSE 8080

# Запускаем приложение
CMD ["./shorturlservice", "in_memory"]