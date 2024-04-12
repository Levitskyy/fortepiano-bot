package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	botAPI, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Printf("ERROR failed to create botAPI: %v", err)
	}

	db, err := sqlx.Connect("postgres")
	if err != nil {
		log.Printf("ERROR failed to connect to db: %v", err)
	}
}
