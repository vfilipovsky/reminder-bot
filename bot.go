package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func start() {
	err := godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env file")
		os.Exit(1)
	}

	store, err := newStorage()

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	defer store.close()

	if err := store.migrate(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))

	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	go remind(bot, store)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := createReminder(update.Message, store)

		if _, err := bot.Send(msg); err != nil {
			logrus.Error(err)
		}
	}
}
