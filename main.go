package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bfinterpreter/pkgbfinterpreter/bfinterpreter"
)

/*
//For extension
type Greater struct {
}

func NewGreater() *Greater {
	return &Greater{}
}

func (a Greater) Expr(bs *bfinterpreter.ByteStack) (err error) {
	t, err := bs.Pop()
	if err != nil {
		return fmt.Errorf("Failed to read ByteStack: %s", err.Error())
	}
	t = t - 1
	bs.Push(t)
	return nil
}
*/
func main() {
	fmt.Println("======================")
	fmt.Printf("Default Interpreter size is 10. %v\n", bfinterpreter.SIZE)
	fmt.Printf("Available Tokens: %s\n", bfinterpreter.GetTokens())
	fmt.Println("Type 'exit' to exit.")
	fmt.Println("======================")
	//for extension
	//bfinterpreter.AddToken(">", NewGreater())
	fmt.Printf("Available Tokens: %s\n", bfinterpreter.GetTokens())
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("bf>")
		input, _ := reader.ReadString('\n')
		if input == "exit\n" {
			break
		}
		output := bfinterpreter.Interpret(input)
		fmt.Println(output)
	}
}
