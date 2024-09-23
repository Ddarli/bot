package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Help(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Commands:\n/help - help\n/list - products list\n"+
		"/read [index] - concrete product\n/create - create new product\n/delete [index] - delete product\n"+
		"/update [id] [new title] - update product\n")

	c.bot.Send(msg)
}
