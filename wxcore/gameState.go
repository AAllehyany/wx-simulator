package wxcore

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// GameState represents the current state of the game.
// It is responsible of keeping track of creation, modification and updating entities.
type GameState struct {
	Entities  map[uuid.UUID]Entity
	MainDecks map[uuid.UUID]CardCollection
	Hands     map[uuid.UUID]CardCollection
}

func NewGame() GameState {
	return GameState{
		Entities:  make(map[uuid.UUID]Entity),
		MainDecks: make(map[uuid.UUID]CardCollection),
		Hands:     make(map[uuid.UUID]CardCollection),
	}
}

// CreateEntity creats a new entity and adds it to the game.
func (g *GameState) CreateEntity(t EntityType, owner uuid.UUID) uuid.UUID {

	entity := NewEntity(t, owner)

	g.Entities[entity.ID] = entity

	return entity.ID
}

// UpdateEntity updates an entity tag with the provided value, or adds a new tag with the given value
func (g *GameState) UpdateEntity(entityId uuid.UUID, tag EntityTag, value int) {

	if e, ok := g.Entities[entityId]; ok {
		e.Tags[tag] = value
	}
}

// RemoveEntity removes an entity from the game if it exists
func (g *GameState) RemoveEntity(entityId uuid.UUID) {
	delete(g.Entities, entityId)
}

func (g *GameState) GetAllCards() []Entity {
	var cards []Entity

	for _, element := range g.Entities {
		if element.Type == Card {
			//fmt.Printf("%+v", element)
			cards = append(cards, element)
		}
	}

	return cards
}

func (g *GameState) PrintCardsInZone(zone Zone) {

	for _, element := range g.Entities {
		fmt.Println("found element")
		if e, ok := element.Tags[Place]; ok && e == int(zone) {
			fmt.Printf("%+v", element)
		}
	}
}

func (g *GameState) CreatePlayerEntity() uuid.UUID {
	p := NewEntity(Player, uuid.NewV4())
	g.Entities[p.ID] = p
	g.MainDecks[p.ID] = NewCardCollection(HiddenZone)
	g.Hands[p.ID] = NewCardCollection(PrivateZone)
	return p.ID
}

func (g *GameState) ShuffleDeck(id uuid.UUID) {

	if deck, ok := g.MainDecks[id]; ok {
		deck.Shuffle()
	}

}

func (g *GameState) Draw(id uuid.UUID, amount int) {}

func ApplyAction(action Action, game *GameState) {
	if entity, ok := game.Entities[action.TargetEntity]; ok {

		if entity.Owner != action.PlayerId {
			return
		}

		entity.Tags[action.Tag] = action.Value
	}
}
