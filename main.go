package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/erobx/anki_duo/cmd"
	"github.com/erobx/golingo/duolingo"
	"github.com/joho/godotenv"
)

const (
	baseUrl = "https://www.duolingo.com"
	abbr    = "es"
)

type Cmdr struct {
	cmd *exec.Cmd
}

func NewCmdr(args []string) Cmdr {
	return Cmdr{}
}

// go run . JWT
func main() {
	//var jwt string
	//fmt.Print("Enter your JWT from Duolingo: ")
	//fmt.Scan(&jwt)

	//duo := createConn(jwt)
	//vocab := duo.GetVocab(jwt, abbr)
	playTest()
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
	//token := os.Getenv("DUO_JWT")

	duo := duolingo.NewDuolingo(username, token, baseUrl, abbr)
	return duo
}

func dumpVocab(vocab []interface{}) {
	data, err := json.Marshal(vocab)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("vocab.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func playTest() {
	data, err := os.ReadFile("vocab.json")
	if err != nil {
		log.Fatal(err)
	}

	n := len(data) - 2
	data = data[2:n]
	words := extractWords(data)
	_ = words

	//for _, word := range words[:2] {
	//	fmt.Printf("Text: %s, Translations: %v\n", word.Text, word.Translations)
	//}
}

type Word struct {
	Text         string   `json:"text"`
	Translations []string `json:"translations"`
}

// have to parse data for each matching {} skipping ,

func extractWords(data []byte) []Word {
	content := strings.TrimSpace(string(data))
	objects := strings.Split(content, "},{")
	var words []Word

	for i, obj := range objects {
		obj = strings.TrimSpace(obj)
		if i == 0 && strings.HasPrefix(obj, "{") {
			obj = obj[1:]
		}
		if i == len(objects)-1 && strings.HasSuffix(obj, "}") {
			obj = obj[:len(obj)-1]
		}

		// Wrap in braces for valid JSON
		obj = fmt.Sprintf("{%s}", obj)

		// Unmarshal the JSON object into a Word struct
		var word Word
		err := json.Unmarshal([]byte(obj), &word)
		if err != nil {
			log.Fatal(err)
		}
		words = append(words, word)
	}

	return words
}
