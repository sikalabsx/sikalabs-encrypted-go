package encrypted

import (
	"github.com/sikalabsx/sikalabs-encrypted-go/internal/decrypt"
)

type TelegramConfig struct {
	Token  string
	ChatID string
}

func GetOndrejSikaCITelegramConfig() (TelegramConfig, error) {
	token, err := decrypt.Decrypt(TELEGRAM_ONDREJSIKA_CI_NOTIFICATION_TOKEN)
	if err != nil {
		return TelegramConfig{}, err
	}

	chatID, err := decrypt.Decrypt(TELEGRAM_ONDREJSIKA_CI_NOTIFICATION_CHAT_ID)
	if err != nil {
		return TelegramConfig{}, err
	}

	return TelegramConfig{
		Token:  token,
		ChatID: chatID,
	}, nil
}
