package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/AAllehyany/wx-simulator/wxcore"
)

func main() {
	pirulukData, err := ioutil.ReadFile("data/piruluk.json")
	if err != nil {
		println("FAILED TO LOAD FILE")
	}

	var cards []*wxcore.CardData

	if err = json.Unmarshal(pirulukData, &cards); err != nil {
		fmt.Printf("FAILED TO LOAD PIRULUK %v", err)
	}

	game := wxcore.NewGame()

	wxcore.CardToEntityParser(cards, &game)

	wxcore.GetAllCards(&game)

}
