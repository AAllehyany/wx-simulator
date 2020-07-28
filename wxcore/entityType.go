package wxcore

// EntityType represents what the entity is for different operations.
type EntityType int

// Represents what the entity is
const (
	Card EntityType = iota
	Player
	Zone
)
