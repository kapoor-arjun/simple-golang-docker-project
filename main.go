package main

import(
	"fmt"
	"log"
	"net/http"
	"html"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "HELLO, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("welcome", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi, Welcome")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}