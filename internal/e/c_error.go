package e

type CError struct {
	code int
	msg  string
}

func New(code int, msg string) *CError {
	return &CError{code: code, msg: msg}
}

func (e *CError) Error() string {
	return e.msg
}

func (e *CError) Code() int {
	return e.code
}

var (
	Unauthorized = func(message string) *CError {
		return New(401, message)
	}
	Authorization = New(403, "没有权限")
	NotFound      = New(404, "路由错误")
	EmptyData     = New(410, "数据不存在")
	Exception     = func(msg string) *CError {
		return New(510, msg)
	}
	Default = New(510, "业务异常，服务端没有响应。")
)
