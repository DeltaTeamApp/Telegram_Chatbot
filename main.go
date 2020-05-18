package main

import (
	"DeltaTeleBot/config"
	"DeltaTeleBot/controller"
	"DeltaTeleBot/handler/ggsheet"
	"log"

	tb "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	config.Init()
	ggsheet.ServiceGGSheetInit()
	teleConfigObj := config.GetTeleConfigObj()

	bot, err := tb.NewBotAPI(teleConfigObj.APIKey)
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tb.NewUpdate(teleConfigObj.UpdateOffset)
	u.Timeout = teleConfigObj.TimeOut

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message.IsCommand() {
			msg := tb.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "/createfwshortlink [range]"

			case "createfwshortlink":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "createfwshortlink : no argument found"
				} else {
					msgChan := controller.GetUpdateChan()
					go controller.CreateFwdAndShortLinks(arg)
					for newMsg := range msgChan {
						if newMsg == "exit" {
							break
						} else {
							msg1 := tb.NewMessage(update.Message.Chat.ID, newMsg)
							msg1.DisableWebPagePreview = true
							if _, err := bot.Send(msg1); err != nil {
								log.Panic(err)
							}
						}
					}
					msg.Text = "complete"
				}

			case "setrebrandapi":
				arg := update.Message.CommandArguments()
				if len(arg) < 1 {
					msg.Text = "createfwshortlink : no argument found"
				} else {
					msg.Text = controller.SetNewRebrandAPIKey(arg)
				}

			default:
				msg.Text = "Unknown command"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
