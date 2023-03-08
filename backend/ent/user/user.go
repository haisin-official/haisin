// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldGa holds the string denoting the ga field in the database.
	FieldGa = "ga"
	// EdgeUUID holds the string denoting the uuid edge name in mutations.
	EdgeUUID = "uuid"
	// Table holds the table name of the user in the database.
	Table = "users"
	// UUIDTable is the table that holds the uuid relation/edge.
	UUIDTable = "urls"
	// UUIDInverseTable is the table name for the Url entity.
	// It exists in this package in order to avoid circular dependency with the "url" package.
	UUIDInverseTable = "urls"
	// UUIDColumn is the table column denoting the uuid relation/edge.
	UUIDColumn = "user_uuid"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldSlug,
	FieldGa,
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

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	SlugValidator func(string) error
	// GaValidator is a validator for the "ga" field. It is called by the builders before save.
	GaValidator func(string) error
)
