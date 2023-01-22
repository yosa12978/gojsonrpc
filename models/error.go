package models

type ErrorResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (e ErrorResp) Error() string {
	return e.Message
}

func NewErrorResp(code int, message string, data ...interface{}) *ErrorResp {
	return &ErrorResp{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ParseError(data ...interface{}) *ErrorResp {
	return NewErrorResp(-32700, "Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.", data...)
}

func InvalidRequest(data ...interface{}) *ErrorResp {
	return NewErrorResp(-32600, "The JSON sent is not a valid Request object.", data...)
}

func InvalidParams(data ...interface{}) *ErrorResp {
	return NewErrorResp(-32602, "Invalid method parameter(s).", data...)
}

func MethodNotFound(data ...interface{}) *ErrorResp {
	return NewErrorResp(-32601, "The method does not exist / is not available.", data...)
}

func InternalError(data ...interface{}) *ErrorResp {
	return NewErrorResp(-32603, "Internal JSON-RPC error.", data...)
}

func ServerError(code int, data ...interface{}) *ErrorResp {
	if code > -32000 || code < -32099 {
		code = -32000
	}
	return NewErrorResp(code, "Server error.", data...)
}
