// Code generated by entc, DO NOT EDIT.

package entry

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the entry type in the database.
	Label = "entry"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the entry in the database.
	Table = "entries"
	// AccountTable is the table the holds the account relation/edge.
	AccountTable = "entries"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_entries"
)

// Columns holds all SQL columns for entry fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldAmount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "entries"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"account_entries",
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
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)