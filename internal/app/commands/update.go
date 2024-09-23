package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

func (c *Commander) Update(message *tgbotapi.Message) {
	params := strings.Split(message.Text, " ")
	idx, _ := strconv.Atoi(params[1])
	product, err := c.productService.Update(idx, params[2])
	if err != nil {
		fmt.Errorf("can't update product: %v", err)
		return
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("%d. %s", idx, product.String()))

	c.bot.Send(msg)
}
