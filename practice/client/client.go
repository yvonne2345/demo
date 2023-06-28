package main

import (
	"log"
	"net/rpc"
	"web.com/server/rpcexample"
)

func main() {
	//make connection to rpc server
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &rpcexample.Args{
		A: 2,
		B: 9,
	}
	//this will store returned result
	//var result rpcexample.Result
	var resp rpcexample.Resp
	//call remote procedure with args
	err = client.Call("Service.Init", args, &resp)
	if err != nil {
		log.Fatalf("error in Service", err)
	}
	//we got our result in result
	log.Printf("%s", resp)
}
