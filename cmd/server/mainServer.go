package main

import (
	"ddbms/lib/manager"
	"ddbms/lib/serv"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func main() {

	// init Logger
	manager.InitLogger()

	// init port number
	port, err := serv.ReadPortNumber()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// init Ops - Operational Parameters
	opServer := new(serv.Ops)
	opServer.Port = port
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(opServer.N, opServer.Port, opServer.Status)

	err = rpc.Register(opServer)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Start Listener
	// Registers an HTTP handler for RPC messages
	rpc.HandleHTTP()

	// Start listening for the requests on port 1234
	listener, err := net.Listen("tcp", ":"+opServer.Port)
	if err != nil {
		log.Fatal("Listener error: ", err)
	}
	http.Serve(listener, nil)
}
