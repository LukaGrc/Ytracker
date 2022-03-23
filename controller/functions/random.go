package functions

import (
	"groupie-tracker/model"
	"math/rand"
	"time"
)

func RandArtist() []int {
	var result []int
	var temp int
	for len(result) < 4 {
		rand.Seed(time.Now().UnixNano())
		temp = rand.Intn(len(model.ArtistTab))
		if len(result) == 0 {
			result = append(result, temp)
		} else {
			if !checkInside(result, temp) {
				result = append(result, temp)
			}
		}
	}
	return result
}

func checkInside(list []int, value int) bool {
	for _, i := range list {
		if i == value {
			return true
		}
	}
	return false
}
