package serializer

// Response 基础序列化器
type Response struct {
	Code errorCode   `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// ErrorResponse 返回正确响应
func OkResponse(data interface{}) *Response {
	return &Response{
		Code: CodeOk,
		Msg:  GetErrorMsg(CodeOk),
		Data: data,
	}
}

// ErrorResponse 返回错误响应
func ErrorResponse(code errorCode) *Response {
	return &Response{
		Code: code,
		Msg:  GetErrorMsg(code),
	}
}
