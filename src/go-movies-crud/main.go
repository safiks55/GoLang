package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// movie and director are corelated
// assuming that every movie has a director
//assuming one movie will only have one director

type Movie struct {
	ID       string    `json:"id"`       // movie id
	ISBN     string    `json:"isbn"`     // unique id to identify a movie
	Title    string    `json:"title"`    // movie title
	Director *Director `json:"director"` // associating a director to a movie using pointers
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...) // just append the next movie to the current movie
			break
		}

	}
	json.NewEncoder(w).Encode(movies) // returns all the remaining movies after deleting
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) // this will return the new movie
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...) // just append the next movie to the current movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // return the updated movie
			return
		}
	}
}

func main() {
	movies = append(movies, Movie{ID: "1", ISBN: "78613121", Title: "The Batman", Director: &Director{FirstName: "Matt", LastName: "Reeves"}})
	movies = append(movies, Movie{ID: "2", ISBN: "45117191", Title: "Dune", Director: &Director{FirstName: "Denis", LastName: "Villeneuve"}})
	movies = append(movies, Movie{ID: "3", ISBN: "9653447", Title: "Shutter Island", Director: &Director{FirstName: "Martin", LastName: "Scorsese"}})
	movies = append(movies, Movie{ID: "4", ISBN: "12666190", Title: "The Prestige", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "5", ISBN: "08699146", Title: "John Wick ", Director: &Director{FirstName: "Chad", LastName: "Stahelski"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at PORT:8800\n")
	log.Fatal(http.ListenAndServe(":8800", r))
}
