// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slack-application/ent/workspace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WorkspaceCreate is the builder for creating a Workspace entity.
type WorkspaceCreate struct {
	config
	mutation *WorkspaceMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WorkspaceCreate) SetName(s string) *WorkspaceCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkspaceCreate) SetDescription(s string) *WorkspaceCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetCreatedBy sets the "created_by" field.
func (wc *WorkspaceCreate) SetCreatedBy(s string) *WorkspaceCreate {
	wc.mutation.SetCreatedBy(s)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkspaceCreate) SetCreatedAt(t time.Time) *WorkspaceCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// Mutation returns the WorkspaceMutation object of the builder.
func (wc *WorkspaceCreate) Mutation() *WorkspaceMutation {
	return wc.mutation
}

// Save creates the Workspace in the database.
func (wc *WorkspaceCreate) Save(ctx context.Context) (*Workspace, error) {
	return withHooks[*Workspace, WorkspaceMutation](ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkspaceCreate) SaveX(ctx context.Context) *Workspace {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkspaceCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkspaceCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkspaceCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Workspace.name"`)}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Workspace.description"`)}
	}
	if _, ok := wc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Workspace.created_by"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workspace.created_at"`)}
	}
	return nil
}

func (wc *WorkspaceCreate) sqlSave(ctx context.Context) (*Workspace, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkspaceCreate) createSpec() (*Workspace, *sqlgraph.CreateSpec) {
	var (
		_node = &Workspace{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workspace.Table, sqlgraph.NewFieldSpec(workspace.FieldID, field.TypeInt))
	)
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(workspace.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workspace.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := wc.mutation.CreatedBy(); ok {
		_spec.SetField(workspace.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workspace.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// WorkspaceCreateBulk is the builder for creating many Workspace entities in bulk.
type WorkspaceCreateBulk struct {
	config
	builders []*WorkspaceCreate
}

// Save creates the Workspace entities in the database.
func (wcb *WorkspaceCreateBulk) Save(ctx context.Context) ([]*Workspace, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workspace, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkspaceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) SaveX(ctx context.Context) []*Workspace {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkspaceCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkspaceCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
