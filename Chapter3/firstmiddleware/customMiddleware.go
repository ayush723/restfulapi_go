package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before reqquest phase!")
		//pass control back to handler
		handler.ServeHTTP(w, r)
		fmt.Println("executing middleware ater response phase!")
	})
}
func mainLogic(w http.ResponseWriter, r *http.Request) {
	//business logic goes here
	fmt.Println("executing middleware....")
	w.Write([]byte("OK"))
}
func main() {
	//handlerfunc returns http handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}
