package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/model"
	"io/ioutil"
	"net/http"
)

func ApiRequestLocation() {

	req, errh := http.Get("https://groupietrackers.herokuapp.com/api/locations")

	if errh != nil {
		fmt.Println(errh)
	}

	x, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(x, &model.LocationTab)
}
