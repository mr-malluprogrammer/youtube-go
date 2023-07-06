package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

type Book struct{
	ID string	`json:"id"`
	ISB	string	`json:"isb"`
	Name string	`json:"name"`
	Author	*Author	`json:"author"`
}

type Author struct{
	FirstName	string	`json:"fName"`
	LastName	string	`json:"lName"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books=append(books[:index], books[index+1:]...)
			fmt.Fprintf(w, "Successfully Removed %s\n", item.Name)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}



func main() {
	r := mux.NewRouter()

	jsonFile, err := os.Open("sample.json")
	if(err!=nil){
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &books)


	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server started: http://localhost:5020/")
	log.Fatal(http.ListenAndServe(":5020",r))

}
