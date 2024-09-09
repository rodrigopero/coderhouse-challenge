package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type CallsFuncControl struct {
	funcName            string
	funcCalls           int32
	ExpectedCalls       int32
	IgnoreCallAssertion bool
}

func (c *CallsFuncControl) SetFuncName(name string) {
	c.funcName = name
}

func (c *CallsFuncControl) IncreaseCallCount() {
	c.funcCalls++
}

func AssertControls(t *testing.T, mock Mockeable) {
	for _, control := range mock.GetFuncControls() {
		if !control.IgnoreCallAssertion {
			assert.Equal(t, control.ExpectedCalls, control.funcCalls,
				fmt.Sprintf("expected calls for func %s does not match actual", control.funcName))
		}
	}
}