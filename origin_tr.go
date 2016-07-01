package main

import (
	"fmt"
	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/mux"
	"net/http"
)

//var web = flag.String("web")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "There is nothing here as of now, but more to come !\n")
}

func main() {

	fmt.Println("Welcome to Topic Reminder\n")
	// Using a simple default serv mux
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
