// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"flehmen-api/ent/predicate"
	"flehmen-api/ent/sukipi"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SukipiDelete is the builder for deleting a Sukipi entity.
type SukipiDelete struct {
	config
	hooks    []Hook
	mutation *SukipiMutation
}

// Where appends a list predicates to the SukipiDelete builder.
func (sd *SukipiDelete) Where(ps ...predicate.Sukipi) *SukipiDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SukipiDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SukipiDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SukipiDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sukipi.Table, sqlgraph.NewFieldSpec(sukipi.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SukipiDeleteOne is the builder for deleting a single Sukipi entity.
type SukipiDeleteOne struct {
	sd *SukipiDelete
}

// Where appends a list predicates to the SukipiDelete builder.
func (sdo *SukipiDeleteOne) Where(ps ...predicate.Sukipi) *SukipiDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SukipiDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sukipi.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SukipiDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
