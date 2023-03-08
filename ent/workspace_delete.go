// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"slack-application/ent/predicate"
	"slack-application/ent/workspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceDelete is the builder for deleting a Workspace entity.
type WorkspaceDelete struct {
	config
	hooks    []Hook
	mutation *WorkspaceMutation
}

// Where appends a list predicates to the WorkspaceDelete builder.
func (wd *WorkspaceDelete) Where(ps ...predicate.Workspace) *WorkspaceDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WorkspaceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, WorkspaceMutation](ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WorkspaceDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WorkspaceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workspace.Table, sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeInt))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WorkspaceDeleteOne is the builder for deleting a single Workspace entity.
type WorkspaceDeleteOne struct {
	wd *WorkspaceDelete
}

// Where appends a list predicates to the WorkspaceDelete builder.
func (wdo *WorkspaceDeleteOne) Where(ps ...predicate.Workspace) *WorkspaceDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WorkspaceDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workspace.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WorkspaceDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
