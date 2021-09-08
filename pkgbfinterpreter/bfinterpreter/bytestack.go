package bfinterpreter

import (
	"errors"
	"strconv"
)

type ByteStack struct {
	items []int
}

func NewByteStack(size int) *ByteStack {
	return &ByteStack{
		items: make([]int, size),
	}
}

// push adds items to end of stack
func (rs *ByteStack) Push(r int) {
	rs.items = append(rs.items, r)
}

// pop removes items from end of stack
func (rs *ByteStack) Pop() (int, error) {
	l := len(rs.items) - 1
	if l < 0 {
		return 0, errors.New("pop empty stack nothing to remove")
	}
	item := rs.items[l]
	rs.items = rs.items[:l]
	return item, nil
}

// Display
func (rs *ByteStack) Display() string {
	var result string
	for {
		l := len(rs.items) - 1
		if l < 0 {
			break
		}
		item := rs.items[l]
		rs.items = rs.items[:l]
		result = result + strconv.Itoa(item)
	}
	return result
}
