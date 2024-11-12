package main

import (
	"fmt"

	"github.com/erobx/anki_duo/cmd"
)

func main() {
	var username, jwt string
	fmt.Print("Enter your Duolingo username: ")
	fmt.Scan(&username)

	fmt.Print("Enter your JWT from Duolingo: ")
	fmt.Scan(&jwt)

	cmdr := cmd.NewCmdr(username, jwt)
	cmdr.ImportVocab("vocab.json")
}
