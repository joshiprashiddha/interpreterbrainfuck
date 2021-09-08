package bfinterpreter

import "fmt"

type Add struct {
}

func NewAdd() Add {
	return Add{}
}

func (a Add) Expr(bs *ByteStack) (err error) {
	t, err := bs.Pop()
	if err != nil {
		return fmt.Errorf("Failed to read bystack: %s", err.Error())
	}
	t = t + 1
	bs.Push(t)
	return nil
}
