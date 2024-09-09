package test

import "testing"

type testMockeable struct{}

func (t testMockeable) GetFuncControls() []*CallsFuncControl {
	return []*CallsFuncControl{
		{
			funcName:            "mock",
			funcCalls:           1,
			ExpectedCalls:       1,
			IgnoreCallAssertion: false,
		},
	}
}

func TestAssertControls(t *testing.T) {
	AssertControls(t, testMockeable{})
}