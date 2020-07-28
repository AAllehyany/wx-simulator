package wxcore

import uuid "github.com/satori/go.uuid"

type CardData struct {
	Name  string
	Type  string
	Color int
	Limit int
	Level int
	Cost  int
	Lrig  string
	Power int
	Class string
	ID    uuid.UUID
}
