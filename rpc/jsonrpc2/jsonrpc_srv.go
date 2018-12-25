package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

//Args: input args
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.B * args.A
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	//注意此处是rpc.Register
	rpc.Register(arith)
	//rpc.HandleHTTP()
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("ResolveTCPAddr fail")
		return
	}
	Listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal("ListenTCP fail", err.Error())
		return
	}

	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Fprint(os.Stderr, "accept error", err.Error())
			continue
		}
		//注意此处是jsonrpc
		jsonrpc.ServeConn(conn)
	}
	//	l, e := net.Listen("tcp", ":1234")
	//	if e != nil {
	//		log.Fatal("listen error", e)
	//	}
	//
	//	http.Serve(l, nil)
}
