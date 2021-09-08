package bfinterpreter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNewSub(t *testing.T) {
	assertions := assert.New(t)
	sub := NewSub()

	if !assertions.Implements(new(Token), sub) {
		t.Fatal("Token Implementation Sub does not honor Token definition")
	}

	if !assertions.NotNil(sub, "sub not initialized") {
		t.Fatal("sub not initialized")
	}

	if !assertions.NotNil(sub, "sub dependency not initialized") {
		t.Fatal("sub dependency not initialized")
	}

}

func TestExpr(t *testing.T) {
	type TestCase struct {
		bs  ByteStack
		err error
	}
	var subTestCases = []TestCase{
		{
			bs:  *NewByteStack(0),
			err: errors.New("Failed to read ByteStack: pop empty stack nothing to remove"),
		},
	}
	t.Run("success", func(t *testing.T) {
		for _, tt := range subTestCases {
			s := NewSub()
			err := s.Expr(&tt.bs)
			assert.EqualError(t, tt.err, err.Error())
		}
	})
}
