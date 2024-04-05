package serv

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type OpArgs struct {
	N      int
	Status string
}

type Ops struct {
	N      int
	Port   string
	Status string
}

// Receive and set config info from Load Balancer
func (o *Ops) SetConfig(args OpArgs, s *string) error {
	o.N = args.N
	o.Status = args.Status
	log.Println(o.N, o.Port, o.Status)

	return nil
}

// Reads port number from programm arguments
func ReadPortNumber() (s string, e error) {

	// Check if we have an arguments -> Args[0] is the program name
	if len(os.Args) == 1 {
		e = errors.New("Not enough arguments to start program!\nPlease specify on which port server should be hosted\n")
		return "", e
	}

	// If the 1st argument exists
	s = os.Args[1]

	// Try to convert argument value to Int
	port, e := strconv.Atoi(s)

	// If fail - incorrect type
	if e != nil {
		e = errors.New("Invalid type of argument!\nPlease write number on which port server should be hosted\n")
		return "", e
	}

	// Check if number is in range [1, 65535]
	if port <= 0 || port > 65535 {
		e = errors.New("Invalid value of port number!\nPlease set port number from range [1, 65535]\n")
		return "", e
	}

	// Now all validations are passed, so port number can be used in porogram
	return s, nil
}
