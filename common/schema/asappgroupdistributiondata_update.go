// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asappgroupdistributiondata"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsAppGroupDistributionDataUpdate is the builder for updating AsAppGroupDistributionData entities.
type AsAppGroupDistributionDataUpdate struct {
	config
	hooks    []Hook
	mutation *AsAppGroupDistributionDataMutation
}

// Where appends a list predicates to the AsAppGroupDistributionDataUpdate builder.
func (aagddu *AsAppGroupDistributionDataUpdate) Where(ps ...predicate.AsAppGroupDistributionData) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.Where(ps...)
	return aagddu
}

// SetAppID sets the "app_id" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetAppID(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetAppID(i)
	return aagddu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableAppID(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetAppID(*i)
	}
	return aagddu
}

// ClearAppID clears the value of the "app_id" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearAppID() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearAppID()
	return aagddu
}

// SetGroupID sets the "group_id" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetGroupID(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetGroupID(i)
	return aagddu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableGroupID(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetGroupID(*i)
	}
	return aagddu
}

// ClearGroupID clears the value of the "group_id" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearGroupID() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearGroupID()
	return aagddu
}

// SetConfig sets the "config" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetConfig(s string) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetConfig(s)
	return aagddu
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableConfig(s *string) *AsAppGroupDistributionDataUpdate {
	if s != nil {
		aagddu.SetConfig(*s)
	}
	return aagddu
}

// ClearConfig clears the value of the "config" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearConfig() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearConfig()
	return aagddu
}

// SetIsDeleted sets the "is_deleted" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetIsDeleted(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ResetIsDeleted()
	aagddu.mutation.SetIsDeleted(i)
	return aagddu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableIsDeleted(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetIsDeleted(*i)
	}
	return aagddu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (aagddu *AsAppGroupDistributionDataUpdate) AddIsDeleted(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.AddIsDeleted(i)
	return aagddu
}

// SetStatus sets the "status" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetStatus(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ResetStatus()
	aagddu.mutation.SetStatus(i)
	return aagddu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableStatus(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetStatus(*i)
	}
	return aagddu
}

// AddStatus adds i to the "status" field.
func (aagddu *AsAppGroupDistributionDataUpdate) AddStatus(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.AddStatus(i)
	return aagddu
}

// ClearStatus clears the value of the "status" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearStatus() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearStatus()
	return aagddu
}

// SetCreateUser sets the "create_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetCreateUser(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ResetCreateUser()
	aagddu.mutation.SetCreateUser(i)
	return aagddu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableCreateUser(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetCreateUser(*i)
	}
	return aagddu
}

// AddCreateUser adds i to the "create_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) AddCreateUser(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.AddCreateUser(i)
	return aagddu
}

// ClearCreateUser clears the value of the "create_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearCreateUser() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearCreateUser()
	return aagddu
}

// SetUpdateUser sets the "update_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetUpdateUser(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ResetUpdateUser()
	aagddu.mutation.SetUpdateUser(i)
	return aagddu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableUpdateUser(i *int64) *AsAppGroupDistributionDataUpdate {
	if i != nil {
		aagddu.SetUpdateUser(*i)
	}
	return aagddu
}

// AddUpdateUser adds i to the "update_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) AddUpdateUser(i int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.AddUpdateUser(i)
	return aagddu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearUpdateUser() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearUpdateUser()
	return aagddu
}

// SetUpdateTime sets the "update_time" field.
func (aagddu *AsAppGroupDistributionDataUpdate) SetUpdateTime(dt date.DateTime) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetUpdateTime(dt)
	return aagddu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearUpdateTime() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearUpdateTime()
	return aagddu
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (aagddu *AsAppGroupDistributionDataUpdate) SetAppxID(id int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetAppxID(id)
	return aagddu
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableAppxID(id *int64) *AsAppGroupDistributionDataUpdate {
	if id != nil {
		aagddu = aagddu.SetAppxID(*id)
	}
	return aagddu
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (aagddu *AsAppGroupDistributionDataUpdate) SetAppx(a *AsMarketApp) *AsAppGroupDistributionDataUpdate {
	return aagddu.SetAppxID(a.ID)
}

// SetGroupxID sets the "groupx" edge to the AsAllGroup entity by ID.
func (aagddu *AsAppGroupDistributionDataUpdate) SetGroupxID(id int64) *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.SetGroupxID(id)
	return aagddu
}

// SetNillableGroupxID sets the "groupx" edge to the AsAllGroup entity by ID if the given value is not nil.
func (aagddu *AsAppGroupDistributionDataUpdate) SetNillableGroupxID(id *int64) *AsAppGroupDistributionDataUpdate {
	if id != nil {
		aagddu = aagddu.SetGroupxID(*id)
	}
	return aagddu
}

// SetGroupx sets the "groupx" edge to the AsAllGroup entity.
func (aagddu *AsAppGroupDistributionDataUpdate) SetGroupx(a *AsAllGroup) *AsAppGroupDistributionDataUpdate {
	return aagddu.SetGroupxID(a.ID)
}

// Mutation returns the AsAppGroupDistributionDataMutation object of the builder.
func (aagddu *AsAppGroupDistributionDataUpdate) Mutation() *AsAppGroupDistributionDataMutation {
	return aagddu.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearAppx() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearAppx()
	return aagddu
}

// ClearGroupx clears the "groupx" edge to the AsAllGroup entity.
func (aagddu *AsAppGroupDistributionDataUpdate) ClearGroupx() *AsAppGroupDistributionDataUpdate {
	aagddu.mutation.ClearGroupx()
	return aagddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aagddu *AsAppGroupDistributionDataUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	aagddu.defaults()
	if len(aagddu.hooks) == 0 {
		affected, err = aagddu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsAppGroupDistributionDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aagddu.mutation = mutation
			affected, err = aagddu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aagddu.hooks) - 1; i >= 0; i-- {
			if aagddu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = aagddu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aagddu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aagddu *AsAppGroupDistributionDataUpdate) SaveX(ctx context.Context) int {
	affected, err := aagddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aagddu *AsAppGroupDistributionDataUpdate) Exec(ctx context.Context) error {
	_, err := aagddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aagddu *AsAppGroupDistributionDataUpdate) ExecX(ctx context.Context) {
	if err := aagddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aagddu *AsAppGroupDistributionDataUpdate) defaults() {
	if _, ok := aagddu.mutation.UpdateTime(); !ok && !aagddu.mutation.UpdateTimeCleared() {
		v := asappgroupdistributiondata.UpdateDefaultUpdateTime()
		aagddu.mutation.SetUpdateTime(v)
	}
}

func (aagddu *AsAppGroupDistributionDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asappgroupdistributiondata.Table,
			Columns: asappgroupdistributiondata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asappgroupdistributiondata.FieldID,
			},
		},
	}
	if ps := aagddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aagddu.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asappgroupdistributiondata.FieldConfig,
		})
	}
	if aagddu.mutation.ConfigCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asappgroupdistributiondata.FieldConfig,
		})
	}
	if value, ok := aagddu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldIsDeleted,
		})
	}
	if value, ok := aagddu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldIsDeleted,
		})
	}
	if value, ok := aagddu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if value, ok := aagddu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if aagddu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if value, ok := aagddu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if value, ok := aagddu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if aagddu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if value, ok := aagddu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if value, ok := aagddu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if aagddu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if aagddu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asappgroupdistributiondata.FieldCreateTime,
		})
	}
	if value, ok := aagddu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateTime,
		})
	}
	if aagddu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asappgroupdistributiondata.FieldUpdateTime,
		})
	}
	if aagddu.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.AppxTable,
			Columns: []string{asappgroupdistributiondata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aagddu.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.AppxTable,
			Columns: []string{asappgroupdistributiondata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aagddu.mutation.GroupxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.GroupxTable,
			Columns: []string{asappgroupdistributiondata.GroupxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asallgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aagddu.mutation.GroupxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.GroupxTable,
			Columns: []string{asappgroupdistributiondata.GroupxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asallgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aagddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asappgroupdistributiondata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsAppGroupDistributionDataUpdateOne is the builder for updating a single AsAppGroupDistributionData entity.
type AsAppGroupDistributionDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsAppGroupDistributionDataMutation
}

// SetAppID sets the "app_id" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetAppID(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetAppID(i)
	return aagdduo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableAppID(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetAppID(*i)
	}
	return aagdduo
}

// ClearAppID clears the value of the "app_id" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearAppID() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearAppID()
	return aagdduo
}

// SetGroupID sets the "group_id" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetGroupID(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetGroupID(i)
	return aagdduo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableGroupID(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetGroupID(*i)
	}
	return aagdduo
}

// ClearGroupID clears the value of the "group_id" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearGroupID() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearGroupID()
	return aagdduo
}

// SetConfig sets the "config" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetConfig(s string) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetConfig(s)
	return aagdduo
}

// SetNillableConfig sets the "config" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableConfig(s *string) *AsAppGroupDistributionDataUpdateOne {
	if s != nil {
		aagdduo.SetConfig(*s)
	}
	return aagdduo
}

// ClearConfig clears the value of the "config" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearConfig() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearConfig()
	return aagdduo
}

// SetIsDeleted sets the "is_deleted" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetIsDeleted(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ResetIsDeleted()
	aagdduo.mutation.SetIsDeleted(i)
	return aagdduo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableIsDeleted(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetIsDeleted(*i)
	}
	return aagdduo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) AddIsDeleted(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.AddIsDeleted(i)
	return aagdduo
}

// SetStatus sets the "status" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetStatus(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ResetStatus()
	aagdduo.mutation.SetStatus(i)
	return aagdduo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableStatus(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetStatus(*i)
	}
	return aagdduo
}

// AddStatus adds i to the "status" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) AddStatus(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.AddStatus(i)
	return aagdduo
}

// ClearStatus clears the value of the "status" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearStatus() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearStatus()
	return aagdduo
}

// SetCreateUser sets the "create_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetCreateUser(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ResetCreateUser()
	aagdduo.mutation.SetCreateUser(i)
	return aagdduo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableCreateUser(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetCreateUser(*i)
	}
	return aagdduo
}

// AddCreateUser adds i to the "create_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) AddCreateUser(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.AddCreateUser(i)
	return aagdduo
}

// ClearCreateUser clears the value of the "create_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearCreateUser() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearCreateUser()
	return aagdduo
}

// SetUpdateUser sets the "update_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetUpdateUser(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ResetUpdateUser()
	aagdduo.mutation.SetUpdateUser(i)
	return aagdduo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableUpdateUser(i *int64) *AsAppGroupDistributionDataUpdateOne {
	if i != nil {
		aagdduo.SetUpdateUser(*i)
	}
	return aagdduo
}

// AddUpdateUser adds i to the "update_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) AddUpdateUser(i int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.AddUpdateUser(i)
	return aagdduo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearUpdateUser() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearUpdateUser()
	return aagdduo
}

// SetUpdateTime sets the "update_time" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetUpdateTime(dt date.DateTime) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetUpdateTime(dt)
	return aagdduo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearUpdateTime() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearUpdateTime()
	return aagdduo
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetAppxID(id int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetAppxID(id)
	return aagdduo
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableAppxID(id *int64) *AsAppGroupDistributionDataUpdateOne {
	if id != nil {
		aagdduo = aagdduo.SetAppxID(*id)
	}
	return aagdduo
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetAppx(a *AsMarketApp) *AsAppGroupDistributionDataUpdateOne {
	return aagdduo.SetAppxID(a.ID)
}

// SetGroupxID sets the "groupx" edge to the AsAllGroup entity by ID.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetGroupxID(id int64) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.SetGroupxID(id)
	return aagdduo
}

// SetNillableGroupxID sets the "groupx" edge to the AsAllGroup entity by ID if the given value is not nil.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetNillableGroupxID(id *int64) *AsAppGroupDistributionDataUpdateOne {
	if id != nil {
		aagdduo = aagdduo.SetGroupxID(*id)
	}
	return aagdduo
}

// SetGroupx sets the "groupx" edge to the AsAllGroup entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SetGroupx(a *AsAllGroup) *AsAppGroupDistributionDataUpdateOne {
	return aagdduo.SetGroupxID(a.ID)
}

// Mutation returns the AsAppGroupDistributionDataMutation object of the builder.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) Mutation() *AsAppGroupDistributionDataMutation {
	return aagdduo.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearAppx() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearAppx()
	return aagdduo
}

// ClearGroupx clears the "groupx" edge to the AsAllGroup entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ClearGroupx() *AsAppGroupDistributionDataUpdateOne {
	aagdduo.mutation.ClearGroupx()
	return aagdduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) Select(field string, fields ...string) *AsAppGroupDistributionDataUpdateOne {
	aagdduo.fields = append([]string{field}, fields...)
	return aagdduo
}

// Save executes the query and returns the updated AsAppGroupDistributionData entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) Save(ctx context.Context) (*AsAppGroupDistributionData, error) {
	var (
		err  error
		node *AsAppGroupDistributionData
	)
	aagdduo.defaults()
	if len(aagdduo.hooks) == 0 {
		node, err = aagdduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsAppGroupDistributionDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aagdduo.mutation = mutation
			node, err = aagdduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aagdduo.hooks) - 1; i >= 0; i-- {
			if aagdduo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = aagdduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aagdduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) SaveX(ctx context.Context) *AsAppGroupDistributionData {
	node, err := aagdduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) Exec(ctx context.Context) error {
	_, err := aagdduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) ExecX(ctx context.Context) {
	if err := aagdduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aagdduo *AsAppGroupDistributionDataUpdateOne) defaults() {
	if _, ok := aagdduo.mutation.UpdateTime(); !ok && !aagdduo.mutation.UpdateTimeCleared() {
		v := asappgroupdistributiondata.UpdateDefaultUpdateTime()
		aagdduo.mutation.SetUpdateTime(v)
	}
}

func (aagdduo *AsAppGroupDistributionDataUpdateOne) sqlSave(ctx context.Context) (_node *AsAppGroupDistributionData, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asappgroupdistributiondata.Table,
			Columns: asappgroupdistributiondata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asappgroupdistributiondata.FieldID,
			},
		},
	}
	id, ok := aagdduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsAppGroupDistributionData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aagdduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asappgroupdistributiondata.FieldID)
		for _, f := range fields {
			if !asappgroupdistributiondata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asappgroupdistributiondata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aagdduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aagdduo.mutation.Config(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asappgroupdistributiondata.FieldConfig,
		})
	}
	if aagdduo.mutation.ConfigCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asappgroupdistributiondata.FieldConfig,
		})
	}
	if value, ok := aagdduo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldIsDeleted,
		})
	}
	if value, ok := aagdduo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldIsDeleted,
		})
	}
	if value, ok := aagdduo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if value, ok := aagdduo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if aagdduo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldStatus,
		})
	}
	if value, ok := aagdduo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if value, ok := aagdduo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if aagdduo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldCreateUser,
		})
	}
	if value, ok := aagdduo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if value, ok := aagdduo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if aagdduo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asappgroupdistributiondata.FieldUpdateUser,
		})
	}
	if aagdduo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asappgroupdistributiondata.FieldCreateTime,
		})
	}
	if value, ok := aagdduo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asappgroupdistributiondata.FieldUpdateTime,
		})
	}
	if aagdduo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asappgroupdistributiondata.FieldUpdateTime,
		})
	}
	if aagdduo.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.AppxTable,
			Columns: []string{asappgroupdistributiondata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aagdduo.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.AppxTable,
			Columns: []string{asappgroupdistributiondata.AppxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapp.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if aagdduo.mutation.GroupxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.GroupxTable,
			Columns: []string{asappgroupdistributiondata.GroupxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asallgroup.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := aagdduo.mutation.GroupxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asappgroupdistributiondata.GroupxTable,
			Columns: []string{asappgroupdistributiondata.GroupxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asallgroup.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsAppGroupDistributionData{config: aagdduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aagdduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asappgroupdistributiondata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}