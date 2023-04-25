package main

import (
	"log"
	"os"
	"time"

	"github.com/nesvoboda/slotwatch/notify"
	"github.com/nesvoboda/slotwatch/slot"
)

func main() {
	proejctName := os.Getenv("PROJECT_NAME")
	teamId := os.Getenv("TEAM_ID")
	cookie := os.Getenv("COOKIE")

	if proejctName == "" {
		log.Fatal("PROJECT_URL is not set")
	}
	if teamId == "" {
		log.Fatal("TEAM_ID is not set")
	}
	if cookie == "" {
		log.Fatal("COOKIE is not set")
	}
	notify.Init()

	watch(proejctName, teamId, cookie)
}

// a global cache to store the slots
var cache = make(map[string]slot.Slot)

func watch(proejctName string, teamId string, cookie string) {
	for {
		slots := slot.GetAll(proejctName, teamId, cookie)
		for _, slot := range slots {
			if _, ok := cache[slot.Id]; !ok {
				cache[slot.Id] = slot
				notify.Send(slot)
			}
		}
		time.Sleep(time.Second * 3)
	}
}
