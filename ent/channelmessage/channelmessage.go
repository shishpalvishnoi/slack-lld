// Code generated by ent, DO NOT EDIT.

package channelmessage

const (
	// Label holds the string label denoting the channelmessage type in the database.
	Label = "channel_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPostgresArrayCol holds the string denoting the postgres_array_col field in the database.
	FieldPostgresArrayCol = "postgres_array_col"
	// Table holds the table name of the channelmessage in the database.
	Table = "channel_messages"
)

// Columns holds all SQL columns for channelmessage fields.
var Columns = []string{
	FieldID,
	FieldPostgresArrayCol,
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