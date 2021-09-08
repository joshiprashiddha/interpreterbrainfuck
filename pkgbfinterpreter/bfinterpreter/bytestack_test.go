package bfinterpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisplay(t *testing.T) {
	type TestCase struct {
		bs     ByteStack
		output string
	}
	var displayTestCases = []TestCase{
		{
			bs:     *NewByteStack(10),
			output: "0000000000",
		},
	}
	t.Run("success", func(t *testing.T) {
		for _, tt := range displayTestCases {
			assert.Equal(t, tt.output, tt.bs.Display())
		}
	})
}
