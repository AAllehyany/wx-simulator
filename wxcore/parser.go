package wxcore

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func CardToEntityParser(cards []*CardData, gameState *GameState, owner uuid.UUID) {

	for _, v := range cards {
		eid := CreateEntity(gameState, Card, owner)
		PraseCard(gameState, eid, v)
	}
}

func PraseCard(gameState *GameState, eid uuid.UUID, card *CardData) {

	cardType := getCardType(card.Type)
	cardZone := getCardZone(card.Type)
	UpdateEntity(gameState, eid, cardType, 1)
	UpdateEntity(gameState, eid, Power, card.Power)
	UpdateEntity(gameState, eid, Cost, card.Cost)
	UpdateEntity(gameState, eid, Level, card.Level)
	UpdateEntity(gameState, eid, Limit, card.Limit)
	UpdateEntity(gameState, eid, cardZone, 1)

}

func getCardType(t string) EntityTag {
	switch ct := strings.ToLower(t); ct {
	case "lrig":
		return Lrig
	case "signi":
		return Signi
	case "arts":
		return Arts
	case "spell":
		return Spell
	case "key":
		return Key
	case "resona":
		return Resona
	}

	return Signi
}

func getCardZone(t string) EntityTag {
	switch ct := strings.ToLower(t); ct {
	case "lrig", "arts":
		return LrigDeck
	default:
		return MainDeck
	}
}
