package schema

import "entgo.io/ent"

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return nil
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return nil
}
