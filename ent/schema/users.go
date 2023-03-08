package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("phone_number"),
		field.String("user_type"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Int("version_id"),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
