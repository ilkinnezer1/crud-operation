package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Struct used instead of db

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{
		Id: "1", Isbn: "312133", Title: "Interstellar", Director: &Director{
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	})
	movies = append(movies, Movie{
		Id: "2", Isbn: "2313134", Title: "V For Vendetta", Director: &Director{
			FirstName: "James",
			LastName:  "McTequie",
		},
	})
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
