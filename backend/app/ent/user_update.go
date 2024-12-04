// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"flehmen-api/ent/mbti"
	"flehmen-api/ent/predicate"
	"flehmen-api/ent/user"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetWeight sets the "weight" field.
func (uu *UserUpdate) SetWeight(f float64) *UserUpdate {
	uu.mutation.ResetWeight()
	uu.mutation.SetWeight(f)
	return uu
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uu *UserUpdate) SetNillableWeight(f *float64) *UserUpdate {
	if f != nil {
		uu.SetWeight(*f)
	}
	return uu
}

// AddWeight adds f to the "weight" field.
func (uu *UserUpdate) AddWeight(f float64) *UserUpdate {
	uu.mutation.AddWeight(f)
	return uu
}

// SetHeight sets the "height" field.
func (uu *UserUpdate) SetHeight(f float64) *UserUpdate {
	uu.mutation.ResetHeight()
	uu.mutation.SetHeight(f)
	return uu
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uu *UserUpdate) SetNillableHeight(f *float64) *UserUpdate {
	if f != nil {
		uu.SetHeight(*f)
	}
	return uu
}

// AddHeight adds f to the "height" field.
func (uu *UserUpdate) AddHeight(f float64) *UserUpdate {
	uu.mutation.AddHeight(f)
	return uu
}

// SetClerkID sets the "clerk_id" field.
func (uu *UserUpdate) SetClerkID(s string) *UserUpdate {
	uu.mutation.SetClerkID(s)
	return uu
}

// SetNillableClerkID sets the "clerk_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableClerkID(s *string) *UserUpdate {
	if s != nil {
		uu.SetClerkID(*s)
	}
	return uu
}

// ClearClerkID clears the value of the "clerk_id" field.
func (uu *UserUpdate) ClearClerkID() *UserUpdate {
	uu.mutation.ClearClerkID()
	return uu
}

// SetIsMale sets the "is_male" field.
func (uu *UserUpdate) SetIsMale(b bool) *UserUpdate {
	uu.mutation.SetIsMale(b)
	return uu
}

// SetNillableIsMale sets the "is_male" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsMale(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsMale(*b)
	}
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetMbtiID sets the "mbti" edge to the Mbti entity by ID.
func (uu *UserUpdate) SetMbtiID(id int) *UserUpdate {
	uu.mutation.SetMbtiID(id)
	return uu
}

// SetNillableMbtiID sets the "mbti" edge to the Mbti entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableMbtiID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetMbtiID(*id)
	}
	return uu
}

// SetMbti sets the "mbti" edge to the Mbti entity.
func (uu *UserUpdate) SetMbti(m *Mbti) *UserUpdate {
	return uu.SetMbtiID(m.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearMbti clears the "mbti" edge to the Mbti entity.
func (uu *UserUpdate) ClearMbti() *UserUpdate {
	uu.mutation.ClearMbti()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uu.mutation.ClerkID(); ok {
		_spec.SetField(user.FieldClerkID, field.TypeString, value)
	}
	if uu.mutation.ClerkIDCleared() {
		_spec.ClearField(user.FieldClerkID, field.TypeString)
	}
	if value, ok := uu.mutation.IsMale(); ok {
		_spec.SetField(user.FieldIsMale, field.TypeBool, value)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if uu.mutation.MbtiCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.MbtiIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetWeight sets the "weight" field.
func (uuo *UserUpdateOne) SetWeight(f float64) *UserUpdateOne {
	uuo.mutation.ResetWeight()
	uuo.mutation.SetWeight(f)
	return uuo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableWeight(f *float64) *UserUpdateOne {
	if f != nil {
		uuo.SetWeight(*f)
	}
	return uuo
}

// AddWeight adds f to the "weight" field.
func (uuo *UserUpdateOne) AddWeight(f float64) *UserUpdateOne {
	uuo.mutation.AddWeight(f)
	return uuo
}

// SetHeight sets the "height" field.
func (uuo *UserUpdateOne) SetHeight(f float64) *UserUpdateOne {
	uuo.mutation.ResetHeight()
	uuo.mutation.SetHeight(f)
	return uuo
}

// SetNillableHeight sets the "height" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableHeight(f *float64) *UserUpdateOne {
	if f != nil {
		uuo.SetHeight(*f)
	}
	return uuo
}

// AddHeight adds f to the "height" field.
func (uuo *UserUpdateOne) AddHeight(f float64) *UserUpdateOne {
	uuo.mutation.AddHeight(f)
	return uuo
}

// SetClerkID sets the "clerk_id" field.
func (uuo *UserUpdateOne) SetClerkID(s string) *UserUpdateOne {
	uuo.mutation.SetClerkID(s)
	return uuo
}

// SetNillableClerkID sets the "clerk_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableClerkID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetClerkID(*s)
	}
	return uuo
}

// ClearClerkID clears the value of the "clerk_id" field.
func (uuo *UserUpdateOne) ClearClerkID() *UserUpdateOne {
	uuo.mutation.ClearClerkID()
	return uuo
}

// SetIsMale sets the "is_male" field.
func (uuo *UserUpdateOne) SetIsMale(b bool) *UserUpdateOne {
	uuo.mutation.SetIsMale(b)
	return uuo
}

// SetNillableIsMale sets the "is_male" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsMale(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsMale(*b)
	}
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetMbtiID sets the "mbti" edge to the Mbti entity by ID.
func (uuo *UserUpdateOne) SetMbtiID(id int) *UserUpdateOne {
	uuo.mutation.SetMbtiID(id)
	return uuo
}

// SetNillableMbtiID sets the "mbti" edge to the Mbti entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMbtiID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetMbtiID(*id)
	}
	return uuo
}

// SetMbti sets the "mbti" edge to the Mbti entity.
func (uuo *UserUpdateOne) SetMbti(m *Mbti) *UserUpdateOne {
	return uuo.SetMbtiID(m.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearMbti clears the "mbti" edge to the Mbti entity.
func (uuo *UserUpdateOne) ClearMbti() *UserUpdateOne {
	uuo.mutation.ClearMbti()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
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
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Weight(); ok {
		_spec.SetField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedWeight(); ok {
		_spec.AddField(user.FieldWeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.Height(); ok {
		_spec.SetField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.AddedHeight(); ok {
		_spec.AddField(user.FieldHeight, field.TypeFloat64, value)
	}
	if value, ok := uuo.mutation.ClerkID(); ok {
		_spec.SetField(user.FieldClerkID, field.TypeString, value)
	}
	if uuo.mutation.ClerkIDCleared() {
		_spec.ClearField(user.FieldClerkID, field.TypeString)
	}
	if value, ok := uuo.mutation.IsMale(); ok {
		_spec.SetField(user.FieldIsMale, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if uuo.mutation.MbtiCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.MbtiIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
