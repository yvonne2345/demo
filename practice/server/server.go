package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"web.com/server/rpcexample"
)

func main() {
	//register Arith object as a service
	Service := new(rpcexample.Service)
	err := rpc.Register(Service)
	if err != nil {
		log.Fatalf("Format of service Arith isn't correct. %s", err)
	}
	rpc.HandleHTTP()
	//Service.Init()
	//start listening for messages on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
