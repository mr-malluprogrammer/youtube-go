package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	ISB    string  `json:"isb"`
	Name   string  `json:"name"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"fName"`
	LastName  string `json:"lName"`
}

var books []Book

// Easy Part

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := r.URL.Query().Get("id")
	params := mux.Vars(r)
	// fmt.Println(params)
	for index, item := range books {
		if item.ID == params["id"] {
			fmt.Println(books[index+1:])
			books = append(books[:index], books[index+1:]...)
			fmt.Fprintf(w, "Successfully Removed %s\n", item.Name)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := r.URL.Query().Get("id")
	params := mux.Vars(r)
	// fmt.Println(params)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()

	jsonFile, err := os.Open("sample.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &books)

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book", createBook).Methods("POST")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/book/{id}", updateBook).Methods("PUT")

	fmt.Println("Server started: http://localhost:5020/")
	log.Fatal(http.ListenAndServe(":5020", r))

}
