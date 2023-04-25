package notify

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
		os.Getenv("TELEGRAM_TOKEN"),
		os.Getenv("TELEGRAM_CHAT_ID"),
		message,
	)
	http.Get(url)
}
