# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /go/src/app

# Копируем файлы Go приложения внутрь образа
COPY . .

# Собираем Go приложение
RUN go build -o main .

# Указываем команду для запуска приложения при старте контейнера
CMD ["./main"]