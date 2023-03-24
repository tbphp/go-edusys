package e

var codeDescriptions = map[int]string{
	Ok:               "success",
	Unauthorized:     "认证失败",
	Authorization:    "没有权限",
	NotFound:         "路由错误",
	NotAllowed:       "请求方式错误",
	EmptyData:        "数据不存在",
	ValidationFailed: "字段验证失败",
	Exception:        "业务异常，服务端没有响应。",
}

func GetMsg(code int) string {
	msg, ok := codeDescriptions[code]
	if ok {
		return msg
	}

	return codeDescriptions[Exception]
}
