# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /go/src/app

# Копируем файлы Go приложения внутрь образа
COPY . .

EXPOSE 8080
# Собираем Go приложение
RUN go get github.com/gorilla/websocket
RUN go get golang.org/x/net
RUN go build -o main .

# Указываем команду для запуска приложения при старте контейнера
CMD ["./main"]