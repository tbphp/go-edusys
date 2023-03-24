package response

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResponse(code int, msg string, data any) *Response {
	return &Response{code, msg, data}
}
