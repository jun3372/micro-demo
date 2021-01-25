package err

type Error struct {
	Err int    `json:"err"`
	Msg string `json:"msg"`
}

func NewError(err int, msg string) *Error {
	return &Error{Err: err, Msg: msg}
}

func (e Error) Error() string {
	return e.Msg
}
