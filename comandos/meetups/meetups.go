package meetups

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
		texto = texto + "Nome: " + meetupp.Name + "\n"
		texto = texto + "Data: " + meetupp.LocalTime + " ás " + meetupp.LocalTime + "Horas" + "\n"
		texto = texto + "Local: " + meetupp.Venue.Name + " - " + meetupp.Venue.Address + "\n"
		texto = texto + "Link: " + meetupp.Link
		texto = texto + "\n \n"
	}

	return string(texto)
}
