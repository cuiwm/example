package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quo struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing error")
	}
	//sync call
	args := &Args{7, 8}
	args2 := Args{1, 1}
	fmt.Println(args)
	fmt.Println(args2)
	fmt.Println(*args)
	fmt.Println("&args", &args)
	fmt.Println("&args2", &args2)
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error", err)
	}
	fmt.Println("Arith %d*%d=%d\n", args.A, args.B, reply)

	q := new(Quo)
	divCall := client.Go("Arith.Divide", args, q, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Fatal("arith eroor", replyCall.Error)
	}
	fmt.Printf("Arith %d/%d = %d ..%d", args.A, args.B, q.Quo, q.Rem)
}
