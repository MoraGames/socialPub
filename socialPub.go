package socialPub

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//Execute this function
func CoreBot() {
	//Get the API_Token
	botToken, err := GetToken("@MG_Telegram_bot")
	if err != nil {
		log.Panic(err)
	}

	//Get the botAPI object
	botAPI, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	botAPI.Debug = true

	//Set the eventUpdate
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updatesChannel, err := botAPI.GetUpdatesChan(update)
	if err != nil {
		log.Panic(err)
	}

	//Run the bot
	for update := range updatesChannel {
		if update.Message == nil {
			//Ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		botAPI.Send(msg)
	}
}
