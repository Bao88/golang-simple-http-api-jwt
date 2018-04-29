package main

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/dgrijalva/jwt-go"
	"strings"
	// "encoding/json"
)

// type Users struct {
// 	Name string 'json:"Name"'
// 	Password string 'json:"Password"'
// }

// type Users []User

func handler(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	Homepage(w, req)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Homepage(w http.ResponseWriter, r *http.Request){

	http.ServeFile(w, r, "views/login/login.html")
	fmt.Println("Endpoint Hit: homepage")
	fmt.Println(r.Method);
}

// func handleRequests() {
// 	// http.HandleFunc("/", homePage)
// 	// Serve the Invoice form to clients
// 	http.Handle("/", http.FileServer(http.Dir("./views/invoice")))
	
// 	log.Fatal(http.ListenAndServe(":8081", nil))
// }



func UserLogin(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w, "Welcome to the UserLogin!")
	fmt.Println("Implement UserLogin")
	fmt.Println(r.Method);
	fmt.Println(r.Body);
	fmt.Println(r);
	r.ParseForm()

	fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Welcome to the UserLogin!")
    // fmt.Fprintf(w, "Hello OK!") // write data to response
	// http.Redirect("/createInvoice")
	// http.ServeFile(w, r, "./views/invoice/index.html")
	// http.ServeFile(w, r, "./views/invoice/scripts/main.js")
	// http.FileServer(http.Dir("./views/invoice/scripts"))
	// http.Redirect(w, r, "/createInvoice", 302)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the CreateInvoice!")
	fmt.Println("Implement CreateInvoice")
	fmt.Println(r.Method);
}

type myData struct {
	Owner string
	Name  string
}

func postInvoice(w http.ResponseWriter, r *http.Request){


	fmt.Fprintf(w, "Welcome to the PostInvoice!")
	fmt.Println("Implement PostInvoice")
	// fmt.Println(r.Method);
	// fmt.Println(r.Body);
	// fmt.Println(r);

	// r.ParseForm()

	// fmt.Println(r.Form) // print information on server side.
    // fmt.Println("path", r.URL.Path)
    // fmt.Println("scheme", r.URL.Scheme)
    // fmt.Println(r.Form["url_long"])
    // for k, v := range r.Form {
    //     fmt.Println("key:", k)
    //     fmt.Println("val:", strings.Join(v, ""))
	// }
	// fmt.Println(r.Form)
	// fmt.Println("Implement UserLogin")
	fmt.Println(r.Method);
	fmt.Println(r.Body);
	fmt.Println(r);
	r.ParseForm()

	fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
	}
	// fmt.Fprintf(w, "Welcome to the UserLogin!")
}

func main(){
	fmt.Println("Started")
	http.Handle("/", http.FileServer(http.Dir("./main")))
	// http.HandleFunc("/", handler)
	// http.Handle("/", http.FileServer(http.Dir("views/invoice")))
	http.HandleFunc("/login", UserLogin)
	http.HandleFunc("/invoice", postInvoice)
	// http.Handle("/createInvoice/", http.StripPrefix("/createInvoice/", http.FileServer(http.Dir("views/invoice"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
	
}

