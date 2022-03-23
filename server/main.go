package main

import (
	"fmt"
	"groupie-tracker/controller/handlers"
	"net/http"
)

func main() {

	// Donner accès à notre dossier assets (CSS & Images)
	fs := http.FileServer(http.Dir("./view/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Définir les fonctions et les routes
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/artists", handlers.ArtistsHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	fmt.Println("Server started at http://localhost:8080")

	http.ListenAndServe("localhost:8080", nil)
}
