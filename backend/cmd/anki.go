package cmd

import (
	"encoding/json"
	"log"
)

type Anki struct {
	objMap map[string]json.RawMessage
}

func NewAnki() Anki {
	return Anki{}
}

func (a Anki) CheckError() {
	var e string
	err := json.Unmarshal(a.objMap["error"], &e)
	if err != nil {
		log.Fatal(err)
	}
	if len(e) != 0 {
		log.Fatal(e)
	}
}

/*
{"result": ["Spanish", "test"], "error": null}
*/
func (a Anki) GetDeckNames() []string {
	var names []string

	req := NewRequest("deckNames")
	body := req.Send()

	err := json.Unmarshal(body, &a.objMap)
	if err != nil {
		log.Fatal(err)
	}

	// check error field in response
	a.CheckError()

	err = json.Unmarshal(a.objMap["result"], &names)
	if err != nil {
		log.Fatal(err)
	}
	return names
}

/*
{"result": 1519323742721, "error": null}
*/
func (a Anki) CreateDeck(name string) int {
	var deckId int
	params := map[string]string{
		"deck": name,
	}
	req := NewParamRequest("createDeck", params)
	body := req.Send()

	err := json.Unmarshal(body, &a.objMap)
	if err != nil {
		log.Fatal(err)
	}

	a.CheckError()

	err = json.Unmarshal(a.objMap["result"], &deckId)
	if err != nil {
		log.Fatal(err)
	}
	return deckId
}

/*
{"result": null, "error": null}
*/
func (a Anki) ImportFile(path string) {
	params := map[string]string{
		"path": path,
	}
	req := NewParamRequest("guiImportFile", params)
	body := req.Send()

	err := json.Unmarshal(body, &a.objMap)
	if err != nil {
		log.Fatal(err)
	}

	a.CheckError()

	var r string
	err = json.Unmarshal(a.objMap["result"], &r)
	if err != nil {
		log.Fatal(err)
	}
	if len(r) != 0 {
		log.Fatal(r)
	}
}
