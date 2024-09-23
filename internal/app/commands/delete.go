package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (c *Commander) Delete(message *tgbotapi.Message) {
	args := message.CommandArguments()
	idx, _ := strconv.Atoi(args)
	_, err := c.productService.Delete(idx)
	if err != nil {
		c.bot.Send(tgbotapi.NewMessage(message.Chat.ID, err.Error()))
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Successfully deleted")
	c.bot.Send(msg)
}
