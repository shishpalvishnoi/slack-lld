// Code generated by ent, DO NOT EDIT.

package channelmessage

const (
	// Label holds the string label denoting the channelmessage type in the database.
	Label = "channel_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessageIds holds the string denoting the messageids field in the database.
	FieldMessageIds = "message_ids"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the channelmessage in the database.
	Table = "channel_messages"
)

// Columns holds all SQL columns for channelmessage fields.
var Columns = []string{
	FieldID,
	FieldMessageIds,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
