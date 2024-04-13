package botbehaviour

import (
	"context"
	"fmt"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserAdder interface {
	Add(ctx context.Context, user model.User) error
}

func CmdViewStart(adder UserAdder) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {

		err := adder.Add(
			ctx,
			model.User{
				Id:   update.Message.From.ID,
				Name: fmt.Sprintf("%s %s", update.Message.From.LastName, update.Message.From.FirstName),
			},
		)
		if err != nil {
			return err
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello, %s!", update.Message.From.UserName))

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}
