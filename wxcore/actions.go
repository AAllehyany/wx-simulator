package wxcore

import uuid "github.com/satori/go.uuid"

// ActionType represents different actions
type ActionType int

// Type of action
const (
	Play ActionType = iota
	Move
	EndTurn
)

// Action describes how to modify the gamestate
type Action struct {
	Type         ActionType
	PlayerId     uuid.UUID
	TargetEntity uuid.UUID
	Tag          EntityTag
	Value        int
}

// NewAction creates a new action with the given parameters
func NewAction(t ActionType, pId, target uuid.UUID, tag EntityTag, value int) Action {
	return Action{
		Type:         t,
		PlayerId:     pId,
		TargetEntity: target,
		Tag:          tag,
		Value:        value,
	}
}
