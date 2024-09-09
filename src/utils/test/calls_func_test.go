package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCallsFuncControl_IncreaseCallCount(t *testing.T) {
	c := CallsFuncControl{}
	assert.Equal(t, int32(0), c.funcCalls)

	c.IncreaseCallCount()
	assert.Equal(t, int32(1), c.funcCalls)

	c.IncreaseCallCount()
	assert.Equal(t, int32(2), c.funcCalls)
}

func TestCallsFuncControl_SetFuncName(t *testing.T) {
	c := CallsFuncControl{}
	assert.Equal(t, "", c.funcName)

	funcName := "myFunc"
	c.SetFuncName(funcName)
	assert.Equal(t, funcName, c.funcName)
}