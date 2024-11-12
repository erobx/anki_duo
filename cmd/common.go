package cmd

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const (
	url     = "http://localhost:8765"
	content = "application/json"
)

func SendRequest(data []byte) []byte {
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

	return body
}
