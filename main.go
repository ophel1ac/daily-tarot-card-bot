package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ophel1ac/daily-tarot-card-bot/model"
	"github.com/ophel1ac/daily-tarot-card-bot/utils"
	"log"
	"math/rand"
	"time"
)

var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🔮", "button_get_random_card"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI(utils.GetBotKey())
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		// Check if we've gotten a message update.
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "[WIP] This bot sends you random tarot card.")
			msg.ReplyMarkup = inlineKeyboard
			// Send the message.
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := bot.Request(callback); err != nil {
				panic(err)
			}

			// Send the message.
			// And finally, send a message containing the data received.
			switch update.CallbackQuery.Data {
			case "button_get_random_card":
				getRandomCardHandler(bot, &update)
			}
		}
	}
}

func getRandomCardHandler(b *tgbotapi.BotAPI, update *tgbotapi.Update) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(22)

	card := model.GetCard("default", v)

	img := card.GetImage()

	msg := tgbotapi.NewPhoto(update.CallbackQuery.Message.Chat.ID, tgbotapi.FileBytes{
		Name:  "default.jpg",
		Bytes: img,
	})
	msg.Caption = card.Name + "\n" + card.MeaningUp

	if _, err := b.Send(msg); err != nil {
		panic(err)
	}
}
