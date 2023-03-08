package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChannelMessage holds the schema definition for the ChannelMessage entity.
type ChannelMessage struct {
	ent.Schema
}

// Fields of the ChannelMessage.
func (ChannelMessage) Fields() []ent.Field {
	return []ent.Field {
		field.String("messageIds"),
		field.Time("created_at"),
	}
}

// Edges of the ChannelMessage.
func (ChannelMessage) Edges() []ent.Edge {
	return nil
}
