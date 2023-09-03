package main

import (
	"github.com/Cc360428/telegram_robot/keyboard"
	"github.com/Cc360428/telegram_robot/pkg/mysql"
	"github.com/Cc360428/telegram_robot/ssh"
	"log"
	"os"

	"github.com/Cc360428/HelpPackage/other"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	mysql.NewClient()
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	cmd := bot.GetUpdatesChan(u)

	for update := range cmd {

		if update.CallbackQuery != nil {
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")

			if other.StrToInt(update.CallbackQuery.Data) <= 3 {
				msg.Text = "游戏列表如下"
			}
			switch update.CallbackQuery.Data {
			case "1":
				msg.ReplyMarkup = mysql.GetGameAllByPlatform(1)
			case "2":
				msg.ReplyMarkup = mysql.GetGameAllByPlatform(2)
			case "3":
				msg.ReplyMarkup = mysql.GetGameAllByPlatform(3)
			case "12":
				client := ssh.RestartServer()
				msg.Text = client

			case "2003":
				msg.Text = "Dummy; \n GameId:2003"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			continue
		}

		// 处理命令行
		if update.Message != nil && update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			// Extract the command from the Message.
			switch update.Message.Command() {
			case "open_keyboard":
				msg.ReplyMarkup = keyboard.NumericKeyboard
			case "close_keyboard":
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			case "status":
				msg.Text = "I'm ok."
			case "start", "help":
				msg.ReplyMarkup = keyboard.NumericKeyboard
				msg.Text = `Hello! `

			default:
				msg.Text = "Sorry 请给出对应的指令"
			}

			if _, err := bot.Send(msg); err != nil {
				log.Println(err)
			}
			continue
		}

		if update.Message != nil && update.Message.Text != "" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch msg.Text {
			case "GameInfo":
				msg.ReplyMarkup = mysql.GetKeyboardAllByType(mysql.UpdateVersion)
			case "ShareInfo":
				msg.ReplyMarkup = mysql.GetKeyboardAllByType(mysql.ShareDocument)
			case "重启服务":
				msg.ReplyMarkup = mysql.GetKeyboardAllByType(mysql.ShellCmd)
			case "other":
				log.Println("------------------------------other")
				msg.ReplyMarkup = mysql.GetKeyboardAllByType(mysql.UpdateVersion)
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}

	}
}
