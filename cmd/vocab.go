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

func (vs *VocabState) ReadFile(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	n := len(data) - 2
	data = data[2:n]
	vs.extractWords(data)
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
