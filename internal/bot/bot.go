package bot

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ViewFunc func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error

type Bot struct {
	api           *tgbotapi.BotAPI
	cmdViews      map[string]ViewFunc
	callbackViews map[string]ViewFunc
	paymentViews  map[string]ViewFunc
}

func New(api *tgbotapi.BotAPI) *Bot {
	return &Bot{api: api}
}

func (b *Bot) RegisterCmdView(cmd string, view ViewFunc) {
	if b.cmdViews == nil {
		b.cmdViews = make(map[string]ViewFunc)
	}

	b.cmdViews[cmd] = view
}

func (b *Bot) RegisterCallbackView(cmd string, view ViewFunc) {
	if b.callbackViews == nil {
		b.callbackViews = make(map[string]ViewFunc)
	}

	b.callbackViews[cmd] = view
}

func (b *Bot) RegisterPaymentView(cmd string, view ViewFunc) {
	if b.paymentViews == nil {
		b.paymentViews = make(map[string]ViewFunc)
	}

	b.paymentViews[cmd] = view
}

func (b *Bot) Run(ctx context.Context) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.api.GetUpdatesChan(u)

	for {
		select {
		case update := <-updates:
			go func() {
				updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Minute)
				defer updateCancel()
				b.handleUpdate(updateCtx, update)
			}()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (b *Bot) handleUpdate(ctx context.Context, update tgbotapi.Update) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("ERROR panic recovered: %v\n%s", p, string(debug.Stack()))
		}
	}()

	var cmd string
	cmdView := func() ViewFunc {
		if update.CallbackQuery != nil {
			cmd = update.CallbackQuery.Data
			cmdView, ok := b.callbackViews[cmd]
			if !ok {
				return nil
			}
			return cmdView
		} else if update.PreCheckoutQuery != nil {
			cmd = "PCQ"
			cmdView, ok := b.cmdViews[cmd]
			if !ok {
				return nil
			}
			return cmdView
		} else if update.ChatJoinRequest != nil {
			cmd = "ChatJoinRequest"
			cmdView, ok := b.callbackViews[cmd]
			if !ok {
				return nil
			}
			return cmdView
		} else if update.Message != nil {
			if update.Message.SuccessfulPayment != nil {
				cmd = update.Message.SuccessfulPayment.InvoicePayload
				cmdView, ok := b.paymentViews[cmd]
				if !ok {
					return nil
				}
				return cmdView
			} else if update.Message.IsCommand() {
				cmd = update.Message.Command()
				cmdView, ok := b.cmdViews[cmd]
				if !ok {
					return nil
				}
				return cmdView
			}
			return nil
		} else {
			return nil
		}
	}()

	if cmdView == nil {
		return
	}
	if err := cmdView(ctx, b.api, update); err != nil {
		log.Printf("ERROR failed to execute view: %v", err)

		if _, err := b.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Внутреняя ошибка")); err != nil {
			log.Printf("ERROR failed to send error message: %v", err)
		}
	}
}
