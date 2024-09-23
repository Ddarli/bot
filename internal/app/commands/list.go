package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) List(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	products := ""
	for i, v := range c.productService.List() {
		products += fmt.Sprintf("%v. %s\n", i+1, v.String())
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, products)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}
