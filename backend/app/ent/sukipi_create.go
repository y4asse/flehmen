// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flehmen-api/ent/mbti"
	"flehmen-api/ent/sukipi"
	"flehmen-api/ent/tweet"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SukipiCreate is the builder for creating a Sukipi entity.
type SukipiCreate struct {
	config
	mutation *SukipiMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SukipiCreate) SetName(s string) *SukipiCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetWeight sets the "weight" field.
func (sc *SukipiCreate) SetWeight(f float64) *SukipiCreate {
	sc.mutation.SetWeight(f)
	return sc
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableWeight(f *float64) *SukipiCreate {
	if f != nil {
		sc.SetWeight(*f)
	}
	return sc
}

// SetHeight sets the "height" field.
func (sc *SukipiCreate) SetHeight(f float64) *SukipiCreate {
	sc.mutation.SetHeight(f)
	return sc
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableHeight(f *float64) *SukipiCreate {
	if f != nil {
		sc.SetHeight(*f)
	}
	return sc
}

// SetXID sets the "x_id" field.
func (sc *SukipiCreate) SetXID(s string) *SukipiCreate {
	sc.mutation.SetXID(s)
	return sc
}

// SetNillableXID sets the "x_id" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableXID(s *string) *SukipiCreate {
	if s != nil {
		sc.SetXID(*s)
	}
	return sc
}

// SetHobby sets the "hobby" field.
func (sc *SukipiCreate) SetHobby(s string) *SukipiCreate {
	sc.mutation.SetHobby(s)
	return sc
}

// SetNillableHobby sets the "hobby" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableHobby(s *string) *SukipiCreate {
	if s != nil {
		sc.SetHobby(*s)
	}
	return sc
}

// SetBirthday sets the "birthday" field.
func (sc *SukipiCreate) SetBirthday(t time.Time) *SukipiCreate {
	sc.mutation.SetBirthday(t)
	return sc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableBirthday(t *time.Time) *SukipiCreate {
	if t != nil {
		sc.SetBirthday(*t)
	}
	return sc
}

// SetShowsSize sets the "showsSize" field.
func (sc *SukipiCreate) SetShowsSize(s string) *SukipiCreate {
	sc.mutation.SetShowsSize(s)
	return sc
}

// SetNillableShowsSize sets the "showsSize" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableShowsSize(s *string) *SukipiCreate {
	if s != nil {
		sc.SetShowsSize(*s)
	}
	return sc
}

// SetFamily sets the "family" field.
func (sc *SukipiCreate) SetFamily(s string) *SukipiCreate {
	sc.mutation.SetFamily(s)
	return sc
}

// SetNillableFamily sets the "family" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableFamily(s *string) *SukipiCreate {
	if s != nil {
		sc.SetFamily(*s)
	}
	return sc
}

// SetNearlyStation sets the "nearly_station" field.
func (sc *SukipiCreate) SetNearlyStation(s string) *SukipiCreate {
	sc.mutation.SetNearlyStation(s)
	return sc
}

// SetNillableNearlyStation sets the "nearly_station" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableNearlyStation(s *string) *SukipiCreate {
	if s != nil {
		sc.SetNearlyStation(*s)
	}
	return sc
}

// SetLikedAt sets the "liked_at" field.
func (sc *SukipiCreate) SetLikedAt(t time.Time) *SukipiCreate {
	sc.mutation.SetLikedAt(t)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SukipiCreate) SetCreatedAt(t time.Time) *SukipiCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SukipiCreate) SetNillableCreatedAt(t *time.Time) *SukipiCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetMbtiID sets the "mbti" edge to the Mbti entity by ID.
func (sc *SukipiCreate) SetMbtiID(id int) *SukipiCreate {
	sc.mutation.SetMbtiID(id)
	return sc
}

// SetNillableMbtiID sets the "mbti" edge to the Mbti entity by ID if the given value is not nil.
func (sc *SukipiCreate) SetNillableMbtiID(id *int) *SukipiCreate {
	if id != nil {
		sc = sc.SetMbtiID(*id)
	}
	return sc
}

// SetMbti sets the "mbti" edge to the Mbti entity.
func (sc *SukipiCreate) SetMbti(m *Mbti) *SukipiCreate {
	return sc.SetMbtiID(m.ID)
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by IDs.
func (sc *SukipiCreate) AddTweetIDs(ids ...int) *SukipiCreate {
	sc.mutation.AddTweetIDs(ids...)
	return sc
}

// AddTweets adds the "tweets" edges to the Tweet entity.
func (sc *SukipiCreate) AddTweets(t ...*Tweet) *SukipiCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTweetIDs(ids...)
}

// Mutation returns the SukipiMutation object of the builder.
func (sc *SukipiCreate) Mutation() *SukipiMutation {
	return sc.mutation
}

// Save creates the Sukipi in the database.
func (sc *SukipiCreate) Save(ctx context.Context) (*Sukipi, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SukipiCreate) SaveX(ctx context.Context) *Sukipi {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SukipiCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SukipiCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SukipiCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := sukipi.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SukipiCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Sukipi.name"`)}
	}
	if _, ok := sc.mutation.LikedAt(); !ok {
		return &ValidationError{Name: "liked_at", err: errors.New(`ent: missing required field "Sukipi.liked_at"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Sukipi.created_at"`)}
	}
	return nil
}

func (sc *SukipiCreate) sqlSave(ctx context.Context) (*Sukipi, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SukipiCreate) createSpec() (*Sukipi, *sqlgraph.CreateSpec) {
	var (
		_node = &Sukipi{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(sukipi.Table, sqlgraph.NewFieldSpec(sukipi.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(sukipi.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Weight(); ok {
		_spec.SetField(sukipi.FieldWeight, field.TypeFloat64, value)
		_node.Weight = &value
	}
	if value, ok := sc.mutation.Height(); ok {
		_spec.SetField(sukipi.FieldHeight, field.TypeFloat64, value)
		_node.Height = &value
	}
	if value, ok := sc.mutation.XID(); ok {
		_spec.SetField(sukipi.FieldXID, field.TypeString, value)
		_node.XID = &value
	}
	if value, ok := sc.mutation.Hobby(); ok {
		_spec.SetField(sukipi.FieldHobby, field.TypeString, value)
		_node.Hobby = &value
	}
	if value, ok := sc.mutation.Birthday(); ok {
		_spec.SetField(sukipi.FieldBirthday, field.TypeTime, value)
		_node.Birthday = &value
	}
	if value, ok := sc.mutation.ShowsSize(); ok {
		_spec.SetField(sukipi.FieldShowsSize, field.TypeString, value)
		_node.ShowsSize = &value
	}
	if value, ok := sc.mutation.Family(); ok {
		_spec.SetField(sukipi.FieldFamily, field.TypeString, value)
		_node.Family = &value
	}
	if value, ok := sc.mutation.NearlyStation(); ok {
		_spec.SetField(sukipi.FieldNearlyStation, field.TypeString, value)
		_node.NearlyStation = &value
	}
	if value, ok := sc.mutation.LikedAt(); ok {
		_spec.SetField(sukipi.FieldLikedAt, field.TypeTime, value)
		_node.LikedAt = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(sukipi.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.MbtiIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sukipi.MbtiTable,
			Columns: []string{sukipi.MbtiColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(mbti.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.sukipi_mbti = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TweetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sukipi.TweetsTable,
			Columns: []string{sukipi.TweetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SukipiCreateBulk is the builder for creating many Sukipi entities in bulk.
type SukipiCreateBulk struct {
	config
	err      error
	builders []*SukipiCreate
}

// Save creates the Sukipi entities in the database.
func (scb *SukipiCreateBulk) Save(ctx context.Context) ([]*Sukipi, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sukipi, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SukipiMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SukipiCreateBulk) SaveX(ctx context.Context) []*Sukipi {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SukipiCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SukipiCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
