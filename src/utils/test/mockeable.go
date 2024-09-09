package test

type Mockeable interface {
	GetFuncControls() []*CallsFuncControl
}