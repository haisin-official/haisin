// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/haisin-official/haisin/ent/url"
)

// Url is the model entity for the Url schema.
type Url struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Url) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Url", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Url fields.
func (u *Url) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case url.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Url.
// Note that you need to call Url.Unwrap() before calling this method if this Url
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Url) Update() *URLUpdateOne {
	return NewURLClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Url entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Url) Unwrap() *Url {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Url is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Url) String() string {
	var builder strings.Builder
	builder.WriteString("Url(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Urls is a parsable slice of Url.
type Urls []*Url
