package handlers

import (
	"fmt"
	"groupie-tracker/controller/api"
	"groupie-tracker/controller/functions"
	"groupie-tracker/model"
	"html/template"
	"net/http"
	"strings"
)

func NotFound(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	files := []string{"view/html/404.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	t.Execute(w, nil)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	files := []string{"view/html/about.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	t.Execute(w, nil)
}

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	files := []string{"view/html/index.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	api.ApiRequestArtists()
	artists := functions.RandArtist()

	var TabArtists []model.Artist

	for i := 0; i < 4; i++ {
		TabArtists = append(TabArtists, model.ArtistTab[artists[i]])
	}

	data := model.MainStruct{
		Artists: TabArtists,
	}

	t.Execute(w, data)
}

///////////////////////////////////////////////////////////////

func ArtistsHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artists" {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	files := []string{"view/html/artists.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	api.ApiRequestArtists()

	data := model.ArtistsStruct{
		TabArt: model.ArtistTab,
	}

	t.Execute(w, data)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artist" {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		errorText := fmt.Sprintf("%v - %v", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		http.Error(w, errorText, http.StatusMethodNotAllowed)
		return
	}

	files := []string{
		"view/html/artist.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	r.ParseForm()
	id := r.Form.Get("id")

	api.ApiRequestArtist(id)
	api.ApiRequestRelation(id)

	if model.Artiste.Id == 0 {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	data := model.ArtistStruct{
		Art: model.Artiste,
		Rel: model.Relation,
	}

	t.Execute(w, data)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/search" {
		NotFound(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		errorText := fmt.Sprintf("%v - %v", http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		http.Error(w, errorText, http.StatusMethodNotAllowed)
		return
	}

	files := []string{
		"view/html/search.html",
		"view/html/templates/head.html",
		"view/html/templates/footer.html",
		"view/html/templates/header/title-area.html",
		"view/html/templates/header/navbar.html"}
	t := template.Must(template.ParseFiles(files...))

	r.ParseForm()
	terms := r.Form.Get("terms")

	api.ApiRequestArtists()

	var TabArtists []model.Artist

	for i := 0; i < len(model.ArtistTab); i++ {
		if strings.Contains(strings.ToLower(model.ArtistTab[i].Name), strings.ToLower(terms)) {
			TabArtists = append(TabArtists, model.ArtistTab[i])
		} else {
			for _, member := range model.ArtistTab[i].Members {
				if strings.Contains(strings.ToLower(member), strings.ToLower(terms)) {
					TabArtists = append(TabArtists, model.ArtistTab[i])
					break
				}
			}
		}

	}

	data := model.MainStruct{
		Artists: TabArtists,
	}

	t.Execute(w, data)
}
