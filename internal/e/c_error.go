package e

type CError struct {
	code int
	msg  string
}

func NewError(code int, msg string) *CError {
	return &CError{code: code, msg: msg}
}

func CodeError(code int) *CError {
	return NewError(code, GetMsg(code))
}

func DefaultError(msg string) *CError {
	return NewError(Exception, msg)
}

func EmptyDataError() *CError {
	return CodeError(EmptyData)
}

func (e *CError) Error() string {
	return e.msg
}

func (e *CError) Code() int {
	return e.code
}
