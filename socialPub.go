package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	log.Print(">>> Executing bot code:\n\n")
	CoreBot()
}

func CoreBot() {
	//Get the API_Token
	botToken, err := GetToken("@MG_Telegram_bot")
	if err != nil {
		log.Panic(err)
	}
	log.Print(">> Token:", botToken, "\n\n")

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

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
			case "help":
			}
		}

		response := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		response.ReplyToMessageID = update.Message.MessageID
		if !strings.Contains(update.Message.Text, "?") {
			response.Text = "Non stai inviando una domanda! Rispondi ad una domanda ricevuta o formula una domanda?"
			botAPI.Send(response)
			continue
		}
		response.Text = "Domanda inviata!"

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		botAPI.Send(response)
	}
}
