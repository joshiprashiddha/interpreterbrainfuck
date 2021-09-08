package bfinterpreter

var Tokens = map[string]Token{
	"+": NewAdd(),
	"-": NewSub(),
	"[": NewOpenBracket(),
	"]": NewCloseBracket(),
}

type Token interface {
	Expr(bs *ByteStack) (err error)
}

func AddToken(symbol string, token Token) {
	Tokens[symbol] = token
}
