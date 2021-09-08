package bfinterpreter

import "fmt"

type Sub struct {
}

func NewSub() *Sub {
	return &Sub{}
}

func (a Sub) Expr(bs *ByteStack) (err error) {
	t, err := bs.Pop()
	if err != nil {
		return fmt.Errorf("Failed to read ByteStack: %s", err.Error())
	}
	t = t - 1
	bs.Push(t)
	return nil
}
