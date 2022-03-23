package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/model"
	"io/ioutil"
	"net/http"
)

func ApiRequestArtists() {

	req, errh := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if errh != nil {
		fmt.Println(errh)
	}

	x, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &model.ArtistTab)
}

func ApiRequestArtist(id string) {

	req, errh := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)

	if errh != nil {
		fmt.Println(errh)
	}

	x, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &model.Artiste)

}
