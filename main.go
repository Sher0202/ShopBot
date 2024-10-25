package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7968779965:AAGvHrad2a7C_RRBIvi3vFy6WHkfoA6Ydg8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			// Ботни /start буйруғи билан бошлаймиз
			if update.Message.Text == "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Салом! Ботга хуш келибсиз!")
				bot.Send(msg)
			}
			if update.Message.Text == "/aboutUs" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Биз сизга ӯз дӯконимиз тӯғрисида айтиб беришим мумкин!")
				bot.Send(msg)
			}
		}
	}
}
