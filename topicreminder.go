package main

import (
	"fmt"
	//"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

// kittens function
func KittensHandler(rw http.ResponseWriter, req *http.Request) {

	rw.Header().Set("Content-Type", "applicatio/json")
	rw.Write([]byte(`{"Kittens":[{"id":1, "name":"fluffy"}, {"id":2, "name":"pinky"}]}`))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hi world %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "There is nothing here as of now, but more to come !\n")
}

func topicHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello world %s\n", r.URL.Path[1:])
	fmt.Fprintf(w, "There is nothing here as of now, but more to come !\n")

}

func main() {

	fmt.Println("Welcome to Topic Reminder")
	fmt.Println("Application Started .. listening on 8080.. ")

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/topics", topicHandler)
	//r.HandleFunc("/speaker", speakerHandler)
	r.HandleFunc("/api/kittens", KittensHandler).Methods("GET")

	//http.Handle("/", "r")
	// Inside the listenandserve() call provide the 'r', stands for router.
	http.ListenAndServe(":8080", r)

}
