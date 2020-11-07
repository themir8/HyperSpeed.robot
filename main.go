package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Главное меню
var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Проверка кнопки"),
		tgbotapi.NewKeyboardButtonContact("Отправить телефон"),
	),
	// tgbotapi.NewKeyboardButtonRow(
	// ),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonLocation("Отправить местоположение"),
	),
)

var (
	bot          *tgbotapi.BotAPI
	err          error
	updChannel   tgbotapi.UpdatesChannel
	update       tgbotapi.Update
	updConfig    tgbotapi.UpdateConfig
	botUser      tgbotapi.User
	phone        string
	locationLink string
	photography  []tgbotapi.PhotoSize
)

func main() {

	bot, err = tgbotapi.NewBotAPI("1446049605:AAFU3Tzrp3SciOBQBS4dROnQFlsntxgy5to")
	if err != nil {
		panic("bot init error: " + err.Error())
	}

	botUser, err = bot.GetMe()
	if err != nil {
		panic("bot getme error: " + err.Error())
	}

	fmt.Printf("auth ok! bot is: %s\n", botUser.FirstName)

	updConfig.Timeout = 60 //
	updConfig.Limit = 1    // Неведомая хуйня, надо разобраться
	updConfig.Offset = 0   //

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		panic("update channel error: " + err.Error())
	}

	for {

		update = <-updChannel

		if update.Message != nil { // Обработка всех входящих сообщений
			if update.Message.IsCommand() { // Обработка команд
				cmdText := update.Message.Command()
				if cmdText == "start" { // Действие на команду start
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {
				if update.Message.Text == mainMenu.Keyboard[0][0].Text { // Обработка кнопки меню
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "/start")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false) // Спрятать клавиатуру
					bot.Send(msg)

				} else if update.Message.Text != "" { // Обработка текстовых сообщений
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					bot.Send(msg)

				} else if update.Message.Photo != nil { // Обработка входящих фотографий
					photography = *update.Message.Photo
					photoLastIndex := len(photography) - 1
					photo := photography[photoLastIndex] // Получаем последний элемент массива (самую большую фотку)
					photoMsg := tgbotapi.NewPhotoShare(update.Message.Chat.ID, photo.FileID)
					bot.Send(photoMsg)
					println("Photo id:", photo.FileID)

				} else if update.Message.Contact != nil { // Запрашиваем номер телефона
					phone = update.Message.Contact.PhoneNumber // Получаем номер телефона
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, phone)
					bot.Send(msg)

				} else if update.Message.Location != nil { // Запрашиваем местоположение
					lat := fmt.Sprint(update.Message.Location.Latitude)                                     // широта
					lon := fmt.Sprint(update.Message.Location.Longitude)                                    // долгота
					locationLink = fmt.Sprintf("https://www.google.com/maps?q=%v,%v&hl=ru&gl=ru", lat, lon) // Создаем ссылку для гугл карт и отправляем
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, locationLink)
					bot.Send(msg)

				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					bot.Send(msg)
				}
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "nil")
			bot.Send(msg)
		}
	}

}
