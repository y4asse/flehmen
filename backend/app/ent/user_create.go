// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flehmen-api/ent/mbti"
	"flehmen-api/ent/specialevent"
	"flehmen-api/ent/user"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetWeight sets the "weight" field.
func (uc *UserCreate) SetWeight(f float64) *UserCreate {
	uc.mutation.SetWeight(f)
	return uc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uc *UserCreate) SetNillableWeight(f *float64) *UserCreate {
	if f != nil {
		uc.SetWeight(*f)
	}
	return uc
}

// SetHeight sets the "height" field.
func (uc *UserCreate) SetHeight(f float64) *UserCreate {
	uc.mutation.SetHeight(f)
	return uc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uc *UserCreate) SetNillableHeight(f *float64) *UserCreate {
	if f != nil {
		uc.SetHeight(*f)
	}
	return uc
}

// SetClerkID sets the "clerk_id" field.
func (uc *UserCreate) SetClerkID(s string) *UserCreate {
	uc.mutation.SetClerkID(s)
	return uc
}

// SetIsMale sets the "is_male" field.
func (uc *UserCreate) SetIsMale(b bool) *UserCreate {
	uc.mutation.SetIsMale(b)
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetMbtiID sets the "mbti" edge to the Mbti entity by ID.
func (uc *UserCreate) SetMbtiID(id int) *UserCreate {
	uc.mutation.SetMbtiID(id)
	return uc
}

// SetNillableMbtiID sets the "mbti" edge to the Mbti entity by ID if the given value is not nil.
func (uc *UserCreate) SetNillableMbtiID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetMbtiID(*id)
	}
	return uc
}

// SetMbti sets the "mbti" edge to the Mbti entity.
func (uc *UserCreate) SetMbti(m *Mbti) *UserCreate {
	return uc.SetMbtiID(m.ID)
}

// AddSpecialEventIDs adds the "special_events" edge to the SpecialEvent entity by IDs.
func (uc *UserCreate) AddSpecialEventIDs(ids ...int) *UserCreate {
	uc.mutation.AddSpecialEventIDs(ids...)
	return uc
}

// AddSpecialEvents adds the "special_events" edges to the SpecialEvent entity.
func (uc *UserCreate) AddSpecialEvents(s ...*SpecialEvent) *UserCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddSpecialEventIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.ClerkID(); !ok {
		return &ValidationError{Name: "clerk_id", err: errors.New(`ent: missing required field "User.clerk_id"`)}
	}
	if _, ok := uc.mutation.IsMale(); !ok {
		return &ValidationError{Name: "is_male", err: errors.New(`ent: missing required field "User.is_male"`)}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
		_node.Weight = value
	}
	if value, ok := uc.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
		_node.Height = value
	}
	if value, ok := uc.mutation.ClerkID(); ok {
		_spec.SetField(user.FieldClerkID, field.TypeString, value)
		_node.ClerkID = value
	}
	if value, ok := uc.mutation.IsMale(); ok {
		_spec.SetField(user.FieldIsMale, field.TypeBool, value)
		_node.IsMale = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := uc.mutation.MbtiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   user.MbtiTable,
			Columns: []string{user.MbtiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mbti.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_mbti = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.SpecialEventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SpecialEventsTable,
			Columns: []string{user.SpecialEventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(specialevent.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}