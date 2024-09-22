package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

func (c *Commander) List(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	products := ""
	for _, v := range c.productService.List() {
		products += v.Title + "\n"
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, strings.TrimSpace(products))

	c.bot.Send(msg)
}
