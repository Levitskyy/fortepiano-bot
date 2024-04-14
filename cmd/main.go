package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/botbehaviour"
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

	bot := bot.New(botAPI)
	bot.RegisterCmdView("start", botbehaviour.CmdViewStart(userStorage))
	bot.RegisterCmdView("menu", botbehaviour.CmdViewOpenMenu())
	bot.RegisterCmdView("PCQ", botbehaviour.CmdViewAnswerPCQ())
	bot.RegisterCmdView("email", botbehaviour.CmdViewUpdateMyEmail(userStorage))
	bot.RegisterCallbackView("myEmailIK", botbehaviour.CmdViewMyEmail(userStorage))
	bot.RegisterCallbackView("menuIK", botbehaviour.CmdViewBackToMenu())
	bot.RegisterCallbackView("buySubsIK", botbehaviour.CmdViewSubOptions())

	for _, v := range []int{1, 3, 6, 12} {
		bot.RegisterCallbackView(fmt.Sprintf("FortePianoChannel_%d", v), botbehaviour.CmdViewBuySub("FortePianoChannel", v))
		bot.RegisterPaymentView(fmt.Sprintf("%s:%d:INV", "FortePianoChannel", v), botbehaviour.CmdViewSuccesfulPayment(groupStorage, subscriptionStorage, "FortePianoChannel", v))
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := bot.Run(ctx); err != nil {
		log.Printf("ERROR failed to run bot: %v", err)
	}
}
