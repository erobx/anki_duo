package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/erobx/anki_duo/cmd"
	"github.com/erobx/golingo/duolingo"
	"github.com/joho/godotenv"
)

const (
	baseUrl = "https://www.duolingo.com"
	abbr    = "es"
)

type Cmdr struct {
}

func NewCmdr() Cmdr {
	return Cmdr{}
}

func main() {
	var jwt string
	fmt.Print("Enter your JWT from Duolingo: ")
	fmt.Scan(&jwt)

	duo := createConn(jwt)
	vocab := duo.GetVocab(abbr)
	dumpVocab(vocab, "data/vocab.json")

	vocabState := cmd.NewVocabState()
	vocabState.ReadFile("data/vocab.json")
	vocabState.ToAnki("data/import.txt")

	cmdr := NewCmdr()
	cmdr.runAnki()
}

func (c *Cmdr) runAnki() {
	fmt.Println("Importing to Anki")
	anki := cmd.NewAnki()
	anki.ImportFile()
}

func createConn(token string) *duolingo.Duolingo {
	err := godotenv.Load()
	if err != nil {
		os.Exit(1)
	}

	username := os.Getenv("DUO_USERNAME")

	duo := duolingo.NewDuolingo(username, token, baseUrl, abbr)
	return duo
}

func dumpVocab(vocab []interface{}, file string) {
	data, err := json.Marshal(vocab)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(file, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
