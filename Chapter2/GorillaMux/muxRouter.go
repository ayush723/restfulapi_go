package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//ArticHandler is a function handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	//mux.Vars returns all path parameters as a map
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is : %v\n", vars["category"])
	fmt.Fprintf(w, "Id is: %v\n", vars["id"])
}
func main() {
	//create a new router
	r := mux.NewRouter()
	//attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		//good practice : enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
