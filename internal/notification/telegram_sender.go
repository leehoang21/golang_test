package notification

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (s *Sender) SendTelegramMessage(message string) error {
	endpoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.teleConfig.BotToken)
	data := url.Values{}
	data.Set("chat_id", s.teleConfig.ChanelID)
	data.Set("text", message)
	data.Set("parse_mode", "Markdown")

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body:", err)
		}
	}(resp.Body)
	return nil
}
