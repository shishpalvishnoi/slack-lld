// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"slack-application/ent/predicate"
	"slack-application/ent/userchannel"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserChannelDelete is the builder for deleting a UserChannel entity.
type UserChannelDelete struct {
	config
	hooks    []Hook
	mutation *UserChannelMutation
}

// Where appends a list predicates to the UserChannelDelete builder.
func (ucd *UserChannelDelete) Where(ps ...predicate.UserChannel) *UserChannelDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserChannelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, UserChannelMutation](ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserChannelDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserChannelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userchannel.Table, sqlgraph.NewFieldSpec(userchannel.FieldID, field.TypeInt))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UserChannelDeleteOne is the builder for deleting a single UserChannel entity.
type UserChannelDeleteOne struct {
	ucd *UserChannelDelete
}

// Where appends a list predicates to the UserChannelDelete builder.
func (ucdo *UserChannelDeleteOne) Where(ps ...predicate.UserChannel) *UserChannelDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UserChannelDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userchannel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserChannelDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
