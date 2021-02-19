package global

type Error struct {
	code    int
	status  int
	message string
	inner   *error
}

func NewError(status int, code int, message string, inner ...error) *Error {
	var e *error

	if inner != nil {
		e = &inner[0]
	} else {
		e = nil
	}

	return &Error{
		code:    code,
		status:  status,
		message: message,
		inner:   e,
	}
}

func (o *Error) Error() string {
	if o.message != "" {
		return o.message
	} else if o.inner != nil {
		return (*o.inner).Error()
	}

	return ""
}

func (o *Error) Code() int {
	return o.code
}

func (o *Error) Status() int {
	return o.status
}
