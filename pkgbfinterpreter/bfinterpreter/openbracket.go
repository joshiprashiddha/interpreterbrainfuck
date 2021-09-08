package bfinterpreter

import "fmt"

type openBracket struct {
}

func NewOpenBracket() openBracket {
	return openBracket{}
}

func (a openBracket) Expr(bs *ByteStack) (err error) {
	t, err := bs.Pop()
	if err != nil {
		return fmt.Errorf("Failed to read ByteStack: %s", err.Error())
	}
	if t == 0 {
		escape = true
	}
	bs.Push(t)
	return nil
}
