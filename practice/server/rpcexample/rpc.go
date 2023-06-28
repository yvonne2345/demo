package rpcexample

import (
	"fmt"
)

// Args Holds arguments to be passed to service Arith in RPC call
type Args struct {
	A, B int
}

// Arith Representss service Arith with method Multiply
type Service int

// Result of RPC call is of this type
type Result int

type Resp string

// Multiply stores product of args.A and args.B in result pointer
func (t *Service) Multiply(args Args, result *Result) error {
	fmt.Println("调用成功multiply")

	*result = Result(args.A * args.B)
	return nil
}

//func (t *Service) Init(args Args, resp *Resp) error {
//	fmt.Println("调用成功init")
//	*resp = "resp"
//	return nil
//}

func (t *Service) Init(args Args, resp *Resp) error {
	fmt.Println("调用成功init")
	*resp = "调用成功init"
	return nil
}
