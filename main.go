package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 page not found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not accepted!", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1 style='color: steelblue'>Hello Mr.MalluProgrammer in hello page</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := Book{
		Title:  "The Gunslinger",
		Author: "Mr.MalluProgrammer",
		Pages:  12,
	}
	json.NewEncoder(w).Encode(book)
}

func Formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error in parsing the form: %v", err)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method not accepted!", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	firstName := r.FormValue("fname")
	lastName := r.FormValue("lname")
	fmt.Fprintf(w, "First Name : %s\nLast Name : %s", firstName, lastName)
}

func main() {
	filePath := http.FileServer(http.Dir("./html"))
	http.Handle("/", filePath)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/form", Formhandler)
	http.HandleFunc("/book", GetBook)
	log.Fatal(http.ListenAndServe(":5100", nil))
}
