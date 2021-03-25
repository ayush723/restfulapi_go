package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Queryhandler(w http.ResponseWriter, r *http.Request) {
	//fetch query parameters as a map
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "GOt parameter id :%s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category :%s!", queryParams["category"][0])
}

func main() {
	// create a new router
	r := mux.NewRouter()
	//attach a path with handler
	r.HandleFunc("/articles", Queryhandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		//good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
