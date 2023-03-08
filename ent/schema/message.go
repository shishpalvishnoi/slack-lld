package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("data"),
		field.String("attachment"),
		field.Int64("user_id"),
		field.Int64("workspace_id"),
		field.Time("created_at"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
