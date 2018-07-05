package main

import (
	"log"

	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/comandos"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/contato"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/github"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/livro"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/meetups"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/midias"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/projeto"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/start"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/telegram"
	"github.com/rafaelvicio/GDGBrasiliaBot/comandos/vagas"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("551297102:AAFPVW8XFdmwtlM0-ez4h3ykEXBbWnJ0DgM")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "meetups":
				msg.Text = meetups.GetMeetups()
			case "github":
				msg.Text = github.GetGitHub()
			case "contato":
				msg.Text = contato.GetContato()
			case "telegram":
				msg.Text = telegram.GetTelegram()
			case "vagas":
				msg.Text = vagas.GetVagas()
			case "comandos":
				msg.Text = comandos.GetComandos()
			case "midias":
				msg.Text = midias.GetMidias()
			case "projeto":
				msg.Text = projeto.GetProjeto()
			case "livros":
				msg.Text = livro.GetLivro()
			case "start":
				msg.Text = start.GetStart(update.Message.From.FirstName)
			default:
				msg.Text = "Desculpe, não conheço esse comando"
			}

			bot.Send(msg)
		}

	}
}
