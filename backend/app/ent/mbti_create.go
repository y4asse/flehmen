// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flehmen-api/ent/mbti"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MbtiCreate is the builder for creating a Mbti entity.
type MbtiCreate struct {
	config
	mutation *MbtiMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (mc *MbtiCreate) SetType(s string) *MbtiCreate {
	mc.mutation.SetType(s)
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MbtiCreate) SetCreatedAt(t time.Time) *MbtiCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MbtiCreate) SetNillableCreatedAt(t *time.Time) *MbtiCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// Mutation returns the MbtiMutation object of the builder.
func (mc *MbtiCreate) Mutation() *MbtiMutation {
	return mc.mutation
}

// Save creates the Mbti in the database.
func (mc *MbtiCreate) Save(ctx context.Context) (*Mbti, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MbtiCreate) SaveX(ctx context.Context) *Mbti {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MbtiCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MbtiCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MbtiCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := mbti.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MbtiCreate) check() error {
	if _, ok := mc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Mbti.type"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Mbti.created_at"`)}
	}
	return nil
}

func (mc *MbtiCreate) sqlSave(ctx context.Context) (*Mbti, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MbtiCreate) createSpec() (*Mbti, *sqlgraph.CreateSpec) {
	var (
		_node = &Mbti{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(mbti.Table, sqlgraph.NewFieldSpec(mbti.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.GetType(); ok {
		_spec.SetField(mbti.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(mbti.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// MbtiCreateBulk is the builder for creating many Mbti entities in bulk.
type MbtiCreateBulk struct {
	config
	err      error
	builders []*MbtiCreate
}

// Save creates the Mbti entities in the database.
func (mcb *MbtiCreateBulk) Save(ctx context.Context) ([]*Mbti, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Mbti, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MbtiMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MbtiCreateBulk) SaveX(ctx context.Context) []*Mbti {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MbtiCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MbtiCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
