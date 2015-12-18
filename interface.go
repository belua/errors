package errors

type StackError interface {
	error
	Stack() []byte
	StackFrames() []StackFrame
	ErrorStack() string
	TypeName() string
}
