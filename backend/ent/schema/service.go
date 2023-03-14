package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
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

// Mixin of the Service.
func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	// return nil
	return []ent.Edge{
		edge.From("user_id", User.Type).
			Unique().
			Required().
			Ref("uuid"),
	}
}
