package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Callback(callback *tgbotapi.CallbackQuery) {
	parseData := CommandData{}
	json.Unmarshal([]byte(callback.Data), &parseData)
	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		fmt.Sprintf("Parsed: %v+\n", parseData),
	)
	c.bot.Send(msg)
}
