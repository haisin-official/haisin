// Code generated by ent, DO NOT EDIT.

package url

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the url type in the database.
	Label = "url"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "urlid"
	// FieldService holds the string denoting the service field in the database.
	FieldService = "service"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeUserID holds the string denoting the user_id edge name in mutations.
	EdgeUserID = "user_id"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "user_id"
	// Table holds the table name of the url in the database.
	Table = "urls"
	// UserIDTable is the table that holds the user_id relation/edge.
	UserIDTable = "urls"
	// UserIDInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserIDInverseTable = "users"
	// UserIDColumn is the table column denoting the user_id relation/edge.
	UserIDColumn = "user_urls"
)

// Columns holds all SQL columns for url fields.
var Columns = []string{
	FieldID,
	FieldService,
	FieldURL,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "urls"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_urls",
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
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// Service defines the type for the "service" enum field.
type Service string

// Service values.
const (
	ServiceTwitter Service = "twitter"
	ServiceYoutube Service = "youtube"
	ServiceFanbox  Service = "fanbox"
)

func (s Service) String() string {
	return string(s)
}

// ServiceValidator is a validator for the "service" field enum values. It is called by the builders before save.
func ServiceValidator(s Service) error {
	switch s {
	case ServiceTwitter, ServiceYoutube, ServiceFanbox:
		return nil
	default:
		return fmt.Errorf("url: invalid enum value for service field: %q", s)
	}
}
