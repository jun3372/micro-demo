package err

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
)

type IError interface {
	i()
	WithData(data interface{}) IError
	WithId(id string) IError
	String() string
	ToString() string
}

type Error struct {
	Err  int         `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	id   string      `json:"id"`
}

func NewError(err int, msg string) *Error {
	return &Error{Err: err, Msg: msg}
}

func (e Error) i() {}

func (e Error) WithData(data interface{}) IError {
	e.Data = data
	return e
}

func (e Error) WithId(id string) IError {
	e.id = id
	return e
}

func (e Error) String() string {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	content, _ := json.Marshal(&e)
	return cast.ToString(content)
}

func (e Error) ToString() string {
	return e.String()
}

func (e Error) Error() string {
	return e.String()
}
