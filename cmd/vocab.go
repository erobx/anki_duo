package cmd

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

type VocabState struct {
	Words []Word
	Total int
}

type Word struct {
	Text         string   `json:"text"`
	Translations []string `json:"translations"`
}

func NewVocabState() *VocabState {
	return &VocabState{}
}

// read vocab from Duo api
func (vs *VocabState) ReadFile(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	n := len(data) - 2
	data = data[2:n]
	vs.extractWords(data)
}

// format vocab for Anki import
// word;translation, translation, translation
func (vs *VocabState) ToAnki(name string) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, w := range vs.Words {
		file.WriteString(w.Text + ";")
		for i, t := range w.Translations {
			if i == len(w.Translations)-1 {
				file.WriteString(t + "\n")
			} else {
				file.WriteString(t + ", ")
			}
		}
	}
}

func (vs *VocabState) extractWords(data []byte) {
	objs := strings.Split(string(data), "},")
	for i, obj := range objs {
		check := strings.Contains(obj, "],[")
		if check {
			bef, aft, _ := strings.Cut(obj, "],[")
			aft = aft + "}"
			vs.addWord(bef)
			vs.addWord(aft)
		} else if i == len(objs)-1 {
			vs.addWord(obj)
		} else {
			obj = obj + "}"
			vs.addWord(obj)
		}
	}
}

func (vs *VocabState) addWord(w string) {
	var word Word
	err := json.Unmarshal([]byte(w), &word)
	if err != nil {
		log.Fatal(err)
	}

	vs.Words = append(vs.Words, word)
	vs.Total += 1
}
