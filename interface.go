package errors

type StackError interface {
	error
	Stack() []byte
	StackFrames() []StackFrame
	ErrorStack() string
}
