// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/haisin-official/haisin/ent/url"
	"github.com/haisin-official/haisin/ent/user"
)

// Url is the model entity for the Url schema.
type Url struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Service holds the value of the "service" field.
	Service url.Service `json:"service,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UrlQuery when eager-loading is set.
	Edges     UrlEdges `json:"edges"`
	user_urls *uuid.UUID
}

// UrlEdges holds the relations/edges for other nodes in the graph.
type UrlEdges struct {
	// UserID holds the value of the user_id edge.
	UserID *User `json:"user_id,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserIDOrErr returns the UserID value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UrlEdges) UserIDOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.UserID == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UserID, nil
	}
	return nil, &NotLoadedError{edge: "user_id"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Url) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case url.FieldService, url.FieldURL:
			values[i] = new(sql.NullString)
		case url.FieldCreatedAt, url.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case url.FieldID:
			values[i] = new(uuid.UUID)
		case url.ForeignKeys[0]: // user_urls
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
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
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case url.FieldService:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field service", values[i])
			} else if value.Valid {
				u.Service = url.Service(value.String)
			}
		case url.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				u.URL = value.String
			}
		case url.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case url.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case url.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_urls", values[i])
			} else if value.Valid {
				u.user_urls = new(uuid.UUID)
				*u.user_urls = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUserID queries the "user_id" edge of the Url entity.
func (u *Url) QueryUserID() *UserQuery {
	return NewURLClient(u.config).QueryUserID(u)
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
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("service=")
	builder.WriteString(fmt.Sprintf("%v", u.Service))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(u.URL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Urls is a parsable slice of Url.
type Urls []*Url
