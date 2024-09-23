package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Create(message *tgbotapi.Message) {
	newProduct := c.productService.Create()
	msg := tgbotapi.NewMessage(message.Chat.ID, "Created new product - "+newProduct.String())

	c.bot.Send(msg)
}
