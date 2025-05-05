package notification

import "base/config"

type Sender struct {
	teleConfig config.TelegramConfig
}

func NewSender(teleConfig config.TelegramConfig) *Sender {
	return &Sender{
		teleConfig: teleConfig,
	}
}
