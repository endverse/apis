package errors

type Errors interface {
	Error() string // extends go.error
	Code() int64
	Message() string
}

var _ error = &Error{}

type Error struct {
	code    int64
	message string
	err     error
}

func New(code int64, message string, err error) *Error {
	return &Error{
		code:    code,
		message: message,
		err:     err,
	}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Code() int64 {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}
