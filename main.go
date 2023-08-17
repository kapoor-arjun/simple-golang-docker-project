package main

import(
	"fmt"
	"log"
	"net/http"
	"html"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)){
		fmt.FPrintf(w, "HELLO, %q", html.EscapeString(r.URL.Path))
	}

	http.HandleFunc("hi", func(w http.ResponseWriter, r *http.Request)){
		fmt.FPrintf(w, "hi")
	}

	log.Fatal(http.ListenAndServe(":8081", nil))
}