package main

import (
	"encoding/json"
	"log"
)

type Request struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

type ParamRequest struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
	Params  any    `json:"params"`
}

func NewRequest(action string) Request {
	return Request{
		Action:  action,
		Version: 6,
	}
}

func NewParamRequest(action string, params any) ParamRequest {
	return ParamRequest{
		Action:  action,
		Version: 6,
		Params:  params,
	}
}

func (r Request) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r ParamRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (req Request) Call() Result {
	data, err := req.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return SendRequest(data)
}

func (req ParamRequest) Call() Result {
	data, err := req.Marshal()
	if err != nil {
		log.Fatal(err)
	}
	return SendRequest(data)
}
