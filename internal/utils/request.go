package utils

import (
	"encoding/json"
	"log"
)

type Request interface {
	Send() []byte
}

type BasicReq struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

type ParamReq struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params"`
}

func NewRequest(action string) BasicReq {
	return BasicReq{
		Action:  action,
		Version: 6,
	}
}

func NewParamRequest(action string, params any) ParamReq {
	return ParamReq{
		Action:  action,
		Version: 6,
		Params:  params,
	}
}

func (req BasicReq) Send() []byte {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	return SendRequest(data)
}

func (req ParamReq) Send() []byte {
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	return SendRequest(data)
}
