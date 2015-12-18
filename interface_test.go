package errors

import "testing"

func TestInterfaceImplemented(t *testing.T) {
	acceptStackError(NewStackError("Test error"))
}

func acceptStackError(err StackError) {
	acceptStandardError(err)
}

func acceptStandardError(err error) {
}
