package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func SendRequest(data []byte) Result {
	var r Result
	reqBody := bytes.NewBuffer(data)
	res, err := http.Post(url, content, reqBody)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
