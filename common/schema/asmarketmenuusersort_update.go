// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketmenuusersort"
	"orginone/common/schema/asuser"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketMenuUserSortUpdate is the builder for updating AsMarketMenuUserSort entities.
type AsMarketMenuUserSortUpdate struct {
	config
	hooks    []Hook
	mutation *AsMarketMenuUserSortMutation
}

// Where appends a list predicates to the AsMarketMenuUserSortUpdate builder.
func (ammusu *AsMarketMenuUserSortUpdate) Where(ps ...predicate.AsMarketMenuUserSort) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.Where(ps...)
	return ammusu
}

// SetUserID sets the "user_id" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetUserID(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.SetUserID(i)
	return ammusu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableUserID(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetUserID(*i)
	}
	return ammusu
}

// ClearUserID clears the value of the "user_id" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearUserID() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearUserID()
	return ammusu
}

// SetMenuID sets the "menu_id" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetMenuID(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.SetMenuID(i)
	return ammusu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableMenuID(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetMenuID(*i)
	}
	return ammusu
}

// ClearMenuID clears the value of the "menu_id" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearMenuID() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearMenuID()
	return ammusu
}

// SetSort sets the "sort" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetSort(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ResetSort()
	ammusu.mutation.SetSort(i)
	return ammusu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableSort(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetSort(*i)
	}
	return ammusu
}

// AddSort adds i to the "sort" field.
func (ammusu *AsMarketMenuUserSortUpdate) AddSort(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.AddSort(i)
	return ammusu
}

// ClearSort clears the value of the "sort" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearSort() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearSort()
	return ammusu
}

// SetIsDeleted sets the "is_deleted" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetIsDeleted(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ResetIsDeleted()
	ammusu.mutation.SetIsDeleted(i)
	return ammusu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableIsDeleted(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetIsDeleted(*i)
	}
	return ammusu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (ammusu *AsMarketMenuUserSortUpdate) AddIsDeleted(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.AddIsDeleted(i)
	return ammusu
}

// SetStatus sets the "status" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetStatus(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ResetStatus()
	ammusu.mutation.SetStatus(i)
	return ammusu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableStatus(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetStatus(*i)
	}
	return ammusu
}

// AddStatus adds i to the "status" field.
func (ammusu *AsMarketMenuUserSortUpdate) AddStatus(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.AddStatus(i)
	return ammusu
}

// ClearStatus clears the value of the "status" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearStatus() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearStatus()
	return ammusu
}

// SetCreateUser sets the "create_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetCreateUser(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ResetCreateUser()
	ammusu.mutation.SetCreateUser(i)
	return ammusu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableCreateUser(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetCreateUser(*i)
	}
	return ammusu
}

// AddCreateUser adds i to the "create_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) AddCreateUser(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.AddCreateUser(i)
	return ammusu
}

// ClearCreateUser clears the value of the "create_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearCreateUser() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearCreateUser()
	return ammusu
}

// SetUpdateUser sets the "update_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetUpdateUser(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ResetUpdateUser()
	ammusu.mutation.SetUpdateUser(i)
	return ammusu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableUpdateUser(i *int64) *AsMarketMenuUserSortUpdate {
	if i != nil {
		ammusu.SetUpdateUser(*i)
	}
	return ammusu
}

// AddUpdateUser adds i to the "update_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) AddUpdateUser(i int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.AddUpdateUser(i)
	return ammusu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearUpdateUser() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearUpdateUser()
	return ammusu
}

// SetUpdateTime sets the "update_time" field.
func (ammusu *AsMarketMenuUserSortUpdate) SetUpdateTime(dt date.DateTime) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.SetUpdateTime(dt)
	return ammusu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ammusu *AsMarketMenuUserSortUpdate) ClearUpdateTime() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearUpdateTime()
	return ammusu
}

// SetUserxID sets the "userx" edge to the AsUser entity by ID.
func (ammusu *AsMarketMenuUserSortUpdate) SetUserxID(id int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.SetUserxID(id)
	return ammusu
}

// SetNillableUserxID sets the "userx" edge to the AsUser entity by ID if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableUserxID(id *int64) *AsMarketMenuUserSortUpdate {
	if id != nil {
		ammusu = ammusu.SetUserxID(*id)
	}
	return ammusu
}

// SetUserx sets the "userx" edge to the AsUser entity.
func (ammusu *AsMarketMenuUserSortUpdate) SetUserx(a *AsUser) *AsMarketMenuUserSortUpdate {
	return ammusu.SetUserxID(a.ID)
}

// SetAppmenuxID sets the "appmenux" edge to the AsMarketMenu entity by ID.
func (ammusu *AsMarketMenuUserSortUpdate) SetAppmenuxID(id int64) *AsMarketMenuUserSortUpdate {
	ammusu.mutation.SetAppmenuxID(id)
	return ammusu
}

// SetNillableAppmenuxID sets the "appmenux" edge to the AsMarketMenu entity by ID if the given value is not nil.
func (ammusu *AsMarketMenuUserSortUpdate) SetNillableAppmenuxID(id *int64) *AsMarketMenuUserSortUpdate {
	if id != nil {
		ammusu = ammusu.SetAppmenuxID(*id)
	}
	return ammusu
}

// SetAppmenux sets the "appmenux" edge to the AsMarketMenu entity.
func (ammusu *AsMarketMenuUserSortUpdate) SetAppmenux(a *AsMarketMenu) *AsMarketMenuUserSortUpdate {
	return ammusu.SetAppmenuxID(a.ID)
}

// Mutation returns the AsMarketMenuUserSortMutation object of the builder.
func (ammusu *AsMarketMenuUserSortUpdate) Mutation() *AsMarketMenuUserSortMutation {
	return ammusu.mutation
}

// ClearUserx clears the "userx" edge to the AsUser entity.
func (ammusu *AsMarketMenuUserSortUpdate) ClearUserx() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearUserx()
	return ammusu
}

// ClearAppmenux clears the "appmenux" edge to the AsMarketMenu entity.
func (ammusu *AsMarketMenuUserSortUpdate) ClearAppmenux() *AsMarketMenuUserSortUpdate {
	ammusu.mutation.ClearAppmenux()
	return ammusu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ammusu *AsMarketMenuUserSortUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ammusu.defaults()
	if len(ammusu.hooks) == 0 {
		affected, err = ammusu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketMenuUserSortMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ammusu.mutation = mutation
			affected, err = ammusu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ammusu.hooks) - 1; i >= 0; i-- {
			if ammusu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = ammusu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ammusu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ammusu *AsMarketMenuUserSortUpdate) SaveX(ctx context.Context) int {
	affected, err := ammusu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ammusu *AsMarketMenuUserSortUpdate) Exec(ctx context.Context) error {
	_, err := ammusu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ammusu *AsMarketMenuUserSortUpdate) ExecX(ctx context.Context) {
	if err := ammusu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ammusu *AsMarketMenuUserSortUpdate) defaults() {
	if _, ok := ammusu.mutation.UpdateTime(); !ok && !ammusu.mutation.UpdateTimeCleared() {
		v := asmarketmenuusersort.UpdateDefaultUpdateTime()
		ammusu.mutation.SetUpdateTime(v)
	}
}

func (ammusu *AsMarketMenuUserSortUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketmenuusersort.Table,
			Columns: asmarketmenuusersort.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketmenuusersort.FieldID,
			},
		},
	}
	if ps := ammusu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ammusu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if value, ok := ammusu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if ammusu.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if value, ok := ammusu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldIsDeleted,
		})
	}
	if value, ok := ammusu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldIsDeleted,
		})
	}
	if value, ok := ammusu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if value, ok := ammusu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if ammusu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if value, ok := ammusu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if value, ok := ammusu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if ammusu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if value, ok := ammusu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if value, ok := ammusu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if ammusu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if ammusu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketmenuusersort.FieldCreateTime,
		})
	}
	if value, ok := ammusu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateTime,
		})
	}
	if ammusu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketmenuusersort.FieldUpdateTime,
		})
	}
	if ammusu.mutation.UserxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.UserxTable,
			Columns: []string{asmarketmenuusersort.UserxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ammusu.mutation.UserxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.UserxTable,
			Columns: []string{asmarketmenuusersort.UserxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ammusu.mutation.AppmenuxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.AppmenuxTable,
			Columns: []string{asmarketmenuusersort.AppmenuxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ammusu.mutation.AppmenuxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.AppmenuxTable,
			Columns: []string{asmarketmenuusersort.AppmenuxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ammusu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketmenuusersort.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsMarketMenuUserSortUpdateOne is the builder for updating a single AsMarketMenuUserSort entity.
type AsMarketMenuUserSortUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsMarketMenuUserSortMutation
}

// SetUserID sets the "user_id" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetUserID(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.SetUserID(i)
	return ammusuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableUserID(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetUserID(*i)
	}
	return ammusuo
}

// ClearUserID clears the value of the "user_id" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearUserID() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearUserID()
	return ammusuo
}

// SetMenuID sets the "menu_id" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetMenuID(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.SetMenuID(i)
	return ammusuo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableMenuID(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetMenuID(*i)
	}
	return ammusuo
}

// ClearMenuID clears the value of the "menu_id" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearMenuID() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearMenuID()
	return ammusuo
}

// SetSort sets the "sort" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetSort(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ResetSort()
	ammusuo.mutation.SetSort(i)
	return ammusuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableSort(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetSort(*i)
	}
	return ammusuo
}

// AddSort adds i to the "sort" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) AddSort(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.AddSort(i)
	return ammusuo
}

// ClearSort clears the value of the "sort" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearSort() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearSort()
	return ammusuo
}

// SetIsDeleted sets the "is_deleted" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetIsDeleted(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ResetIsDeleted()
	ammusuo.mutation.SetIsDeleted(i)
	return ammusuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableIsDeleted(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetIsDeleted(*i)
	}
	return ammusuo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) AddIsDeleted(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.AddIsDeleted(i)
	return ammusuo
}

// SetStatus sets the "status" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetStatus(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ResetStatus()
	ammusuo.mutation.SetStatus(i)
	return ammusuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableStatus(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetStatus(*i)
	}
	return ammusuo
}

// AddStatus adds i to the "status" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) AddStatus(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.AddStatus(i)
	return ammusuo
}

// ClearStatus clears the value of the "status" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearStatus() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearStatus()
	return ammusuo
}

// SetCreateUser sets the "create_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetCreateUser(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ResetCreateUser()
	ammusuo.mutation.SetCreateUser(i)
	return ammusuo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableCreateUser(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetCreateUser(*i)
	}
	return ammusuo
}

// AddCreateUser adds i to the "create_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) AddCreateUser(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.AddCreateUser(i)
	return ammusuo
}

// ClearCreateUser clears the value of the "create_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearCreateUser() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearCreateUser()
	return ammusuo
}

// SetUpdateUser sets the "update_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetUpdateUser(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ResetUpdateUser()
	ammusuo.mutation.SetUpdateUser(i)
	return ammusuo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableUpdateUser(i *int64) *AsMarketMenuUserSortUpdateOne {
	if i != nil {
		ammusuo.SetUpdateUser(*i)
	}
	return ammusuo
}

// AddUpdateUser adds i to the "update_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) AddUpdateUser(i int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.AddUpdateUser(i)
	return ammusuo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearUpdateUser() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearUpdateUser()
	return ammusuo
}

// SetUpdateTime sets the "update_time" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetUpdateTime(dt date.DateTime) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.SetUpdateTime(dt)
	return ammusuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearUpdateTime() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearUpdateTime()
	return ammusuo
}

// SetUserxID sets the "userx" edge to the AsUser entity by ID.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetUserxID(id int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.SetUserxID(id)
	return ammusuo
}

// SetNillableUserxID sets the "userx" edge to the AsUser entity by ID if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableUserxID(id *int64) *AsMarketMenuUserSortUpdateOne {
	if id != nil {
		ammusuo = ammusuo.SetUserxID(*id)
	}
	return ammusuo
}

// SetUserx sets the "userx" edge to the AsUser entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetUserx(a *AsUser) *AsMarketMenuUserSortUpdateOne {
	return ammusuo.SetUserxID(a.ID)
}

// SetAppmenuxID sets the "appmenux" edge to the AsMarketMenu entity by ID.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetAppmenuxID(id int64) *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.SetAppmenuxID(id)
	return ammusuo
}

// SetNillableAppmenuxID sets the "appmenux" edge to the AsMarketMenu entity by ID if the given value is not nil.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetNillableAppmenuxID(id *int64) *AsMarketMenuUserSortUpdateOne {
	if id != nil {
		ammusuo = ammusuo.SetAppmenuxID(*id)
	}
	return ammusuo
}

// SetAppmenux sets the "appmenux" edge to the AsMarketMenu entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SetAppmenux(a *AsMarketMenu) *AsMarketMenuUserSortUpdateOne {
	return ammusuo.SetAppmenuxID(a.ID)
}

// Mutation returns the AsMarketMenuUserSortMutation object of the builder.
func (ammusuo *AsMarketMenuUserSortUpdateOne) Mutation() *AsMarketMenuUserSortMutation {
	return ammusuo.mutation
}

// ClearUserx clears the "userx" edge to the AsUser entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearUserx() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearUserx()
	return ammusuo
}

// ClearAppmenux clears the "appmenux" edge to the AsMarketMenu entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ClearAppmenux() *AsMarketMenuUserSortUpdateOne {
	ammusuo.mutation.ClearAppmenux()
	return ammusuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ammusuo *AsMarketMenuUserSortUpdateOne) Select(field string, fields ...string) *AsMarketMenuUserSortUpdateOne {
	ammusuo.fields = append([]string{field}, fields...)
	return ammusuo
}

// Save executes the query and returns the updated AsMarketMenuUserSort entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) Save(ctx context.Context) (*AsMarketMenuUserSort, error) {
	var (
		err  error
		node *AsMarketMenuUserSort
	)
	ammusuo.defaults()
	if len(ammusuo.hooks) == 0 {
		node, err = ammusuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketMenuUserSortMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ammusuo.mutation = mutation
			node, err = ammusuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ammusuo.hooks) - 1; i >= 0; i-- {
			if ammusuo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = ammusuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ammusuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ammusuo *AsMarketMenuUserSortUpdateOne) SaveX(ctx context.Context) *AsMarketMenuUserSort {
	node, err := ammusuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ammusuo *AsMarketMenuUserSortUpdateOne) Exec(ctx context.Context) error {
	_, err := ammusuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ammusuo *AsMarketMenuUserSortUpdateOne) ExecX(ctx context.Context) {
	if err := ammusuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ammusuo *AsMarketMenuUserSortUpdateOne) defaults() {
	if _, ok := ammusuo.mutation.UpdateTime(); !ok && !ammusuo.mutation.UpdateTimeCleared() {
		v := asmarketmenuusersort.UpdateDefaultUpdateTime()
		ammusuo.mutation.SetUpdateTime(v)
	}
}

func (ammusuo *AsMarketMenuUserSortUpdateOne) sqlSave(ctx context.Context) (_node *AsMarketMenuUserSort, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketmenuusersort.Table,
			Columns: asmarketmenuusersort.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketmenuusersort.FieldID,
			},
		},
	}
	id, ok := ammusuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsMarketMenuUserSort.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ammusuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketmenuusersort.FieldID)
		for _, f := range fields {
			if !asmarketmenuusersort.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asmarketmenuusersort.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ammusuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ammusuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if value, ok := ammusuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if ammusuo.mutation.SortCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldSort,
		})
	}
	if value, ok := ammusuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldIsDeleted,
		})
	}
	if value, ok := ammusuo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldIsDeleted,
		})
	}
	if value, ok := ammusuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if value, ok := ammusuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if ammusuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldStatus,
		})
	}
	if value, ok := ammusuo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if value, ok := ammusuo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if ammusuo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldCreateUser,
		})
	}
	if value, ok := ammusuo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if value, ok := ammusuo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if ammusuo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketmenuusersort.FieldUpdateUser,
		})
	}
	if ammusuo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketmenuusersort.FieldCreateTime,
		})
	}
	if value, ok := ammusuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketmenuusersort.FieldUpdateTime,
		})
	}
	if ammusuo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketmenuusersort.FieldUpdateTime,
		})
	}
	if ammusuo.mutation.UserxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.UserxTable,
			Columns: []string{asmarketmenuusersort.UserxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ammusuo.mutation.UserxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.UserxTable,
			Columns: []string{asmarketmenuusersort.UserxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ammusuo.mutation.AppmenuxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.AppmenuxTable,
			Columns: []string{asmarketmenuusersort.AppmenuxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketmenu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ammusuo.mutation.AppmenuxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketmenuusersort.AppmenuxTable,
			Columns: []string{asmarketmenuusersort.AppmenuxColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketmenu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsMarketMenuUserSort{config: ammusuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ammusuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketmenuusersort.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
