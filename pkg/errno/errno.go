package errno

import (
	"fmt"
)

type Errno struct {
    HTTP    int
    Code    string
    Message string
}

func (err *Errno) Error() string {
    return err.Message
}

func (err *Errno) WithMessage(format string, args ...interface{}) *Errno {
    err.Message = fmt.Sprintf(format, args...)
    return err
}

// Decode 尝试从 err 中解析出业务错误码和错误信息.
func Decode(err error) (int, string, string) {
    if err == nil {
        return OK.HTTP, OK.Code, OK.Message
    }

	/*
		Err := Errno{}
		// 如果是上述定义的几种错误类型，则转换
		if errors.As(err, &Err) {
			return Err.HTTP, Err.Code, Err.Error()
		}
	*/
	
    switch typed := err.(type) {
    case *Errno:
        return typed.HTTP, typed.Code, typed.Message
    default:
    }

    // 默认返回未知错误码和错误信息. 该错误代表服务端出错
    return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}