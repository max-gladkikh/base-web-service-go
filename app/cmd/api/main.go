package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" // Профилировщик для отслеживания производительности
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("Hello from Go!"))
	})

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       5 * time.Second,   // Максимальное время на чтение запроса
		WriteTimeout:      10 * time.Second,  // Максимальное время на запись ответа
		IdleTimeout:       120 * time.Second, // Время жизни простаивающего соединения (Keep-Alive)
		ReadHeaderTimeout: 2 * time.Second,   // Время на чтение заголовков (защита от медленных атак)
	}

	fmt.Println("Сервер запущен на :8080...")
	// Запуск HTTP-сервера
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Ошибка сервера: %v\n", err)
	}
}
