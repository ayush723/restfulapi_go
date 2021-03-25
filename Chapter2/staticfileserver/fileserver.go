package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	//mapping to methods is possible with httprouter

	router.ServeFiles("/static/*filepath", http.Dir("C:/Users/ayush/Documents/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
