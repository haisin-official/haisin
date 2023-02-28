package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
			StorageKey("uuid").
			Unique(),
		field.Enum("service").
			Values("twitter", "youtube", "fanbox"),
		field.String("url").
			MaxLen(2083).
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now()),
		field.Time("updated_at").
			Default(time.Now()),
	}
}

// Edges of the Url.
func (Url) Edges() []ent.Edge {
	return nil
}
