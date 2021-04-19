package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Arith int

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (a *Arith) Divide(args *Args, reply *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = args.A / args.B
	reply.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
