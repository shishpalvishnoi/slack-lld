package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Workspace holds the schema definition for the Workspace entity.
type Workspace struct {
	ent.Schema
}

// Fields of the Workspace.
func (Workspace) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("created_by"),
		field.Time("created_at"),
	}
}

// Edges of the Workspace.
func (Workspace) Edges() []ent.Edge {
	return nil
}
