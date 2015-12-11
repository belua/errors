package errors

import "testing"

func TestInterfaceImplemented(t *testing.T) {
	acceptStackError(New("Test error"))
}

func acceptStackError(err StackError) {
	acceptStandardError(err)
}

func acceptStandardError(err error) {
}
