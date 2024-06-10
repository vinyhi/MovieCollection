package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

var movies []Movie

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

var port = os.Getenv("PORT")

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", CreateMovie).Methods("POST")
	router.HandleFunc("/movies", GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Movie{})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]

			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}