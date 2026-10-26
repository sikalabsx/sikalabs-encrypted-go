package encrypted

import (
	"strconv"

	"github.com/sikalabsx/sikalabs-encrypted-go/internal/decrypt"
)

type TelegramConfig struct {
	Token  string
	ChatID int64
}

func GetOndrejSikaCITelegramConfig() (TelegramConfig, error) {
	token, err := decrypt.Decrypt(TELEGRAM_ONDREJSIKA_CI_NOTIFICATION_TOKEN)
	if err != nil {
		return TelegramConfig{}, err
	}

	chatIDStr, err := decrypt.Decrypt(TELEGRAM_ONDREJSIKA_CI_NOTIFICATION_CHAT_ID)
	if err != nil {
		return TelegramConfig{}, err
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return TelegramConfig{}, err
	}

	return TelegramConfig{
		Token:  token,
		ChatID: chatID,
	}, nil
}

func GetSikalabsNotificationV1TelegramConfig() (TelegramConfig, error) {
	token, err := decrypt.DecryptWithSuffix(SIKALABS_NOTIFICATION_V1_TOKEN, "_SIKALABS_NOTIFICATION_V1")
	if err != nil {
		return TelegramConfig{}, err
	}

	chatIDStr, err := decrypt.DecryptWithSuffix(SIKALABS_NOTIFICATION_V1_CHAT_ID, "_SIKALABS_NOTIFICATION_V1")
	if err != nil {
		return TelegramConfig{}, err
	}

	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		return TelegramConfig{}, err
	}

	return TelegramConfig{
		Token:  token,
		ChatID: chatID,
	}, nil
}
