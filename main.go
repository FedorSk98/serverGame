package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // Подключенные клиенты
var broadcast = make(chan []byte)            // Канал сообщений
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func main() {
	// Настройка маршрута WebSocket
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/hello", hello)

	// Запуск обработки сообщений
	go handleMessages()

	// Запуск сервера на порту 8080
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "<h1>hello</h1>")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Апгрейд HTTP-запроса в WebSocket

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	// Закрытие подключения при выходе из функции
	defer ws.Close()

	// Добавление клиента в список подключенных
	clients[ws] = true

	// Бесконечный цикл чтения сообщений с WebSocket
	for {
		var msg []byte
		// Чтение сообщения
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v", err)
			delete(clients, ws)
			break
		}

		log.Printf("чтение сообщения: %v", string(msg))
		// Отправка сообщения в канал
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Получение сообщения из канала
		msg := <-broadcast

		// Отправка сообщения всем подключенным клиентам
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
