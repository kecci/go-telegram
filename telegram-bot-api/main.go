package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kecci/go-telegram/constant"
)

const (
	CHAT_ID = int64(-744073805)

	COMMAND_REGISTER = "/register"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(constant.BOT_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// WAITING MESSAGE
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case COMMAND_REGISTER:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Successfully registered")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "command not found, use /register")
				bot.Send(msg)
			}
		}
	}
}
