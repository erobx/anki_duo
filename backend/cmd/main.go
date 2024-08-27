package main

import (
	"fmt"
)

// Globals for call to anki application
const (
	url     = "http://localhost:8765"
	content = "application/json"
)

func main() {
	r := GetDeckNames()
	//r := FindCards()
	fmt.Printf("%+v\n", r)
}
