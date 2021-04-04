package main

import (
	jsonparse "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

//args holds arguments passed to json rpc service
type Args struct {
	Id string
}

//book struct holds book json structure
type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}

type JSONServer struct{}

//give book details

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	// read JSON file and load data
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}
	//unmarshal JSON raw data into books array
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// iterate over each book to find the given book

	for _, book := range books {
		if book.Id == args.Id {
			// if book found , fill the reply with it
			*reply = book
			break
		}
	}
	return nil
}
func main() {
	//create a new rpc server
	s := rpc.NewServer() // register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// register the service by creating a new json server
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}
