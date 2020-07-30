package wxcore

import (
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

type CardCollection struct {
	Cards      []uuid.UUID
	Visibility ZoneProperty
}

type ZoneProperty int

const (
	PublicZone ZoneProperty = iota
	PrivateZone
	HiddenZone
)

func (cc *CardCollection) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cc.Cards), func(i, j int) { cc.Cards[i], cc.Cards[j] = cc.Cards[j], cc.Cards[i] })
}

func (cc *CardCollection) AddTop(id uuid.UUID) {
	cc.Cards = append([]uuid.UUID{id}, cc.Cards...)
}

func (cc *CardCollection) AddBottom(id uuid.UUID) {
	cc.Cards = append(cc.Cards, id)
}

func NewCardCollection(visibility ZoneProperty) CardCollection {
	return CardCollection{
		Cards:      []uuid.UUID{},
		Visibility: visibility,
	}
}
