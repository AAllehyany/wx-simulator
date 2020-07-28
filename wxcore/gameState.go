package wxcore

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GameState represents the current state of the game.
// It is responsible of keeping track of creation, modification and updating entities.
type GameState struct {
	Entities map[uuid.UUID]Entity
}

func NewGame() GameState {
	return GameState{
		Entities: make(map[uuid.UUID]Entity),
	}
}

// CreateEntity creats a new entity and adds it to the game.
func CreateEntity(gameState *GameState, t EntityType, owner uuid.UUID) uuid.UUID {

	entity := NewEntity(t, owner)

	gameState.Entities[entity.ID] = entity

	return entity.ID
}

// UpdateEntity updates an entity tag with the provided value, or adds a new tag with the given value
func UpdateEntity(gameState *GameState, entityId uuid.UUID, tag EntityTag, value int) {

	if e, ok := gameState.Entities[entityId]; ok {
		e.Tags[tag] = value
	}
}

// RemoveEntity removes an entity from the game if it exists
func RemoveEntity(gameState *GameState, entityId uuid.UUID) {
	delete(gameState.Entities, entityId)
}

func GetAllCards(gameState *GameState) []*Entity {
	var cards []*Entity

	for _, element := range gameState.Entities {
		if element.Type == Card {
			fmt.Printf("%+v", element)
		}
	}

	return cards
}
