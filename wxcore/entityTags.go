package wxcore

// EntityTag is used to represent what properties an entity has.
type EntityTag int

// A property tag to add to an entity.
const (
	Lrig EntityTag = iota
	Signi
	Arts
	Key
	Spell
	Resona

	Blue
	White
	Red
	Green
	Black
	Colorless

	Hand
	MainDeck
	Trash
	LrigDeck
	LrigTrash
	LrigZone
	Field
	Ener

	Player1
	Player2

	Cost
	Power
	Limit
	Level
	Class

	Stand
	Rest

	Attacking
	Atacked

	UnableToAttack

	Targetable
	Untargetable

	Health

	CurrentTurn
	TurnsLeft
	NextTurn
)
