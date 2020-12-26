package schema

import "github.com/facebookincubator/ent"

// Reserveroom holds the schema definition for the Reserveroom entity.
type Reserveroom struct {
	ent.Schema
}

// Fields of the Reserveroom.
func (Reserveroom) Fields() []ent.Field {
	return nil
}

// Edges of the Reserveroom.
func (Reserveroom) Edges() []ent.Edge {
	return nil
}
