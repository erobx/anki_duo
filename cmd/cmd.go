package cmd

import (
	"github.com/erobx/anki_duo/internal/anki"
	"github.com/erobx/anki_duo/internal/vocab"
	"github.com/erobx/golingo/duolingo"
)

const (
	baseUrl = "https://www.duolingo.com"
	abbr    = "es"
)

type Cmdr struct {
	duo  *duolingo.Duolingo
	vs   *vocab.VocabState
	anki anki.Anki
}

func NewCmdr(username, jwt string) Cmdr {
	return Cmdr{
		duo:  createConn(username, jwt),
		vs:   vocab.NewVocabState(),
		anki: anki.NewAnki(),
	}
}

func (c *Cmdr) ImportVocab(name string) {
	name = "data/" + name
	vocab := c.duo.GetVocab(abbr)
	c.vs.DumpVocab(vocab, name)
	c.vs.ReadFile(name)
	c.vs.ToAnki("data/import.txt")
	c.anki.ImportFile()
}

func createConn(username, token string) *duolingo.Duolingo {
	duo := duolingo.NewDuolingo(username, token, baseUrl, abbr)
	return duo
}
