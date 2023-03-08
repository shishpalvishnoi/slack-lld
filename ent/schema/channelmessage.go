package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// ChannelMessage holds the schema definition for the ChannelMessage entity.
type ChannelMessage struct {
	ent.Schema
}

// Fields of the ChannelMessage.
func (ChannelMessage) Fields() []ent.Field {
	return []ent.Field{
		field.Other("postgres_array_col", &pgtype.Int4Array{}).
			SchemaType(map[string]string{
				dialect.Postgres: "integer[]",
			}),
	}
}

// Edges of the ChannelMessage.
func (ChannelMessage) Edges() []ent.Edge {
	return nil
}
