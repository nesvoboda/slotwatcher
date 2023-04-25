package notify

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"

	"github.com/nesvoboda/slotwatch/slot"
)

func makeMessage(slot slot.Slot) string {
	formatStart := slot.Start.Format("Monday 2 January 15:04")
	formatEnd := slot.End.Format("15:04")

	message := fmt.Sprintf(
		"New slot available: %s - %s",
		formatStart,
		formatEnd,
	)

	return message
}

var botToken string
var chatId string

func Init() {
	botToken := os.Getenv("TELEGRAM_TOKEN")
	chatId := os.Getenv("TELEGRAM_CHAT_ID")

	if botToken == "" {
		log.Fatal("TELEGRAM_TOKEN is not set")
	}
	if chatId == "" {
		log.Fatal("TELEGRAM_CHAT_ID is not set")
	}
}

func Send(slot slot.Slot) {
	message := makeMessage(slot)

	// send a telegram bot message
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken,
		chatId,
		message,
	)
	_, err := http.Get(url)
	if err != nil {
		log.Error(err)
	}
	log.Debug("Sent message: ", message, " to chat: ", chatId)
}
