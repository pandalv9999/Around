package main

import (
	"fmt"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("started-service")
	getEnvVars()
	mySignInKey := os.Getenv("SECRET_KEY")

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySignInKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r := mux.NewRouter()

	r.Handle("/upload", jwtMiddleware.Handler(http.HandlerFunc(uploadHandler))).Methods("POST", "OPTIONS")
	r.Handle("/search", jwtMiddleware.Handler(http.HandlerFunc(searchHandler))).Methods("GET", "OPTIONS")
	r.Handle("/signup", http.HandlerFunc(signUpHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(signInHandler)).Methods("POST", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", r))

}
