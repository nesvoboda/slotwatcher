package main

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/nesvoboda/slotwatch/notify"
	"github.com/nesvoboda/slotwatch/slot"
)

func main() {
	projectName := os.Getenv("PROJECT_NAME")
	teamId := os.Getenv("TEAM_ID")
	cookie := os.Getenv("COOKIE")

	if projectName == "" {
		log.Fatal("PROJECT_NAME is not set")
	}
	if teamId == "" {
		log.Fatal("TEAM_ID is not set")
	}
	if cookie == "" {
		log.Fatal("COOKIE is not set")
	}
	notify.Init()
	log.SetLevel(log.DebugLevel)

	log.Info("Starting watcher")

	watch(projectName, teamId, cookie)
}

// a global cache to store the slots
var cache = make(map[string]slot.Slot)

func watch(projectName string, teamId string, cookie string) {
	for {
		slots := slot.GetAll(projectName, teamId, cookie)
		log.Debug("Got slots: ", slots)
		for _, slot := range slots {
			if _, ok := cache[slot.Id]; !ok {
				cache[slot.Id] = slot
				notify.Send(slot)
				log.Debug("New slot: ", slot.Id, " ", slot.Start)
			}
			log.Debug("Slot already in cache: ", slot.Id, " ", slot.Start)
		}
		time.Sleep(time.Second * 3)
	}
}
