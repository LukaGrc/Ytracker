package model

// Main

type MainStruct struct {
	Artists []Artist
}

// Artists

var ArtistTab []Artist

type ArtistsStruct struct {
	TabArt []Artist
}

var Artiste Artist

type ArtistStruct struct {
	Art Artist
	Rel Relations
}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Dates

var DatesTab DateStruct

type DateStruct struct {
	Index []Date `json:"index"`
}

type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Locations

var LocationTab LocationStruct

type LocationStruct struct {
	Index []Locations `json:"index"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Relations

var Relation Relations

type RelationStruct struct {
	Index []Relations `json:"index"`
}

type Relations struct {
	Id int `json:"id"`
	// DatesLocations []string `json:"dateslocations"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
