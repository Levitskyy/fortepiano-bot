package botbehaviour

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/botkeyboard"
	"github.com/Levitskyy/fortepiano-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EmailGetter interface {
	GetEmail(ctx context.Context, user model.User) (string, error)
}

type EmailSetter interface {
	UpdateEmail(ctx context.Context, user model.User, email string) error
}

// Callback
func CmdViewMyEmail(getter EmailGetter) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		email, err := getter.GetEmail(
			ctx,
			model.User{
				Id: update.FromChat().ChatConfig().ChatID,
			},
		)
		if err != nil {
			return err
		}
		reply := tgbotapi.NewMessage(update.CallbackQuery.From.ID, fmt.Sprintf("Ваша почта: %s\n\nЧтобы изменить почту напишите /email <Ваша почта>", email))
		reply.ReplyMarkup = botkeyboard.BackToMenuInlineKeyboard

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		if _, err := bot.Request(callback); err != nil {
			return err
		}

		return nil
	}
}

// Command
func CmdViewUpdateMyEmail(setter EmailSetter) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		email := update.Message.CommandArguments()
		if _, err := mail.ParseAddress(email); err != nil {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Некорректный адрес эл.почты")
			reply.ReplyMarkup = botkeyboard.BackToMenuInlineKeyboard
			if _, err := bot.Send(reply); err != nil {
				return err
			}

			return nil
		}
		err := setter.UpdateEmail(
			ctx,
			model.User{
				Id: update.FromChat().ChatConfig().ChatID,
			},
			email,
		)
		if err != nil {
			return err
		}

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Адрес эл.почты обновлён")
		reply.ReplyMarkup = botkeyboard.MenuInlineKeyboard
		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}
}
