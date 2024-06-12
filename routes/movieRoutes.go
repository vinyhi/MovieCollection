package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"hash/fnv"
)

var movies []Movie

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

var port = os.Getenv("PORT")
var dataVersion uint32 // simplistic versioning/checksum

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", CreateMovie).Methods("POST")
	router.HandleFunc("/movies", GetMovies).Methods("GET", "HEAD") // HEAD for checking version
	router.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteCar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func updateDataVersion() {
	// A rudimentary way of creating a version/checksum of the movies slice
	h := fnv.New32a()
	for _, m := range movies {
		h.Write([]byte(m.ID + m.Title + m.Year))
	}
	dataVersion = h.Sum32()
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie) // Adding the movie to the slice
	updateDataVersion()            // Updating the version after modification

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ETag", fmt.Sprintf("%v", dataVersion)) // Set ETag to current data version
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
	found := false
	for index, item := range movies {
		if item.ID == params["id"] {
			found = true
			movies = append(movies[:index], movies[index+1:]...) // removing old
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie) // adding updated
			updateDataVersion()
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	found := false
	for index, item := range movies {
		if item.ID == params["id"] {
			found = true
			movies = append(movies[:index], movies[index+1:]...)
			updateDataVersion()
			break
		}
	}

	if !found {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}