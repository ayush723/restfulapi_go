package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func main() {
	// create a web service
	webservice := new(restful.WebService)
	//create a route and attach it to the handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// add the service to application
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}
func pingTime(req *restful.Request, resp *restful.Response) {
	// write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
