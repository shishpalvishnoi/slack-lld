package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserChannel holds the schema definition for the UserChannel entity.
type UserChannel struct {
	ent.Schema
}

// Fields of the UserChannel.
func (UserChannel) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("channel_id"),
		field.String("role_type"),
	}
}

// Edges of the UserChannel.
func (UserChannel) Edges() []ent.Edge {
	return nil
}
