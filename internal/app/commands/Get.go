package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Get(message *tgbotapi.Message) {
	args := message.CommandArguments()
	idx, err := strconv.Atoi(args)
	if err != nil {
		panic("Illegal index")
	}
	prod, err := c.productService.GetById(idx)
	if err != nil {
		panic("Product not found")
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, prod.Title)

	c.bot.Send(msg)
}
