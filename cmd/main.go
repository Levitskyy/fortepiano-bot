package main

import (
	"log"

	"github.com/Levitskyy/fortepiano-bot/internal/config"
	"github.com/Levitskyy/fortepiano-bot/internal/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	botAPI, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err != nil {
		log.Printf("ERROR failed to create botAPI: %v", err)
	}

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	if err != nil {
		log.Printf("ERROR failed to connect to db: %v", err)
	}
	defer db.Close()

	var (
		userStorage         = storage.NewUserStorage(db)
		groupStorage        = storage.NewGroupStorage(db)
		subscriptionStorage = storage.NewSubscriptionStorage(db)
	)
}
