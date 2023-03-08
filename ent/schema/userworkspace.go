package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UserWorkspace holds the schema definition for the UserWorkspace entity.
type UserWorkspace struct {
	ent.Schema
}

// Fields of the UserWorkspace.
func (UserWorkspace) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("workspace_id"),
		field.String("role_type"),
	}
}

// Edges of the UserWorkspace.
func (UserWorkspace) Edges() []ent.Edge {
	return nil
}
