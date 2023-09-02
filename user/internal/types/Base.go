package types

type Base struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewErrorEntityDefault() *Response {
	base := Base{
		Code:    "1001",
		Message: "系统异常，请联系管理员",
	}
	return &Response{
		base,
		nil,
	}
}

func NewErrorEntity(code string, message string) *Response {
	base := Base{
		Code:    code,
		Message: message,
	}
	return &Response{
		base,
		nil,
	}
}

func NewSuccessEntity(data interface{}) *Response {
	base := Base{
		Code:    "000000",
		Message: "success",
	}
	return &Response{
		base,
		data,
	}
}
