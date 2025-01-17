// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asperson"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsPersonDelete is the builder for deleting a AsPerson entity.
type AsPersonDelete struct {
	config
	hooks    []Hook
	mutation *AsPersonMutation
}

// Where appends a list predicates to the AsPersonDelete builder.
func (apd *AsPersonDelete) Where(ps ...predicate.AsPerson) *AsPersonDelete {
	apd.mutation.Where(ps...)
	return apd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (apd *AsPersonDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(apd.hooks) == 0 {
		affected, err = apd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsPersonMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			apd.mutation = mutation
			affected, err = apd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(apd.hooks) - 1; i >= 0; i-- {
			if apd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = apd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (apd *AsPersonDelete) ExecX(ctx context.Context) int {
	n, err := apd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (apd *AsPersonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asperson.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asperson.FieldID,
			},
		},
	}
	if ps := apd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, apd.driver, _spec)
}

// AsPersonDeleteOne is the builder for deleting a single AsPerson entity.
type AsPersonDeleteOne struct {
	apd *AsPersonDelete
}

// Exec executes the deletion query.
func (apdo *AsPersonDeleteOne) Exec(ctx context.Context) error {
	n, err := apdo.apd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asperson.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (apdo *AsPersonDeleteOne) ExecX(ctx context.Context) {
	apdo.apd.ExecX(ctx)
}
