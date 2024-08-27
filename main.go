package main

import (
	"fmt"

	"github.com/erobx/anki_duo/backend/cmd"
)

func main() {
	anki := cmd.NewAnki()
	fmt.Println("Importing file...")
	anki.ImportFile("/home/erob/Documents/projects/anki_duo/backend/scripts/import.txt")
}
