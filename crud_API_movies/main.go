package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "43552", Title: "Movie love", Director: &Director{Firstname: "Aom", Lastname: "Ket"}})
	movies = append(movies, Movie{Id: "2", Isbn: "34552", Title: "Fast felios", Director: &Director{Firstname: "Adum", Lastname: "Hisler"}})
	movies = append(movies, Movie{Id: "3", Isbn: "23332", Title: "Catch me", Director: &Director{Firstname: "John", Lastname: "Siri"}})
	movies = append(movies, Movie{Id: "4", Isbn: "34532", Title: "32Desember", Director: &Director{Firstname: "Futur", Lastname: "Hosne"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8080 \n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
