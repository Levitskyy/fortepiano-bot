package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN" env-required:"true"`
	DatabaseDSN      string `env:"DATABASE_DSN" env-required:"true"`
	PaymentToken     string `env:"PAYMENT_TOKEN" env-required:"true"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Printf("ERROR config load error: %v", err)
		}
	})

	return cfg
}
