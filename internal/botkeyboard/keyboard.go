package botkeyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var StartKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Старт", "testK"),
	),
)

var GetMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/menu"),
	),
)

var MenuInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Моя эл.почта", "myEmailIK"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Купить подписку", "buySubsIK"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Мои подписки", "getMySubsIK"),
	),
)

var SubOptionsInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Курсы FortePiano (мес)", "fortePianoCourseIK"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "FortePianoChannel_1"),
		tgbotapi.NewInlineKeyboardButtonData("3", "FortePianoChannel_3"),
		tgbotapi.NewInlineKeyboardButtonData("6", "FortePianoChannel_6"),
		tgbotapi.NewInlineKeyboardButtonData("12", "FortePianoChannel_12"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<- Назад в меню", "menuIK"),
	),
)

var BackToMenuInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("<- Вернуться в меню", "menuIK"),
	),
)
