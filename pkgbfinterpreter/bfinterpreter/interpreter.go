package bfinterpreter

import (
	"fmt"
	"reflect"
	"strings"
)

const SIZE = 10

var escape bool

func GetTokens() string {
	keys := reflect.ValueOf(Tokens).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strings.Join(strkeys, ",")
}

func Interpret(input string) string {
	bs := NewByteStack(SIZE)
	for _, r := range input {
		if string(r) == "\n" {
			break
		}
		if escape {
			if string(r) == "]" {
				err := Tokens[string(r)].Expr(bs)
				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			err := Tokens[string(r)].Expr(bs)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return bs.Display()
}
