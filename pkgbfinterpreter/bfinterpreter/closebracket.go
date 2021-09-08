package bfinterpreter

import "fmt"

type closeBracket struct {
}

func NewCloseBracket() closeBracket {
	return closeBracket{}
}

func (a closeBracket) Expr(bs *ByteStack) (err error) {
	t, err := bs.Pop()
	if err != nil {
		return fmt.Errorf("Failed to read bystack: %s", err.Error())
	}
	if t == 0 {
		escape = false
	}
	bs.Push(t)
	return nil
}
