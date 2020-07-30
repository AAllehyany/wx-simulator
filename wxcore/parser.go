package wxcore

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func CardToEntityParser(cards []*CardData, gameState *GameState, owner uuid.UUID) {

	for _, v := range cards {
		eid := gameState.CreateEntity(Card, owner)
		PraseCard(gameState, eid, v)
	}
}

func PraseCard(gameState *GameState, eid uuid.UUID, card *CardData) {

	cardType := getCardType(card.Type)
	cardZone := getCardZone(card.Type)
	gameState.UpdateEntity(eid, cardType, 1)
	gameState.UpdateEntity(eid, Power, card.Power)
	gameState.UpdateEntity(eid, Cost, card.Cost)
	gameState.UpdateEntity(eid, Level, card.Level)
	gameState.UpdateEntity(eid, Limit, card.Limit)
	gameState.UpdateEntity(eid, Place, int(cardZone))
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

func getCardZone(t string) Zone {
	switch ct := strings.ToLower(t); ct {
	case "lrig", "arts":
		return LrigDeck
	default:
		return MainDeck
	}
}
