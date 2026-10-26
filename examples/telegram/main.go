package main

import (
	"log"

	"github.com/sikalabs/slu/utils/telegram_utils"
	"github.com/sikalabsx/sikalabs-encrypted-go/pkg/encrypted"
)

func main() {
	telegramConfig, err := encrypted.GetOndrejSikaCITelegramConfig()
	handleError(err)

	err = telegram_utils.TelegramSendMessage(
		telegramConfig.Token, telegramConfig.ChatID,
		"TEST: 1: This is a test message from SikaLabs Encrypted Go package.",
	)
	handleError(err)

	telegramConfig, err = encrypted.GetSikalabsNotificationV1TelegramConfig()
	handleError(err)

	err = telegram_utils.TelegramSendMessage(
		telegramConfig.Token, telegramConfig.ChatID,
		"TEST: 2: This is a test message from SikaLabs Encrypted Go package.",
	)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
