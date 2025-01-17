// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asallgroup"
	"orginone/common/schema/asmarketapp"
	"orginone/common/schema/asmarketapppurchase"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppPurchaseUpdate is the builder for updating AsMarketAppPurchase entities.
type AsMarketAppPurchaseUpdate struct {
	config
	hooks    []Hook
	mutation *AsMarketAppPurchaseMutation
}

// Where appends a list predicates to the AsMarketAppPurchaseUpdate builder.
func (amapu *AsMarketAppPurchaseUpdate) Where(ps ...predicate.AsMarketAppPurchase) *AsMarketAppPurchaseUpdate {
	amapu.mutation.Where(ps...)
	return amapu
}

// SetAppID sets the "app_id" field.
func (amapu *AsMarketAppPurchaseUpdate) SetAppID(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetAppID(i)
	return amapu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableAppID(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetAppID(*i)
	}
	return amapu
}

// ClearAppID clears the value of the "app_id" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearAppID() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearAppID()
	return amapu
}

// SetTenantID sets the "tenant_id" field.
func (amapu *AsMarketAppPurchaseUpdate) SetTenantID(s string) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetTenantID(s)
	return amapu
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableTenantID(s *string) *AsMarketAppPurchaseUpdate {
	if s != nil {
		amapu.SetTenantID(*s)
	}
	return amapu
}

// ClearTenantID clears the value of the "tenant_id" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearTenantID() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearTenantID()
	return amapu
}

// SetGroupID sets the "group_id" field.
func (amapu *AsMarketAppPurchaseUpdate) SetGroupID(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetGroupID(i)
	return amapu
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableGroupID(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetGroupID(*i)
	}
	return amapu
}

// ClearGroupID clears the value of the "group_id" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearGroupID() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearGroupID()
	return amapu
}

// SetUseStatus sets the "use_status" field.
func (amapu *AsMarketAppPurchaseUpdate) SetUseStatus(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.ResetUseStatus()
	amapu.mutation.SetUseStatus(i)
	return amapu
}

// SetNillableUseStatus sets the "use_status" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableUseStatus(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetUseStatus(*i)
	}
	return amapu
}

// AddUseStatus adds i to the "use_status" field.
func (amapu *AsMarketAppPurchaseUpdate) AddUseStatus(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.AddUseStatus(i)
	return amapu
}

// ClearUseStatus clears the value of the "use_status" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearUseStatus() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearUseStatus()
	return amapu
}

// SetRemark sets the "remark" field.
func (amapu *AsMarketAppPurchaseUpdate) SetRemark(s string) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetRemark(s)
	return amapu
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableRemark(s *string) *AsMarketAppPurchaseUpdate {
	if s != nil {
		amapu.SetRemark(*s)
	}
	return amapu
}

// ClearRemark clears the value of the "remark" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearRemark() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearRemark()
	return amapu
}

// SetIsDeleted sets the "is_deleted" field.
func (amapu *AsMarketAppPurchaseUpdate) SetIsDeleted(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.ResetIsDeleted()
	amapu.mutation.SetIsDeleted(i)
	return amapu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableIsDeleted(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetIsDeleted(*i)
	}
	return amapu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (amapu *AsMarketAppPurchaseUpdate) AddIsDeleted(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.AddIsDeleted(i)
	return amapu
}

// SetStatus sets the "status" field.
func (amapu *AsMarketAppPurchaseUpdate) SetStatus(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.ResetStatus()
	amapu.mutation.SetStatus(i)
	return amapu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableStatus(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetStatus(*i)
	}
	return amapu
}

// AddStatus adds i to the "status" field.
func (amapu *AsMarketAppPurchaseUpdate) AddStatus(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.AddStatus(i)
	return amapu
}

// ClearStatus clears the value of the "status" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearStatus() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearStatus()
	return amapu
}

// SetCreateUser sets the "create_user" field.
func (amapu *AsMarketAppPurchaseUpdate) SetCreateUser(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.ResetCreateUser()
	amapu.mutation.SetCreateUser(i)
	return amapu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableCreateUser(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetCreateUser(*i)
	}
	return amapu
}

// AddCreateUser adds i to the "create_user" field.
func (amapu *AsMarketAppPurchaseUpdate) AddCreateUser(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.AddCreateUser(i)
	return amapu
}

// ClearCreateUser clears the value of the "create_user" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearCreateUser() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearCreateUser()
	return amapu
}

// SetUpdateUser sets the "update_user" field.
func (amapu *AsMarketAppPurchaseUpdate) SetUpdateUser(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.ResetUpdateUser()
	amapu.mutation.SetUpdateUser(i)
	return amapu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableUpdateUser(i *int64) *AsMarketAppPurchaseUpdate {
	if i != nil {
		amapu.SetUpdateUser(*i)
	}
	return amapu
}

// AddUpdateUser adds i to the "update_user" field.
func (amapu *AsMarketAppPurchaseUpdate) AddUpdateUser(i int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.AddUpdateUser(i)
	return amapu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearUpdateUser() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearUpdateUser()
	return amapu
}

// SetUpdateTime sets the "update_time" field.
func (amapu *AsMarketAppPurchaseUpdate) SetUpdateTime(dt date.DateTime) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetUpdateTime(dt)
	return amapu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (amapu *AsMarketAppPurchaseUpdate) ClearUpdateTime() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearUpdateTime()
	return amapu
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (amapu *AsMarketAppPurchaseUpdate) SetAppxID(id int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetAppxID(id)
	return amapu
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableAppxID(id *int64) *AsMarketAppPurchaseUpdate {
	if id != nil {
		amapu = amapu.SetAppxID(*id)
	}
	return amapu
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (amapu *AsMarketAppPurchaseUpdate) SetAppx(a *AsMarketApp) *AsMarketAppPurchaseUpdate {
	return amapu.SetAppxID(a.ID)
}

// SetGroupxID sets the "groupx" edge to the AsAllGroup entity by ID.
func (amapu *AsMarketAppPurchaseUpdate) SetGroupxID(id int64) *AsMarketAppPurchaseUpdate {
	amapu.mutation.SetGroupxID(id)
	return amapu
}

// SetNillableGroupxID sets the "groupx" edge to the AsAllGroup entity by ID if the given value is not nil.
func (amapu *AsMarketAppPurchaseUpdate) SetNillableGroupxID(id *int64) *AsMarketAppPurchaseUpdate {
	if id != nil {
		amapu = amapu.SetGroupxID(*id)
	}
	return amapu
}

// SetGroupx sets the "groupx" edge to the AsAllGroup entity.
func (amapu *AsMarketAppPurchaseUpdate) SetGroupx(a *AsAllGroup) *AsMarketAppPurchaseUpdate {
	return amapu.SetGroupxID(a.ID)
}

// Mutation returns the AsMarketAppPurchaseMutation object of the builder.
func (amapu *AsMarketAppPurchaseUpdate) Mutation() *AsMarketAppPurchaseMutation {
	return amapu.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (amapu *AsMarketAppPurchaseUpdate) ClearAppx() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearAppx()
	return amapu
}

// ClearGroupx clears the "groupx" edge to the AsAllGroup entity.
func (amapu *AsMarketAppPurchaseUpdate) ClearGroupx() *AsMarketAppPurchaseUpdate {
	amapu.mutation.ClearGroupx()
	return amapu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (amapu *AsMarketAppPurchaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	amapu.defaults()
	if len(amapu.hooks) == 0 {
		affected, err = amapu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppPurchaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amapu.mutation = mutation
			affected, err = amapu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amapu.hooks) - 1; i >= 0; i-- {
			if amapu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amapu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amapu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (amapu *AsMarketAppPurchaseUpdate) SaveX(ctx context.Context) int {
	affected, err := amapu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (amapu *AsMarketAppPurchaseUpdate) Exec(ctx context.Context) error {
	_, err := amapu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amapu *AsMarketAppPurchaseUpdate) ExecX(ctx context.Context) {
	if err := amapu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amapu *AsMarketAppPurchaseUpdate) defaults() {
	if _, ok := amapu.mutation.UpdateTime(); !ok && !amapu.mutation.UpdateTimeCleared() {
		v := asmarketapppurchase.UpdateDefaultUpdateTime()
		amapu.mutation.SetUpdateTime(v)
	}
}

func (amapu *AsMarketAppPurchaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketapppurchase.Table,
			Columns: asmarketapppurchase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketapppurchase.FieldID,
			},
		},
	}
	if ps := amapu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amapu.mutation.TenantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketapppurchase.FieldTenantID,
		})
	}
	if amapu.mutation.TenantIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asmarketapppurchase.FieldTenantID,
		})
	}
	if value, ok := amapu.mutation.UseStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if value, ok := amapu.mutation.AddedUseStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if amapu.mutation.UseStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if value, ok := amapu.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketapppurchase.FieldRemark,
		})
	}
	if amapu.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asmarketapppurchase.FieldRemark,
		})
	}
	if value, ok := amapu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldIsDeleted,
		})
	}
	if value, ok := amapu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldIsDeleted,
		})
	}
	if value, ok := amapu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if value, ok := amapu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if amapu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if value, ok := amapu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if value, ok := amapu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if amapu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if value, ok := amapu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if value, ok := amapu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if amapu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if amapu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketapppurchase.FieldCreateTime,
		})
	}
	if value, ok := amapu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateTime,
		})
	}
	if amapu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketapppurchase.FieldUpdateTime,
		})
	}
	if amapu.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.AppxTable,
			Columns: []string{asmarketapppurchase.AppxColumn},
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
	if nodes := amapu.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.AppxTable,
			Columns: []string{asmarketapppurchase.AppxColumn},
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
	if amapu.mutation.GroupxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.GroupxTable,
			Columns: []string{asmarketapppurchase.GroupxColumn},
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
	if nodes := amapu.mutation.GroupxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.GroupxTable,
			Columns: []string{asmarketapppurchase.GroupxColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, amapu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketapppurchase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsMarketAppPurchaseUpdateOne is the builder for updating a single AsMarketAppPurchase entity.
type AsMarketAppPurchaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsMarketAppPurchaseMutation
}

// SetAppID sets the "app_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetAppID(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetAppID(i)
	return amapuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableAppID(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetAppID(*i)
	}
	return amapuo
}

// ClearAppID clears the value of the "app_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearAppID() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearAppID()
	return amapuo
}

// SetTenantID sets the "tenant_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetTenantID(s string) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetTenantID(s)
	return amapuo
}

// SetNillableTenantID sets the "tenant_id" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableTenantID(s *string) *AsMarketAppPurchaseUpdateOne {
	if s != nil {
		amapuo.SetTenantID(*s)
	}
	return amapuo
}

// ClearTenantID clears the value of the "tenant_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearTenantID() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearTenantID()
	return amapuo
}

// SetGroupID sets the "group_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetGroupID(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetGroupID(i)
	return amapuo
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableGroupID(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetGroupID(*i)
	}
	return amapuo
}

// ClearGroupID clears the value of the "group_id" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearGroupID() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearGroupID()
	return amapuo
}

// SetUseStatus sets the "use_status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetUseStatus(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ResetUseStatus()
	amapuo.mutation.SetUseStatus(i)
	return amapuo
}

// SetNillableUseStatus sets the "use_status" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableUseStatus(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetUseStatus(*i)
	}
	return amapuo
}

// AddUseStatus adds i to the "use_status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) AddUseStatus(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.AddUseStatus(i)
	return amapuo
}

// ClearUseStatus clears the value of the "use_status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearUseStatus() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearUseStatus()
	return amapuo
}

// SetRemark sets the "remark" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetRemark(s string) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetRemark(s)
	return amapuo
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableRemark(s *string) *AsMarketAppPurchaseUpdateOne {
	if s != nil {
		amapuo.SetRemark(*s)
	}
	return amapuo
}

// ClearRemark clears the value of the "remark" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearRemark() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearRemark()
	return amapuo
}

// SetIsDeleted sets the "is_deleted" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetIsDeleted(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ResetIsDeleted()
	amapuo.mutation.SetIsDeleted(i)
	return amapuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableIsDeleted(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetIsDeleted(*i)
	}
	return amapuo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) AddIsDeleted(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.AddIsDeleted(i)
	return amapuo
}

// SetStatus sets the "status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetStatus(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ResetStatus()
	amapuo.mutation.SetStatus(i)
	return amapuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableStatus(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetStatus(*i)
	}
	return amapuo
}

// AddStatus adds i to the "status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) AddStatus(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.AddStatus(i)
	return amapuo
}

// ClearStatus clears the value of the "status" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearStatus() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearStatus()
	return amapuo
}

// SetCreateUser sets the "create_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetCreateUser(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ResetCreateUser()
	amapuo.mutation.SetCreateUser(i)
	return amapuo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableCreateUser(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetCreateUser(*i)
	}
	return amapuo
}

// AddCreateUser adds i to the "create_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) AddCreateUser(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.AddCreateUser(i)
	return amapuo
}

// ClearCreateUser clears the value of the "create_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearCreateUser() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearCreateUser()
	return amapuo
}

// SetUpdateUser sets the "update_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetUpdateUser(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ResetUpdateUser()
	amapuo.mutation.SetUpdateUser(i)
	return amapuo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableUpdateUser(i *int64) *AsMarketAppPurchaseUpdateOne {
	if i != nil {
		amapuo.SetUpdateUser(*i)
	}
	return amapuo
}

// AddUpdateUser adds i to the "update_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) AddUpdateUser(i int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.AddUpdateUser(i)
	return amapuo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearUpdateUser() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearUpdateUser()
	return amapuo
}

// SetUpdateTime sets the "update_time" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetUpdateTime(dt date.DateTime) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetUpdateTime(dt)
	return amapuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearUpdateTime() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearUpdateTime()
	return amapuo
}

// SetAppxID sets the "appx" edge to the AsMarketApp entity by ID.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetAppxID(id int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetAppxID(id)
	return amapuo
}

// SetNillableAppxID sets the "appx" edge to the AsMarketApp entity by ID if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableAppxID(id *int64) *AsMarketAppPurchaseUpdateOne {
	if id != nil {
		amapuo = amapuo.SetAppxID(*id)
	}
	return amapuo
}

// SetAppx sets the "appx" edge to the AsMarketApp entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetAppx(a *AsMarketApp) *AsMarketAppPurchaseUpdateOne {
	return amapuo.SetAppxID(a.ID)
}

// SetGroupxID sets the "groupx" edge to the AsAllGroup entity by ID.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetGroupxID(id int64) *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.SetGroupxID(id)
	return amapuo
}

// SetNillableGroupxID sets the "groupx" edge to the AsAllGroup entity by ID if the given value is not nil.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetNillableGroupxID(id *int64) *AsMarketAppPurchaseUpdateOne {
	if id != nil {
		amapuo = amapuo.SetGroupxID(*id)
	}
	return amapuo
}

// SetGroupx sets the "groupx" edge to the AsAllGroup entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) SetGroupx(a *AsAllGroup) *AsMarketAppPurchaseUpdateOne {
	return amapuo.SetGroupxID(a.ID)
}

// Mutation returns the AsMarketAppPurchaseMutation object of the builder.
func (amapuo *AsMarketAppPurchaseUpdateOne) Mutation() *AsMarketAppPurchaseMutation {
	return amapuo.mutation
}

// ClearAppx clears the "appx" edge to the AsMarketApp entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearAppx() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearAppx()
	return amapuo
}

// ClearGroupx clears the "groupx" edge to the AsAllGroup entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) ClearGroupx() *AsMarketAppPurchaseUpdateOne {
	amapuo.mutation.ClearGroupx()
	return amapuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (amapuo *AsMarketAppPurchaseUpdateOne) Select(field string, fields ...string) *AsMarketAppPurchaseUpdateOne {
	amapuo.fields = append([]string{field}, fields...)
	return amapuo
}

// Save executes the query and returns the updated AsMarketAppPurchase entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) Save(ctx context.Context) (*AsMarketAppPurchase, error) {
	var (
		err  error
		node *AsMarketAppPurchase
	)
	amapuo.defaults()
	if len(amapuo.hooks) == 0 {
		node, err = amapuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketAppPurchaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amapuo.mutation = mutation
			node, err = amapuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(amapuo.hooks) - 1; i >= 0; i-- {
			if amapuo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amapuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amapuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (amapuo *AsMarketAppPurchaseUpdateOne) SaveX(ctx context.Context) *AsMarketAppPurchase {
	node, err := amapuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (amapuo *AsMarketAppPurchaseUpdateOne) Exec(ctx context.Context) error {
	_, err := amapuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amapuo *AsMarketAppPurchaseUpdateOne) ExecX(ctx context.Context) {
	if err := amapuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amapuo *AsMarketAppPurchaseUpdateOne) defaults() {
	if _, ok := amapuo.mutation.UpdateTime(); !ok && !amapuo.mutation.UpdateTimeCleared() {
		v := asmarketapppurchase.UpdateDefaultUpdateTime()
		amapuo.mutation.SetUpdateTime(v)
	}
}

func (amapuo *AsMarketAppPurchaseUpdateOne) sqlSave(ctx context.Context) (_node *AsMarketAppPurchase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketapppurchase.Table,
			Columns: asmarketapppurchase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketapppurchase.FieldID,
			},
		},
	}
	id, ok := amapuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsMarketAppPurchase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := amapuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketapppurchase.FieldID)
		for _, f := range fields {
			if !asmarketapppurchase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asmarketapppurchase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := amapuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amapuo.mutation.TenantID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketapppurchase.FieldTenantID,
		})
	}
	if amapuo.mutation.TenantIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asmarketapppurchase.FieldTenantID,
		})
	}
	if value, ok := amapuo.mutation.UseStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if value, ok := amapuo.mutation.AddedUseStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if amapuo.mutation.UseStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldUseStatus,
		})
	}
	if value, ok := amapuo.mutation.Remark(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asmarketapppurchase.FieldRemark,
		})
	}
	if amapuo.mutation.RemarkCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: asmarketapppurchase.FieldRemark,
		})
	}
	if value, ok := amapuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldIsDeleted,
		})
	}
	if value, ok := amapuo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldIsDeleted,
		})
	}
	if value, ok := amapuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if value, ok := amapuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if amapuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldStatus,
		})
	}
	if value, ok := amapuo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if value, ok := amapuo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if amapuo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldCreateUser,
		})
	}
	if value, ok := amapuo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if value, ok := amapuo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if amapuo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketapppurchase.FieldUpdateUser,
		})
	}
	if amapuo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketapppurchase.FieldCreateTime,
		})
	}
	if value, ok := amapuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketapppurchase.FieldUpdateTime,
		})
	}
	if amapuo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketapppurchase.FieldUpdateTime,
		})
	}
	if amapuo.mutation.AppxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.AppxTable,
			Columns: []string{asmarketapppurchase.AppxColumn},
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
	if nodes := amapuo.mutation.AppxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.AppxTable,
			Columns: []string{asmarketapppurchase.AppxColumn},
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
	if amapuo.mutation.GroupxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.GroupxTable,
			Columns: []string{asmarketapppurchase.GroupxColumn},
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
	if nodes := amapuo.mutation.GroupxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketapppurchase.GroupxTable,
			Columns: []string{asmarketapppurchase.GroupxColumn},
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
	_node = &AsMarketAppPurchase{config: amapuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, amapuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketapppurchase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
