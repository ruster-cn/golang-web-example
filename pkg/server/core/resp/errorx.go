package resp

type BaseError struct {
	code int
	msg  string
}

func (p *BaseError) Error() string {
	return p.msg
}

func (p *BaseError) Code() int {
	return p.code
}

func NewError(code int, msg string) *BaseError {
	return &BaseError{code: code, msg: msg}
}

func NewCodeError(code int) *BaseError {
	return &BaseError{code: code, msg: errMap[code]}
}
