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

	player1 := wxcore.CreatePlayerEntity(&game)
	player2 := wxcore.CreatePlayerEntity(&game)

	wxcore.CardToEntityParser(cards, &game, player1)
	wxcore.CardToEntityParser(cards, &game, player2)

	cc := wxcore.GetAllCards(&game)
	fmt.Printf("%+v", cc)

}
