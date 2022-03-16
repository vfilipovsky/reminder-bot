package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const defaultText = "remind you about "

type reminder struct {
	ID         uint
	UserId     int64
	ChatId     int64
	MessageId  int
	Message    string
	When       int64
	IsReported bool
}

type parsedMessage struct {
	text string
	when int64
}

func remind(bot *tgbotapi.BotAPI, store storage) {
	for {
		reminders, err := store.find()

		if err != nil {
			logrus.Error(err)
			continue
		}

		for _, remind := range reminders {
			remind := remind
			go func() {
				_, err := bot.Send(remind.createMessage())

				if err != nil {
					logrus.Error(err)
					return
				}

				remind.IsReported = true
				err = store.save(remind)

				if err != nil {
					logrus.Error(err)
					return
				}
			}()
		}

		time.Sleep(1 * time.Minute)
	}
}

func parse(text string) *parsedMessage {
	return &parsedMessage{text: "not implemented yet", when: 12345678}
}

func (r *reminder) createMessage() *tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(r.ChatId, defaultText+r.Message)
	msg.ReplyToMessageID = r.MessageId

	return &msg
}

func createReminder(message *tgbotapi.Message, store storage) *tgbotapi.MessageConfig {
	response := "reminder created"
	parsedMessage := parse(message.Text)

	reminder := &reminder{
		UserId:     message.From.ID,
		ChatId:     message.Chat.ID,
		MessageId:  message.MessageID,
		Message:    parsedMessage.text,
		When:       parsedMessage.when,
		IsReported: false,
	}

	err := store.save(reminder)

	if err != nil {
		logrus.Error(err)
		response = "unexpected error occurred"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, response)
	msg.ReplyToMessageID = message.MessageID

	return reminder.createMessage()
}
