# brainfuckinterpreter
brainfuckinterpreter in golang

**Implementation Details**

Program handles tokens +, -, [, ]
Developer can extend functionality as required.

**How it run**

At root run followings two commands

`go build .`

`./bfinterpreter`

```
======================
Default Interpreter size is 10. 10
Available Tokens: +,-,[,]
Type 'exit' to exit.
======================
bf>

```

Enter sequence of + or - or []
[ will lead to inner loop until ]. Loop calculation is escaped or considered based on stack pointer value at the moment.
Enter "exit" to exit

It will output state of stack.


**How to run tests**

Run `go test -count=1 ./...` to run tests

**Extentend as below**

Create Token which implements Token interface

```
type Token interface {
	Expr(bs *ByteStack) (err error)
}
```

example adding new token `>` is as simple as

```
bfinterpreter.AddToken(">", NewGreater())
```

where Greater implements Token

```
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
```



**Few Example Runs**

```
bf>+++++++++++++++++++++++++++++++++++++
37000000000
bf>+++++++++
9000000000
bf>+++++++++++++++++++------------------
1000000000
bf>+++++++++++++++------
9000000000
bf>[+++++++]++++++++++++++-----[+]
10000000000
bf>[+]++++++++++++-----
7000000000
```




