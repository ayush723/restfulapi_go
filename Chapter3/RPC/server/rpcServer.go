package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	//fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	//create a new rpc server
	timeserver := new(TimeServer)
	//register rpc server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// listen for requests on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen errr:", e)
	}
	http.Serve(l, nil)
}
