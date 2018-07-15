package meetups

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetMeetups() string {

	resp, err := http.Get("http://api.meetup.com/GDG-Brasilia/events")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var meetups []Meetup
	json.Unmarshal(body, &meetups)

	err = json.Unmarshal(body, &meetups)

	if err != nil {
		log.Panicln(err)

	}

	var texto string

	for indice := range meetups {

		meetupp := meetups[indice]

		t, _ := time.Parse("2006-01-02", meetupp.LocalDate)
		data := t.Format("02/01/2006")

		texto = texto + "Nome: " + meetupp.Name + "\n"
		texto = texto + "Data: " + data + " ás " + meetupp.LocalTime + " Horas" + "\n"
		texto = texto + "Local: " + meetupp.Venue.Address + " - " + meetupp.Venue.City + "\n"
		texto = texto + "Link: " + meetupp.Link
		texto = texto + "\n \n"
	}

	return string(texto)
}
