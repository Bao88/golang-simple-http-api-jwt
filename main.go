package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// type Users struct {
// 	Name string 'json:"Name"'
// 	Password string 'json:"Password"'
// }

// type Users []User

// func homePage(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Welcome to the Homepage!")
// 	fmt.Println("Endpoint Hit: homepage")
// }

// func handleRequests() {
// 	// http.HandleFunc("/", homePage)
// 	// Serve the Invoice form to clients
// 	http.Handle("/", http.FileServer(http.Dir("./views/invoice")))
	
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }

func UserLogin(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the UserLogin!")
	fmt.Println("Implement UserLogin")
}

func CreateInvoice(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the CreateInvoice!")
	fmt.Println("Implement CreateInvoice")
}

func main(){
	// fmt.Printf("hello, world\n")
	fmt.Println("Started")
	router := mux.NewRouter();
	router.Handle("/", http.FileServer(http.Dir("./views/invoice")))
	router.HandleFunc("/login", UserLogin).Methods("GET")
	router.HandleFunc("/invoice", CreateInvoice).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", router))
	
	// handleRequests()
}