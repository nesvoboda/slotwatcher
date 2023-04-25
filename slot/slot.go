package slot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Slot struct {
	Ids   string    `json:"ids"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Id    string    `json:"id"`
	Title string    `json:"title"`
}

func GetAll(projectUrl string, teamId string, token string) []Slot {
	// get the slots for the next 3 days
	start := time.Now()
	end := start.Add(time.Hour * 24 * 3)

	url := fmt.Sprintf(
		"https://projects.intra.42.fr/projects/%s/slots.json?team_id=%s&start=%s&end=%s",
		projectUrl,
		teamId,
		start.Format("2006-01-02"),
		end.Format("2006-01-02"),
	)

	cookie := http.Cookie{
		Name:  "_intra_42_session_production",
		Value: token,
	}

	slots := makeReq(url, &cookie)
	return slots
}

// make a http get request with a custom session cookie
func makeReq(url string, sessionCookie *http.Cookie) []Slot {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(sessionCookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var slots []Slot
	err = json.NewDecoder(resp.Body).Decode(&slots)
	if err != nil {
		log.Fatal(err)
	}
	return slots
}
