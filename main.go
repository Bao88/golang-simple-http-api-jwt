package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

// source https://gist.github.com/otiai10/22ad21fbe48f37f14c6b2218e9d110a5
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

const user = "test"
const password ="test"
const userAge = 30

func createTokenString() string {
	// Embed User information to `token`
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
		Name: user,
		Age: userAge,
	})
	// token -> string. Only server knows this secret (foobar).
	tokenstring, err := token.SignedString([]byte("foobar"))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstring
}

func UserLogin(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" { // Accepts only POST
		r.ParseForm()

		name := r.FormValue("name")
		pass := r.FormValue("password")
		if name == user && password == pass {
			// Create token and send it back
			tokenstring := createTokenString() 
			fmt.Fprintf(w, tokenstring)
		} else {
			fmt.Fprintf(w, "Fail")
		}
	} else {
		fmt.Fprintf(w, "Currently supporting only POST /login")
	}
}

func postInvoice(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" { // Accepts only POST
		// Test if user is authirized to access this protected area
		clientToken := r.Header.Get("Authorization")

		token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("foobar"), nil
		})
		// log.Println(token.Claims, err)

		if token.Valid {
			// Autorization success, store json to file
			r.ParseForm()
			// We only wants the key, which is the JSON element
			// Surely there are a better way to access the key only without using a for loop and range.
			for k := range r.Form {
				fmt.Println(k)
				err = ioutil.WriteFile("output.txt", []byte(k), 0644)
				if err != nil {
					panic(err)
				}
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "JWT Autorization failed")
		}	
	} else {
		fmt.Fprintf(w, "Currently supporting POST /invoice only")
	}
}

func main(){
	// Routing
	fmt.Println("Server started. Available at localhost:8000")
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("main/"))))
	http.HandleFunc("/login", UserLogin)
	http.HandleFunc("/invoice", postInvoice)
	http.Handle("/createInvoice/", http.StripPrefix("/createInvoice/", http.FileServer(http.Dir("views/invoice"))))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

