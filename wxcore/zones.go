package wxcore

type Zone int

const (
	Hand Zone = iota
	MainDeck
	Trash
	LrigDeck
	LrigTrash
	LrigZone
	Field
	Ener
)
