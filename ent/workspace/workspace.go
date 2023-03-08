// Code generated by ent, DO NOT EDIT.

package workspace

const (
	// Label holds the string label denoting the workspace type in the database.
	Label = "workspace"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the workspace in the database.
	Table = "workspaces"
)

// Columns holds all SQL columns for workspace fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldCreatedBy,
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
