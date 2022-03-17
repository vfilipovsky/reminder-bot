package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const defaultText = "remind you about "

type reminder struct {
	ID        uint
	UserId    int64
	ChatId    int64
	MessageId int
	Message   string
	NotifyAt  int64
}

type parsedMessage struct {
	text string
	when int64
}

func remind(bot *tgbotapi.BotAPI, store storage) {
	for range time.Tick(5 * time.Second) {
		reminders, err := store.find()

		if err != nil {
			logrus.Error(err)
			continue
		}

		for _, reminder := range reminders {
			r := reminder
			go func() {
				_, err := bot.Send(newMessage(r.ChatId, defaultText+r.Message, r.MessageId))

				if err != nil {
					logrus.Error(err)
					return
				}

				err = store.delete(r.ID)

				if err != nil {
					logrus.Error(err)
					return
				}
			}()
		}
	}
}

func parse(text string) *parsedMessage {
	return &parsedMessage{text: "not implemented yet", when: 1647533120}
}

func createReminder(message *tgbotapi.Message, store storage) *tgbotapi.MessageConfig {
	response := "reminder created"
	parsedMessage := parse(message.Text)

	reminder := &reminder{
		UserId:    message.From.ID,
		ChatId:    message.Chat.ID,
		MessageId: message.MessageID,
		Message:   parsedMessage.text,
		NotifyAt:  parsedMessage.when,
	}

	err := store.save(reminder)

	if err != nil {
		logrus.Error(err)
		response = "unexpected error occurred"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ReplyToMessageID = message.MessageID

	return newMessage(message.Chat.ID, response, message.MessageID)
}
