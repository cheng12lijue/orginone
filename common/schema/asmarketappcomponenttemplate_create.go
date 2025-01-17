// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asmarketappcomponenttemplate"
	"orginone/common/schema/asmarketappusertemplate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppComponentTemplateCreate is the builder for creating a AsMarketAppComponentTemplate entity.
type AsMarketAppComponentTemplateCreate struct {
	config
	mutation *AsMarketAppComponentTemplateMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetName(s string) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetName(s)
	return amactc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableName(s *string) *AsMarketAppComponentTemplateCreate {
	if s != nil {
		amactc.SetName(*s)
	}
	return amactc
}

// SetConfig sets the "config" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetConfig(s string) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetConfig(s)
	return amactc
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableConfig(s *string) *AsMarketAppComponentTemplateCreate {
	if s != nil {
		amactc.SetConfig(*s)
	}
	return amactc
}

// SetIsDefault sets the "is_default" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetIsDefault(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetIsDefault(i)
	return amactc
}

// SetNillableIsDefault sets the "is_default" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableIsDefault(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetIsDefault(*i)
	}
	return amactc
}

// SetIsDeleted sets the "is_deleted" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetIsDeleted(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetIsDeleted(i)
	return amactc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableIsDeleted(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetIsDeleted(*i)
	}
	return amactc
}

// SetStatus sets the "status" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetStatus(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetStatus(i)
	return amactc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableStatus(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetStatus(*i)
	}
	return amactc
}

// SetCreateUser sets the "create_user" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetCreateUser(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetCreateUser(i)
	return amactc
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableCreateUser(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetCreateUser(*i)
	}
	return amactc
}

// SetUpdateUser sets the "update_user" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetUpdateUser(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetUpdateUser(i)
	return amactc
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableUpdateUser(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetUpdateUser(*i)
	}
	return amactc
}

// SetCreateTime sets the "create_time" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetCreateTime(dt date.DateTime) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetCreateTime(dt)
	return amactc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableCreateTime(dt *date.DateTime) *AsMarketAppComponentTemplateCreate {
	if dt != nil {
		amactc.SetCreateTime(*dt)
	}
	return amactc
}

// SetUpdateTime sets the "update_time" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetUpdateTime(dt date.DateTime) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetUpdateTime(dt)
	return amactc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableUpdateTime(dt *date.DateTime) *AsMarketAppComponentTemplateCreate {
	if dt != nil {
		amactc.SetUpdateTime(*dt)
	}
	return amactc
}

// SetID sets the "id" field.
func (amactc *AsMarketAppComponentTemplateCreate) SetID(i int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.SetID(i)
	return amactc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (amactc *AsMarketAppComponentTemplateCreate) SetNillableID(i *int64) *AsMarketAppComponentTemplateCreate {
	if i != nil {
		amactc.SetID(*i)
	}
	return amactc
}

// AddAppUserTemplateIDs adds the "appUserTemplates" edge to the AsMarketAppUserTemplate entity by IDs.
func (amactc *AsMarketAppComponentTemplateCreate) AddAppUserTemplateIDs(ids ...int64) *AsMarketAppComponentTemplateCreate {
	amactc.mutation.AddAppUserTemplateIDs(ids...)
	return amactc
}

// AddAppUserTemplates adds the "appUserTemplates" edges to the AsMarketAppUserTemplate entity.
func (amactc *AsMarketAppComponentTemplateCreate) AddAppUserTemplates(a ...*AsMarketAppUserTemplate) *AsMarketAppComponentTemplateCreate {
	ids := make([]int64, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return amactc.AddAppUserTemplateIDs(ids...)
}

// Mutation returns the AsMarketAppComponentTemplateMutation object of the builder.
func (amactc *AsMarketAppComponentTemplateCreate) Mutation() *AsMarketAppComponentTemplateMutation {
	return amactc.mutation
}

// Save creates the AsMarketAppComponentTemplate in the database.
func (amactc *AsMarketAppComponentTemplateCreate) Save(ctx context.Context) (*AsMarketAppComponentTemplate, error) {
	var (
		err  error
		node *AsMarketAppComponentTemplate
	)
	amactc.defaults()
	if len(amactc.hooks) == 0 {
		if err = amactc.check(); err != nil {
			return nil, err
		}
		node, err = amactc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppComponentTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = amactc.check(); err != nil {
				return nil, err
			}
			amactc.mutation = mutation
			if node, err = amactc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(amactc.hooks) - 1; i >= 0; i-- {
			if amactc.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amactc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amactc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (amactc *AsMarketAppComponentTemplateCreate) SaveX(ctx context.Context) *AsMarketAppComponentTemplate {
	v, err := amactc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amactc *AsMarketAppComponentTemplateCreate) Exec(ctx context.Context) error {
	_, err := amactc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amactc *AsMarketAppComponentTemplateCreate) ExecX(ctx context.Context) {
	if err := amactc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amactc *AsMarketAppComponentTemplateCreate) defaults() {
	if _, ok := amactc.mutation.IsDeleted(); !ok {
		v := asmarketappcomponenttemplate.DefaultIsDeleted
		amactc.mutation.SetIsDeleted(v)
	}
	if _, ok := amactc.mutation.Status(); !ok {
		v := asmarketappcomponenttemplate.DefaultStatus
		amactc.mutation.SetStatus(v)
	}
	if _, ok := amactc.mutation.CreateTime(); !ok {
		v := asmarketappcomponenttemplate.DefaultCreateTime()
		amactc.mutation.SetCreateTime(v)
	}
	if _, ok := amactc.mutation.UpdateTime(); !ok {
		v := asmarketappcomponenttemplate.DefaultUpdateTime()
		amactc.mutation.SetUpdateTime(v)
	}
	if _, ok := amactc.mutation.ID(); !ok {
		v := asmarketappcomponenttemplate.DefaultID()
		amactc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (amactc *AsMarketAppComponentTemplateCreate) check() error {
	if _, ok := amactc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`schema: missing required field "AsMarketAppComponentTemplate.is_deleted"`)}
	}
	return nil
}

func (amactc *AsMarketAppComponentTemplateCreate) sqlSave(ctx context.Context) (*AsMarketAppComponentTemplate, error) {
	_node, _spec := amactc.createSpec()
	if err := sqlgraph.CreateNode(ctx, amactc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (amactc *AsMarketAppComponentTemplateCreate) createSpec() (*AsMarketAppComponentTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &AsMarketAppComponentTemplate{config: amactc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asmarketappcomponenttemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappcomponenttemplate.FieldID,
			},
		}
	)
	if id, ok := amactc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := amactc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldName,
		})
		_node.Name = value
	}
	if value, ok := amactc.mutation.Config(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldConfig,
		})
		_node.Config = value
	}
	if value, ok := amactc.mutation.IsDefault(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldIsDefault,
		})
		_node.IsDefault = value
	}
	if value, ok := amactc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := amactc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := amactc.mutation.CreateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldCreateUser,
		})
		_node.CreateUser = value
	}
	if value, ok := amactc.mutation.UpdateUser(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldUpdateUser,
		})
		_node.UpdateUser = value
	}
	if value, ok := amactc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := amactc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketappcomponenttemplate.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if nodes := amactc.mutation.AppUserTemplatesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   asmarketappcomponenttemplate.AppUserTemplatesTable,
			Columns: []string{asmarketappcomponenttemplate.AppUserTemplatesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketappusertemplate.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AsMarketAppComponentTemplateCreateBulk is the builder for creating many AsMarketAppComponentTemplate entities in bulk.
type AsMarketAppComponentTemplateCreateBulk struct {
	config
	builders []*AsMarketAppComponentTemplateCreate
}

// Save creates the AsMarketAppComponentTemplate entities in the database.
func (amactcb *AsMarketAppComponentTemplateCreateBulk) Save(ctx context.Context) ([]*AsMarketAppComponentTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(amactcb.builders))
	nodes := make([]*AsMarketAppComponentTemplate, len(amactcb.builders))
	mutators := make([]Mutator, len(amactcb.builders))
	for i := range amactcb.builders {
		func(i int, root context.Context) {
			builder := amactcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AsMarketAppComponentTemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, amactcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, amactcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, amactcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (amactcb *AsMarketAppComponentTemplateCreateBulk) SaveX(ctx context.Context) []*AsMarketAppComponentTemplate {
	v, err := amactcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (amactcb *AsMarketAppComponentTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := amactcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amactcb *AsMarketAppComponentTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := amactcb.Exec(ctx); err != nil {
		panic(err)
	}
}
