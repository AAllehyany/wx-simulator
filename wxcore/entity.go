package wxcore

import uuid "github.com/satori/go.uuid"

// Entity is an object with a propery map to define which properties it has
type Entity struct {
	ID uuid.UUID

	// Tags is the map of properties the entity has.
	// Anything greater than 0 indicates the property is active.
	// To remove a property/tag, either set its value to 0 or remove it from the map.
	Tags map[EntityTag]int

	Type  EntityType
	Owner uuid.UUID
}

// NewEntity creates a new entity with a random ID and no Tags
func NewEntity(t EntityType, owner uuid.UUID) Entity {
	return Entity{
		ID:    uuid.NewV4(),
		Tags:  make(map[EntityTag]int),
		Type:  t,
		Owner: owner,
	}
}
