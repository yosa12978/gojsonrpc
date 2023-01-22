package models

import (
	"encoding/json"
	"reflect"
)

type Request struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Id      interface{} `json:"id"`
}

type Batch []Request

func ParseRequest(b []byte) (*Request, *Batch, error) {
	var req interface{}
	err := json.Unmarshal(b, req)
	if err != nil {
		return nil, nil, ParseError(err.Error())
	}
	if string(b[0]) == "[" && string(b[len(b)-1]) == "]" {
		return nil, req.(*Batch), nil
	}
	return req.(*Request), nil, nil
}

func (r *Request) ProcRequest(obj interface{}) {
	var vals []reflect.Value
	reflect.ValueOf(obj).MethodByName(r.Method).Call(vals)
}

func (b *Batch) ProcBatch(obj interface{}) {
	for _, v := range *b {
		v.ProcRequest(obj)
	}
}
