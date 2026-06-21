package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-TYpe", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

// func updateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	m = json.NewDecoder(r.Body).Decode(&movie)

// 	for index, movie := range movies {
// 		if movie.ID == params["id"] {
// 			movies[index] = Movie{m}
// 			return
// 		}
// 	}
// 	http.Error(w, "Movie not found", http.StatusNotFound)
// }

var movie1 Movie = Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}}
var movie2 Movie = Movie{ID: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}}

func main() {
	r := mux.NewRouter()

	// Add some movies to the "DB"
	movies = append(movies, movie1)
	movies = append(movies, movie2)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("movies/{id}", updateMovie).Methods("PUT")
	// r.HandleFunc("movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting movie server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
