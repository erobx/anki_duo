package main

func GetDeckNames() Result {
	req := NewRequest("deckNames")
	return req.Call()
}

func FindCards() Result {
	params := map[string]string{
		"query": "deck:current",
	}
	req := NewParamRequest("findCards", params)
	res := req.Call()
	for x, i := range res.Res {
		i = int(i.(float64))
		res.Res[x] = i
	}
	return res
}
