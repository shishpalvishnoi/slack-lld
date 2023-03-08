// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slack-application/ent/predicate"
	"slack-application/ent/users"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersUpdate is the builder for updating Users entities.
type UsersUpdate struct {
	config
	hooks    []Hook
	mutation *UsersMutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uu *UsersUpdate) Where(ps ...predicate.Users) *UsersUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UsersUpdate) SetName(s string) *UsersUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UsersUpdate) SetEmail(s string) *UsersUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPhoneNumber sets the "phone_number" field.
func (uu *UsersUpdate) SetPhoneNumber(s string) *UsersUpdate {
	uu.mutation.SetPhoneNumber(s)
	return uu
}

// SetUserType sets the "user_type" field.
func (uu *UsersUpdate) SetUserType(s string) *UsersUpdate {
	uu.mutation.SetUserType(s)
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UsersUpdate) SetCreatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UsersUpdate) SetUpdatedAt(t time.Time) *UsersUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetVersionID sets the "version_id" field.
func (uu *UsersUpdate) SetVersionID(i int) *UsersUpdate {
	uu.mutation.ResetVersionID()
	uu.mutation.SetVersionID(i)
	return uu
}

// AddVersionID adds i to the "version_id" field.
func (uu *UsersUpdate) AddVersionID(i int) *UsersUpdate {
	uu.mutation.AddVersionID(i)
	return uu
}

// Mutation returns the UsersMutation object of the builder.
func (uu *UsersUpdate) Mutation() *UsersMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UsersUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UsersMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UsersUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UsersUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UsersUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UsersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(users.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PhoneNumber(); ok {
		_spec.SetField(users.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := uu.mutation.UserType(); ok {
		_spec.SetField(users.FieldUserType, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.VersionID(); ok {
		_spec.SetField(users.FieldVersionID, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedVersionID(); ok {
		_spec.AddField(users.FieldVersionID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UsersUpdateOne is the builder for updating a single Users entity.
type UsersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UsersMutation
}

// SetName sets the "name" field.
func (uuo *UsersUpdateOne) SetName(s string) *UsersUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UsersUpdateOne) SetEmail(s string) *UsersUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPhoneNumber sets the "phone_number" field.
func (uuo *UsersUpdateOne) SetPhoneNumber(s string) *UsersUpdateOne {
	uuo.mutation.SetPhoneNumber(s)
	return uuo
}

// SetUserType sets the "user_type" field.
func (uuo *UsersUpdateOne) SetUserType(s string) *UsersUpdateOne {
	uuo.mutation.SetUserType(s)
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UsersUpdateOne) SetCreatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UsersUpdateOne) SetUpdatedAt(t time.Time) *UsersUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetVersionID sets the "version_id" field.
func (uuo *UsersUpdateOne) SetVersionID(i int) *UsersUpdateOne {
	uuo.mutation.ResetVersionID()
	uuo.mutation.SetVersionID(i)
	return uuo
}

// AddVersionID adds i to the "version_id" field.
func (uuo *UsersUpdateOne) AddVersionID(i int) *UsersUpdateOne {
	uuo.mutation.AddVersionID(i)
	return uuo
}

// Mutation returns the UsersMutation object of the builder.
func (uuo *UsersUpdateOne) Mutation() *UsersMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UsersUpdate builder.
func (uuo *UsersUpdateOne) Where(ps ...predicate.Users) *UsersUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UsersUpdateOne) Select(field string, fields ...string) *UsersUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Users entity.
func (uuo *UsersUpdateOne) Save(ctx context.Context) (*Users, error) {
	return withHooks[*Users, UsersMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UsersUpdateOne) SaveX(ctx context.Context) *Users {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UsersUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UsersUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UsersUpdateOne) sqlSave(ctx context.Context) (_node *Users, err error) {
	_spec := sqlgraph.NewUpdateSpec(users.Table, users.Columns, sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Users.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, users.FieldID)
		for _, f := range fields {
			if !users.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != users.FieldID {
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
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(users.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(users.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PhoneNumber(); ok {
		_spec.SetField(users.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := uuo.mutation.UserType(); ok {
		_spec.SetField(users.FieldUserType, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(users.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(users.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.VersionID(); ok {
		_spec.SetField(users.FieldVersionID, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedVersionID(); ok {
		_spec.AddField(users.FieldVersionID, field.TypeInt, value)
	}
	_node = &Users{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{users.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}