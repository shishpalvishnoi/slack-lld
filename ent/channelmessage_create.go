// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"slack-application/ent/channelmessage"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jackc/pgtype"
)

// ChannelMessageCreate is the builder for creating a ChannelMessage entity.
type ChannelMessageCreate struct {
	config
	mutation *ChannelMessageMutation
	hooks    []Hook
}

// SetPostgresArrayCol sets the "postgres_array_col" field.
func (cmc *ChannelMessageCreate) SetPostgresArrayCol(pg *pgtype.Int4Array) *ChannelMessageCreate {
	cmc.mutation.SetPostgresArrayCol(pg)
	return cmc
}

// Mutation returns the ChannelMessageMutation object of the builder.
func (cmc *ChannelMessageCreate) Mutation() *ChannelMessageMutation {
	return cmc.mutation
}

// Save creates the ChannelMessage in the database.
func (cmc *ChannelMessageCreate) Save(ctx context.Context) (*ChannelMessage, error) {
	return withHooks[*ChannelMessage, ChannelMessageMutation](ctx, cmc.sqlSave, cmc.mutation, cmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cmc *ChannelMessageCreate) SaveX(ctx context.Context) *ChannelMessage {
	v, err := cmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmc *ChannelMessageCreate) Exec(ctx context.Context) error {
	_, err := cmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmc *ChannelMessageCreate) ExecX(ctx context.Context) {
	if err := cmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cmc *ChannelMessageCreate) check() error {
	if _, ok := cmc.mutation.PostgresArrayCol(); !ok {
		return &ValidationError{Name: "postgres_array_col", err: errors.New(`ent: missing required field "ChannelMessage.postgres_array_col"`)}
	}
	return nil
}

func (cmc *ChannelMessageCreate) sqlSave(ctx context.Context) (*ChannelMessage, error) {
	if err := cmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cmc.mutation.id = &_node.ID
	cmc.mutation.done = true
	return _node, nil
}

func (cmc *ChannelMessageCreate) createSpec() (*ChannelMessage, *sqlgraph.CreateSpec) {
	var (
		_node = &ChannelMessage{config: cmc.config}
		_spec = sqlgraph.NewCreateSpec(channelmessage.Table, sqlgraph.NewFieldSpec(channelmessage.FieldID, field.TypeInt))
	)
	if value, ok := cmc.mutation.PostgresArrayCol(); ok {
		_spec.SetField(channelmessage.FieldPostgresArrayCol, field.TypeOther, value)
		_node.PostgresArrayCol = value
	}
	return _node, _spec
}

// ChannelMessageCreateBulk is the builder for creating many ChannelMessage entities in bulk.
type ChannelMessageCreateBulk struct {
	config
	builders []*ChannelMessageCreate
}

// Save creates the ChannelMessage entities in the database.
func (cmcb *ChannelMessageCreateBulk) Save(ctx context.Context) ([]*ChannelMessage, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cmcb.builders))
	nodes := make([]*ChannelMessage, len(cmcb.builders))
	mutators := make([]Mutator, len(cmcb.builders))
	for i := range cmcb.builders {
		func(i int, root context.Context) {
			builder := cmcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMessageMutation)
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
					_, err = mutators[i+1].Mutate(root, cmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cmcb *ChannelMessageCreateBulk) SaveX(ctx context.Context) []*ChannelMessage {
	v, err := cmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cmcb *ChannelMessageCreateBulk) Exec(ctx context.Context) error {
	_, err := cmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmcb *ChannelMessageCreateBulk) ExecX(ctx context.Context) {
	if err := cmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
