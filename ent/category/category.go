// Code generated by ent, DO NOT EDIT.

package category

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// AdminTable is the table that holds the admin relation/edge.
	AdminTable = "categories"
	// AdminInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminInverseTable = "users"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "admin_id"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "categories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"admin_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
