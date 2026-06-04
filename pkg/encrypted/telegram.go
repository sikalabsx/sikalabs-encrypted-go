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
