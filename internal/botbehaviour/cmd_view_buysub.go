package botbehaviour

import (
	"context"
	"fmt"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/botkeyboard"
	"github.com/Levitskyy/fortepiano-bot/internal/config"
	"github.com/Levitskyy/fortepiano-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubUpdater interface {
	UpdateEndDate(ctx context.Context, subscription model.Subscription, days int) error
}

type IdGetter interface {
	GetId(ctx context.Context, group model.Group) (int64, error)
}

// Callback
func CmdViewSubOptions() bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		reply := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "Варианты подписок:")
		reply.ReplyMarkup = botkeyboard.SubOptionsInlineKeyboard

		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}

}

func CmdViewBuySub(groupName string, months int) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		price := tgbotapi.LabeledPrice{Label: "Руб", Amount: 100000 * months}
		prices := []tgbotapi.LabeledPrice{price}
		invoice := tgbotapi.NewInvoice(
			update.CallbackQuery.From.ID,
			fmt.Sprintf("Подписка на %s", groupName),
			fmt.Sprintf("Подписка на %s на %d мес.", groupName, months),
			fmt.Sprintf("%s:%d:INV", groupName, months),
			config.Get().PaymentToken,
			"",
			"RUB",
			prices,
		)
		invoice.SuggestedTipAmounts = []int{}

		if _, err := bot.Send(invoice); err != nil {
			return err
		}

		return nil
	}
}

// PreCheckoutQuery
func CmdViewAnswerPCQ() bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		reply := tgbotapi.PreCheckoutConfig{
			PreCheckoutQueryID: update.PreCheckoutQuery.ID,
			OK:                 true,
		}
		if _, err := bot.Request(reply); err != nil {
			return err
		}

		return nil
	}
}

func CmdViewSuccesfulPayment(getter IdGetter, updater SubUpdater, group string, months int) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		id, err := getter.GetId(ctx, model.Group{Name: group})
		if err != nil {
			reply := tgbotapi.NewMessage(update.Message.From.ID, "Ошибка получения, обратитесь к администратору")
			if _, err := bot.Send(reply); err != nil {
				return err
			}
			return err
		}

		err = updater.UpdateEndDate(ctx, model.Subscription{UserId: update.Message.From.ID, GroupId: id}, 30*months)
		if err != nil {
			reply := tgbotapi.NewMessage(update.Message.From.ID, "Ошибка получения, обратитесь к администратору")
			if _, err := bot.Send(reply); err != nil {
				return err
			}
			return err
		}

		reply := tgbotapi.NewMessage(update.Message.From.ID, fmt.Sprintf("Вы продлили подписку на %d мес.", months))
		reply.ReplyMarkup = botkeyboard.MenuInlineKeyboard
		if _, err := bot.Send(reply); err != nil {
			return err
		}

		return nil
	}

}
