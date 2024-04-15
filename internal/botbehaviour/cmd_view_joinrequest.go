package botbehaviour

import (
	"context"

	"github.com/Levitskyy/fortepiano-bot/internal/bot"
	"github.com/Levitskyy/fortepiano-bot/internal/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SubChecker interface {
	IsSubActive(ctx context.Context, subscription model.Subscription) (bool, error)
}

func CmdViewChatJoinRequest(checker SubChecker) bot.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		if ok, _ := checker.IsSubActive(
			ctx,
			model.Subscription{UserId: update.ChatJoinRequest.From.ID, GroupId: update.ChatJoinRequest.Chat.ID},
		); ok {
			request := tgbotapi.ApproveChatJoinRequestConfig{
				ChatConfig: tgbotapi.ChatConfig{ChatID: update.ChatJoinRequest.Chat.ID, SuperGroupUsername: ""},
				UserID:     update.ChatJoinRequest.From.ID,
			}
			if _, err := bot.Request(request); err != nil {
				return err
			}

			return nil
		}

		request := tgbotapi.DeclineChatJoinRequest{
			ChatConfig: tgbotapi.ChatConfig{ChatID: update.ChatJoinRequest.Chat.ID, SuperGroupUsername: ""},
			UserID:     update.ChatJoinRequest.From.ID,
		}
		if _, err := bot.Request(request); err != nil {
			return err
		}

		return nil
	}
}
