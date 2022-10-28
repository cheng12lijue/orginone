// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"fmt"
	"orginone/common/schema/asmarketappusertemplate"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppUserTemplateDelete is the builder for deleting a AsMarketAppUserTemplate entity.
type AsMarketAppUserTemplateDelete struct {
	config
	hooks    []Hook
	mutation *AsMarketAppUserTemplateMutation
}

// Where appends a list predicates to the AsMarketAppUserTemplateDelete builder.
func (amautd *AsMarketAppUserTemplateDelete) Where(ps ...predicate.AsMarketAppUserTemplate) *AsMarketAppUserTemplateDelete {
	amautd.mutation.Where(ps...)
	return amautd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amautd *AsMarketAppUserTemplateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(amautd.hooks) == 0 {
		affected, err = amautd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppUserTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amautd.mutation = mutation
			affected, err = amautd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amautd.hooks) - 1; i >= 0; i-- {
			if amautd.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amautd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amautd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (amautd *AsMarketAppUserTemplateDelete) ExecX(ctx context.Context) int {
	n, err := amautd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amautd *AsMarketAppUserTemplateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: asmarketappusertemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappusertemplate.FieldID,
			},
		},
	}
	if ps := amautd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, amautd.driver, _spec)
}

// AsMarketAppUserTemplateDeleteOne is the builder for deleting a single AsMarketAppUserTemplate entity.
type AsMarketAppUserTemplateDeleteOne struct {
	amautd *AsMarketAppUserTemplateDelete
}

// Exec executes the deletion query.
func (amautdo *AsMarketAppUserTemplateDeleteOne) Exec(ctx context.Context) error {
	n, err := amautdo.amautd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{asmarketappusertemplate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amautdo *AsMarketAppUserTemplateDeleteOne) ExecX(ctx context.Context) {
	amautdo.amautd.ExecX(ctx)
}