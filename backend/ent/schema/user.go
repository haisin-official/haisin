package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StorageKey("uuid").
			Unique(),
		field.String("email").
			MaxLen(100).
			NotEmpty().
			Unique(),
		field.String("slug").
			MaxLen(30).
			NotEmpty().
			Unique(),
		field.String("ga").
			MaxLen(100).
			Optional(),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
