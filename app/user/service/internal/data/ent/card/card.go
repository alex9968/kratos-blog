// Code generated by entc, DO NOT EDIT.

package card

import (
	"time"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCardNo holds the string denoting the card_no field in the database.
	FieldCardNo = "card_no"
	// FieldCcv holds the string denoting the ccv field in the database.
	FieldCcv = "ccv"
	// FieldExpires holds the string denoting the expires field in the database.
	FieldExpires = "expires"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "cards"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "owner_id"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCardNo,
	FieldCcv,
	FieldExpires,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"owner_id",
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
