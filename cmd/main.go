package main

import (
	"log"
	"user-management-server/database"
)

func main() {
	if err := database.ConnectDatabase(); err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	log.Println("Приложение успешно подключено и запущено.")
}
