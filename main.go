package main

import (
	"fmt"
	"log"
	"net/http"
)

// type Users struct {
// 	Name string 'json:"Name"'
// 	Password string 'json:"Password"'
// }

// type Users []User

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func handleRequests() {
	// http.HandleFunc("/", homePage)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	// fmt.Printf("hello, world\n")
	handleRequests()
}