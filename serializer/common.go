package serializer

// Response 基础序列化器
type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
}
