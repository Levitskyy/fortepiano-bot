package botbehaviour

import (
	"context"
	"fmt"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/botkeyboard"
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
		reply.ReplyMarkup = botkeyboard.GetMenuKeyboard

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}

func CmdViewOpenMenu() bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Меню")
		reply.ReplyMarkup = botkeyboard.MenuInlineKeyboard

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}

func CmdViewBackToMenu() bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		reply := tgbotapi.NewEditMessageTextAndMarkup(
			update.CallbackQuery.From.ID,
			update.CallbackQuery.Message.MessageID,
			"Меню",
			botkeyboard.MenuInlineKeyboard,
		)

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}

func CmdViewSkipCallback() bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		if _, err := bot.Request(callback); err != nil {
			return err
		}

		return nil
	}
}
