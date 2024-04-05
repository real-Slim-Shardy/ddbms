package lb

import (
	"ddbms/lib/serv"
	"errors"
	"net/rpc"
)

func SendCfgToServer(address string, request string, args serv.OpArgs) error {
	// DialHTTP connects to an HTTP RPC server at the specified network
	client, err := rpc.DialHTTP("tcp", ":10010")
	if err != nil {
		return errors.New("Client connection error: ")
	}
	defer client.Close()

	// Sending the arguments and reply variable address to the server as well
	var str string
	err = client.Call(request, args, &str)
	if err != nil {
		return err
	}
	return nil
}
