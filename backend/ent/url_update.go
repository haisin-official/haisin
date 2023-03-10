// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/haisin-official/haisin/ent/predicate"
	"github.com/haisin-official/haisin/ent/url"
	"github.com/haisin-official/haisin/ent/user"
)

// URLUpdate is the builder for updating Url entities.
type URLUpdate struct {
	config
	hooks    []Hook
	mutation *URLMutation
}

// Where appends a list predicates to the URLUpdate builder.
func (uu *URLUpdate) Where(ps ...predicate.Url) *URLUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdateTime sets the "update_time" field.
func (uu *URLUpdate) SetUpdateTime(t time.Time) *URLUpdate {
	uu.mutation.SetUpdateTime(t)
	return uu
}

// SetService sets the "service" field.
func (uu *URLUpdate) SetService(u url.Service) *URLUpdate {
	uu.mutation.SetService(u)
	return uu
}

// SetURL sets the "url" field.
func (uu *URLUpdate) SetURL(s string) *URLUpdate {
	uu.mutation.SetURL(s)
	return uu
}

// SetUserIDID sets the "user_id" edge to the User entity by ID.
func (uu *URLUpdate) SetUserIDID(id uuid.UUID) *URLUpdate {
	uu.mutation.SetUserIDID(id)
	return uu
}

// SetUserID sets the "user_id" edge to the User entity.
func (uu *URLUpdate) SetUserID(u *User) *URLUpdate {
	return uu.SetUserIDID(u.ID)
}

// Mutation returns the URLMutation object of the builder.
func (uu *URLUpdate) Mutation() *URLMutation {
	return uu.mutation
}

// ClearUserID clears the "user_id" edge to the User entity.
func (uu *URLUpdate) ClearUserID() *URLUpdate {
	uu.mutation.ClearUserID()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *URLUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks[int, URLMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *URLUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *URLUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *URLUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *URLUpdate) defaults() {
	if _, ok := uu.mutation.UpdateTime(); !ok {
		v := url.UpdateDefaultUpdateTime()
		uu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *URLUpdate) check() error {
	if v, ok := uu.mutation.Service(); ok {
		if err := url.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Url.service": %w`, err)}
		}
	}
	if v, ok := uu.mutation.URL(); ok {
		if err := url.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Url.url": %w`, err)}
		}
	}
	if _, ok := uu.mutation.UserIDID(); uu.mutation.UserIDCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Url.user_id"`)
	}
	return nil
}

func (uu *URLUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(url.Table, url.Columns, sqlgraph.NewFieldSpec(url.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(url.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uu.mutation.Service(); ok {
		_spec.SetField(url.FieldService, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeString, value)
	}
	if uu.mutation.UserIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.UserIDTable,
			Columns: []string{url.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.UserIDTable,
			Columns: []string{url.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// URLUpdateOne is the builder for updating a single Url entity.
type URLUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *URLMutation
}

// SetUpdateTime sets the "update_time" field.
func (uuo *URLUpdateOne) SetUpdateTime(t time.Time) *URLUpdateOne {
	uuo.mutation.SetUpdateTime(t)
	return uuo
}

// SetService sets the "service" field.
func (uuo *URLUpdateOne) SetService(u url.Service) *URLUpdateOne {
	uuo.mutation.SetService(u)
	return uuo
}

// SetURL sets the "url" field.
func (uuo *URLUpdateOne) SetURL(s string) *URLUpdateOne {
	uuo.mutation.SetURL(s)
	return uuo
}

// SetUserIDID sets the "user_id" edge to the User entity by ID.
func (uuo *URLUpdateOne) SetUserIDID(id uuid.UUID) *URLUpdateOne {
	uuo.mutation.SetUserIDID(id)
	return uuo
}

// SetUserID sets the "user_id" edge to the User entity.
func (uuo *URLUpdateOne) SetUserID(u *User) *URLUpdateOne {
	return uuo.SetUserIDID(u.ID)
}

// Mutation returns the URLMutation object of the builder.
func (uuo *URLUpdateOne) Mutation() *URLMutation {
	return uuo.mutation
}

// ClearUserID clears the "user_id" edge to the User entity.
func (uuo *URLUpdateOne) ClearUserID() *URLUpdateOne {
	uuo.mutation.ClearUserID()
	return uuo
}

// Where appends a list predicates to the URLUpdate builder.
func (uuo *URLUpdateOne) Where(ps ...predicate.Url) *URLUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *URLUpdateOne) Select(field string, fields ...string) *URLUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Url entity.
func (uuo *URLUpdateOne) Save(ctx context.Context) (*Url, error) {
	uuo.defaults()
	return withHooks[*Url, URLMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *URLUpdateOne) SaveX(ctx context.Context) *Url {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *URLUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *URLUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *URLUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdateTime(); !ok {
		v := url.UpdateDefaultUpdateTime()
		uuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *URLUpdateOne) check() error {
	if v, ok := uuo.mutation.Service(); ok {
		if err := url.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Url.service": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.URL(); ok {
		if err := url.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Url.url": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.UserIDID(); uuo.mutation.UserIDCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Url.user_id"`)
	}
	return nil
}

func (uuo *URLUpdateOne) sqlSave(ctx context.Context) (_node *Url, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(url.Table, url.Columns, sqlgraph.NewFieldSpec(url.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Url.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, url.FieldID)
		for _, f := range fields {
			if !url.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != url.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(url.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.Service(); ok {
		_spec.SetField(url.FieldService, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeString, value)
	}
	if uuo.mutation.UserIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.UserIDTable,
			Columns: []string{url.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.UserIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.UserIDTable,
			Columns: []string{url.UserIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Url{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
