package models

type Response struct {
	Version string      `json:"jsonrpc"`
	Result  string      `json:"result,omitempty"`
	Error   ErrorResp   `json:"error,omitempty"`
	Id      interface{} `json:"id"`
}

func NewResponse(result string, id interface{}) Response {
	return Response{
		Version: Version,
		Result:  result,
		Id:      id,
	}
}

func NewResponseError(errResp *ErrorResp, id interface{}) Response {
	return Response{
		Version: Version,
		Error:   *errResp,
		Id:      id,
	}
}
