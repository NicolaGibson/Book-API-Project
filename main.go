package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title string `json:"Title"`
	Author string `json:"Author"`
	Genre string `json:"Genre"`
}
//Type Books is an array holding several instances of article
type Books []Book

func Children(w http.ResponseWriter, r *http.Request){
	books := Books{
	Book{Title:"Alice in Wonderland", Author:"Lewis Carroll", Genre:"Children"},
		{Title:"The Chronicles of Narnia", Author:"C.S. Lewis", Genre:"Children"},
		 {Title:"The Secret Garden", Author:"Frances Hodgson Burnett", Genre:"Children"},
		{Title:"The Little Princess", Author:"Frances Hodgson Burnett", Genre:"Children"},
		{Title:"The Hobbit", Author:"J.R.R. Tolkein", Genre:"Children"},
	}
	fmt.Println("Endpoint hit: Children's books")
	json.NewEncoder(w).Encode(books)
}

func Young_Adult(w http.ResponseWriter, r *http.Request){
	books := Books{
		Book{Title:"The Fault in Our Stars", Author:"Angie Thoomas", Genre:"Young Adult"},
		Book{Title:"The Hate U Give", Author:"John Green", Genre:"Young Adult"},
		Book{Title:"The Book Thief", Author:"Markus Zusak", Genre:"Young Adult"},
		Book{Title:"Eleanor & Park", Author:"Rainbow Rowell", Genre:"Young Adult"},
		Book{Title:"Children of Blood and Bone", Author:"Tomi Adeyemi", Genre:"Young Adult"},
	}
	fmt.Println("Endpoint hit: Young adult books")
	json.NewEncoder(w).Encode(books)
}

func Horror (w http.ResponseWriter, r *http.Request){
	books := Books{
		Book{Title:"Dracula", Author:"Bram Stoker", Genre:"Horror"},
		Book{Title:"Frakenstein", Author:"Mary Shelley", Genre:"Horror"},
		Book{Title:"It", Author:"Stephen King", Genre:"Horror"},
		Book{Title:"The Exorcist", Author:"William Peter Blatty", Genre:"Horror"},
		Book{Title:"The Woman in Black", Author:"Susan Hill", Genre:"Horror"},
	}
	fmt.Println("Endpoint hit: Horror books")
	json.NewEncoder(w).Encode(books)
}
func Homepage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Homepage hit")
}

func handleRequests(){
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/children", Children)
	http.HandleFunc("/YA", Young_Adult)
	http.HandleFunc("/horror", Horror)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main(){
	handleRequests()
}
