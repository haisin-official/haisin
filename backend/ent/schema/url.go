package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Url holds the schema definition for the Url entity.
type Url struct {
	ent.Schema
}

// Fields of the Url.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Unique(),
		field.Enum("service").
			Values("twitter", "youtube", "fanbox"),
		field.String("url").
			MaxLen(2083).
			NotEmpty(),
	}
}

// Mixin of the Url.
func (Url) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	// return nil
	return []ent.Edge{
		edge.From("user_id", User.Type).
			Unique().
			Required().
			Ref("user_id"),
	}
}
