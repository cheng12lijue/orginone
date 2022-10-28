// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"orginone/common/schema/asmarketapprole"
	"orginone/common/schema/asmarketmenu"
	"orginone/common/schema/asmarketrolemenu"
	"orginone/common/schema/predicate"
	"orginone/common/tools/date"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketRoleMenuUpdate is the builder for updating AsMarketRoleMenu entities.
type AsMarketRoleMenuUpdate struct {
	config
	hooks    []Hook
	mutation *AsMarketRoleMenuMutation
}

// Where appends a list predicates to the AsMarketRoleMenuUpdate builder.
func (amrmu *AsMarketRoleMenuUpdate) Where(ps ...predicate.AsMarketRoleMenu) *AsMarketRoleMenuUpdate {
	amrmu.mutation.Where(ps...)
	return amrmu
}

// SetRoleID sets the "role_id" field.
func (amrmu *AsMarketRoleMenuUpdate) SetRoleID(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.SetRoleID(i)
	return amrmu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableRoleID(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetRoleID(*i)
	}
	return amrmu
}

// ClearRoleID clears the value of the "role_id" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearRoleID() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearRoleID()
	return amrmu
}

// SetMenuID sets the "menu_id" field.
func (amrmu *AsMarketRoleMenuUpdate) SetMenuID(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.SetMenuID(i)
	return amrmu
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableMenuID(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetMenuID(*i)
	}
	return amrmu
}

// ClearMenuID clears the value of the "menu_id" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearMenuID() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearMenuID()
	return amrmu
}

// SetIsDeleted sets the "is_deleted" field.
func (amrmu *AsMarketRoleMenuUpdate) SetIsDeleted(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.ResetIsDeleted()
	amrmu.mutation.SetIsDeleted(i)
	return amrmu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableIsDeleted(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetIsDeleted(*i)
	}
	return amrmu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (amrmu *AsMarketRoleMenuUpdate) AddIsDeleted(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.AddIsDeleted(i)
	return amrmu
}

// SetStatus sets the "status" field.
func (amrmu *AsMarketRoleMenuUpdate) SetStatus(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.ResetStatus()
	amrmu.mutation.SetStatus(i)
	return amrmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableStatus(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetStatus(*i)
	}
	return amrmu
}

// AddStatus adds i to the "status" field.
func (amrmu *AsMarketRoleMenuUpdate) AddStatus(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.AddStatus(i)
	return amrmu
}

// ClearStatus clears the value of the "status" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearStatus() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearStatus()
	return amrmu
}

// SetCreateUser sets the "create_user" field.
func (amrmu *AsMarketRoleMenuUpdate) SetCreateUser(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.ResetCreateUser()
	amrmu.mutation.SetCreateUser(i)
	return amrmu
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableCreateUser(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetCreateUser(*i)
	}
	return amrmu
}

// AddCreateUser adds i to the "create_user" field.
func (amrmu *AsMarketRoleMenuUpdate) AddCreateUser(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.AddCreateUser(i)
	return amrmu
}

// ClearCreateUser clears the value of the "create_user" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearCreateUser() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearCreateUser()
	return amrmu
}

// SetUpdateUser sets the "update_user" field.
func (amrmu *AsMarketRoleMenuUpdate) SetUpdateUser(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.ResetUpdateUser()
	amrmu.mutation.SetUpdateUser(i)
	return amrmu
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableUpdateUser(i *int64) *AsMarketRoleMenuUpdate {
	if i != nil {
		amrmu.SetUpdateUser(*i)
	}
	return amrmu
}

// AddUpdateUser adds i to the "update_user" field.
func (amrmu *AsMarketRoleMenuUpdate) AddUpdateUser(i int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.AddUpdateUser(i)
	return amrmu
}

// ClearUpdateUser clears the value of the "update_user" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearUpdateUser() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearUpdateUser()
	return amrmu
}

// SetUpdateTime sets the "update_time" field.
func (amrmu *AsMarketRoleMenuUpdate) SetUpdateTime(dt date.DateTime) *AsMarketRoleMenuUpdate {
	amrmu.mutation.SetUpdateTime(dt)
	return amrmu
}

// ClearUpdateTime clears the value of the "update_time" field.
func (amrmu *AsMarketRoleMenuUpdate) ClearUpdateTime() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearUpdateTime()
	return amrmu
}

// SetMenuxID sets the "menux" edge to the AsMarketMenu entity by ID.
func (amrmu *AsMarketRoleMenuUpdate) SetMenuxID(id int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.SetMenuxID(id)
	return amrmu
}

// SetNillableMenuxID sets the "menux" edge to the AsMarketMenu entity by ID if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableMenuxID(id *int64) *AsMarketRoleMenuUpdate {
	if id != nil {
		amrmu = amrmu.SetMenuxID(*id)
	}
	return amrmu
}

// SetMenux sets the "menux" edge to the AsMarketMenu entity.
func (amrmu *AsMarketRoleMenuUpdate) SetMenux(a *AsMarketMenu) *AsMarketRoleMenuUpdate {
	return amrmu.SetMenuxID(a.ID)
}

// SetRolexID sets the "rolex" edge to the AsMarketAppRole entity by ID.
func (amrmu *AsMarketRoleMenuUpdate) SetRolexID(id int64) *AsMarketRoleMenuUpdate {
	amrmu.mutation.SetRolexID(id)
	return amrmu
}

// SetNillableRolexID sets the "rolex" edge to the AsMarketAppRole entity by ID if the given value is not nil.
func (amrmu *AsMarketRoleMenuUpdate) SetNillableRolexID(id *int64) *AsMarketRoleMenuUpdate {
	if id != nil {
		amrmu = amrmu.SetRolexID(*id)
	}
	return amrmu
}

// SetRolex sets the "rolex" edge to the AsMarketAppRole entity.
func (amrmu *AsMarketRoleMenuUpdate) SetRolex(a *AsMarketAppRole) *AsMarketRoleMenuUpdate {
	return amrmu.SetRolexID(a.ID)
}

// Mutation returns the AsMarketRoleMenuMutation object of the builder.
func (amrmu *AsMarketRoleMenuUpdate) Mutation() *AsMarketRoleMenuMutation {
	return amrmu.mutation
}

// ClearMenux clears the "menux" edge to the AsMarketMenu entity.
func (amrmu *AsMarketRoleMenuUpdate) ClearMenux() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearMenux()
	return amrmu
}

// ClearRolex clears the "rolex" edge to the AsMarketAppRole entity.
func (amrmu *AsMarketRoleMenuUpdate) ClearRolex() *AsMarketRoleMenuUpdate {
	amrmu.mutation.ClearRolex()
	return amrmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (amrmu *AsMarketRoleMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	amrmu.defaults()
	if len(amrmu.hooks) == 0 {
		affected, err = amrmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amrmu.mutation = mutation
			affected, err = amrmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amrmu.hooks) - 1; i >= 0; i-- {
			if amrmu.hooks[i] == nil {
				return 0, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amrmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amrmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (amrmu *AsMarketRoleMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := amrmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (amrmu *AsMarketRoleMenuUpdate) Exec(ctx context.Context) error {
	_, err := amrmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amrmu *AsMarketRoleMenuUpdate) ExecX(ctx context.Context) {
	if err := amrmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amrmu *AsMarketRoleMenuUpdate) defaults() {
	if _, ok := amrmu.mutation.UpdateTime(); !ok && !amrmu.mutation.UpdateTimeCleared() {
		v := asmarketrolemenu.UpdateDefaultUpdateTime()
		amrmu.mutation.SetUpdateTime(v)
	}
}

func (amrmu *AsMarketRoleMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketrolemenu.Table,
			Columns: asmarketrolemenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketrolemenu.FieldID,
			},
		},
	}
	if ps := amrmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amrmu.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldIsDeleted,
		})
	}
	if value, ok := amrmu.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldIsDeleted,
		})
	}
	if value, ok := amrmu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if value, ok := amrmu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if amrmu.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if value, ok := amrmu.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if value, ok := amrmu.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if amrmu.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if value, ok := amrmu.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if value, ok := amrmu.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if amrmu.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if amrmu.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketrolemenu.FieldCreateTime,
		})
	}
	if value, ok := amrmu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateTime,
		})
	}
	if amrmu.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketrolemenu.FieldUpdateTime,
		})
	}
	if amrmu.mutation.MenuxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.MenuxTable,
			Columns: []string{asmarketrolemenu.MenuxColumn},
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
	if nodes := amrmu.mutation.MenuxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.MenuxTable,
			Columns: []string{asmarketrolemenu.MenuxColumn},
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
	if amrmu.mutation.RolexCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.RolexTable,
			Columns: []string{asmarketrolemenu.RolexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapprole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := amrmu.mutation.RolexIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.RolexTable,
			Columns: []string{asmarketrolemenu.RolexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, amrmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AsMarketRoleMenuUpdateOne is the builder for updating a single AsMarketRoleMenu entity.
type AsMarketRoleMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AsMarketRoleMenuMutation
}

// SetRoleID sets the "role_id" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetRoleID(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.SetRoleID(i)
	return amrmuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableRoleID(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetRoleID(*i)
	}
	return amrmuo
}

// ClearRoleID clears the value of the "role_id" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearRoleID() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearRoleID()
	return amrmuo
}

// SetMenuID sets the "menu_id" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetMenuID(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.SetMenuID(i)
	return amrmuo
}

// SetNillableMenuID sets the "menu_id" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableMenuID(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetMenuID(*i)
	}
	return amrmuo
}

// ClearMenuID clears the value of the "menu_id" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearMenuID() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearMenuID()
	return amrmuo
}

// SetIsDeleted sets the "is_deleted" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetIsDeleted(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ResetIsDeleted()
	amrmuo.mutation.SetIsDeleted(i)
	return amrmuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableIsDeleted(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetIsDeleted(*i)
	}
	return amrmuo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) AddIsDeleted(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.AddIsDeleted(i)
	return amrmuo
}

// SetStatus sets the "status" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetStatus(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ResetStatus()
	amrmuo.mutation.SetStatus(i)
	return amrmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableStatus(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetStatus(*i)
	}
	return amrmuo
}

// AddStatus adds i to the "status" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) AddStatus(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.AddStatus(i)
	return amrmuo
}

// ClearStatus clears the value of the "status" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearStatus() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearStatus()
	return amrmuo
}

// SetCreateUser sets the "create_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetCreateUser(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ResetCreateUser()
	amrmuo.mutation.SetCreateUser(i)
	return amrmuo
}

// SetNillableCreateUser sets the "create_user" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableCreateUser(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetCreateUser(*i)
	}
	return amrmuo
}

// AddCreateUser adds i to the "create_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) AddCreateUser(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.AddCreateUser(i)
	return amrmuo
}

// ClearCreateUser clears the value of the "create_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearCreateUser() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearCreateUser()
	return amrmuo
}

// SetUpdateUser sets the "update_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetUpdateUser(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ResetUpdateUser()
	amrmuo.mutation.SetUpdateUser(i)
	return amrmuo
}

// SetNillableUpdateUser sets the "update_user" field if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableUpdateUser(i *int64) *AsMarketRoleMenuUpdateOne {
	if i != nil {
		amrmuo.SetUpdateUser(*i)
	}
	return amrmuo
}

// AddUpdateUser adds i to the "update_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) AddUpdateUser(i int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.AddUpdateUser(i)
	return amrmuo
}

// ClearUpdateUser clears the value of the "update_user" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearUpdateUser() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearUpdateUser()
	return amrmuo
}

// SetUpdateTime sets the "update_time" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetUpdateTime(dt date.DateTime) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.SetUpdateTime(dt)
	return amrmuo
}

// ClearUpdateTime clears the value of the "update_time" field.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearUpdateTime() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearUpdateTime()
	return amrmuo
}

// SetMenuxID sets the "menux" edge to the AsMarketMenu entity by ID.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetMenuxID(id int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.SetMenuxID(id)
	return amrmuo
}

// SetNillableMenuxID sets the "menux" edge to the AsMarketMenu entity by ID if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableMenuxID(id *int64) *AsMarketRoleMenuUpdateOne {
	if id != nil {
		amrmuo = amrmuo.SetMenuxID(*id)
	}
	return amrmuo
}

// SetMenux sets the "menux" edge to the AsMarketMenu entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetMenux(a *AsMarketMenu) *AsMarketRoleMenuUpdateOne {
	return amrmuo.SetMenuxID(a.ID)
}

// SetRolexID sets the "rolex" edge to the AsMarketAppRole entity by ID.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetRolexID(id int64) *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.SetRolexID(id)
	return amrmuo
}

// SetNillableRolexID sets the "rolex" edge to the AsMarketAppRole entity by ID if the given value is not nil.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetNillableRolexID(id *int64) *AsMarketRoleMenuUpdateOne {
	if id != nil {
		amrmuo = amrmuo.SetRolexID(*id)
	}
	return amrmuo
}

// SetRolex sets the "rolex" edge to the AsMarketAppRole entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) SetRolex(a *AsMarketAppRole) *AsMarketRoleMenuUpdateOne {
	return amrmuo.SetRolexID(a.ID)
}

// Mutation returns the AsMarketRoleMenuMutation object of the builder.
func (amrmuo *AsMarketRoleMenuUpdateOne) Mutation() *AsMarketRoleMenuMutation {
	return amrmuo.mutation
}

// ClearMenux clears the "menux" edge to the AsMarketMenu entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearMenux() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearMenux()
	return amrmuo
}

// ClearRolex clears the "rolex" edge to the AsMarketAppRole entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) ClearRolex() *AsMarketRoleMenuUpdateOne {
	amrmuo.mutation.ClearRolex()
	return amrmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (amrmuo *AsMarketRoleMenuUpdateOne) Select(field string, fields ...string) *AsMarketRoleMenuUpdateOne {
	amrmuo.fields = append([]string{field}, fields...)
	return amrmuo
}

// Save executes the query and returns the updated AsMarketRoleMenu entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) Save(ctx context.Context) (*AsMarketRoleMenu, error) {
	var (
		err  error
		node *AsMarketRoleMenu
	)
	amrmuo.defaults()
	if len(amrmuo.hooks) == 0 {
		node, err = amrmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AsMarketRoleMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amrmuo.mutation = mutation
			node, err = amrmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(amrmuo.hooks) - 1; i >= 0; i-- {
			if amrmuo.hooks[i] == nil {
				return nil, fmt.Errorf("schema: uninitialized hook (forgotten import schema/runtime?)")
			}
			mut = amrmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amrmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (amrmuo *AsMarketRoleMenuUpdateOne) SaveX(ctx context.Context) *AsMarketRoleMenu {
	node, err := amrmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (amrmuo *AsMarketRoleMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := amrmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amrmuo *AsMarketRoleMenuUpdateOne) ExecX(ctx context.Context) {
	if err := amrmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amrmuo *AsMarketRoleMenuUpdateOne) defaults() {
	if _, ok := amrmuo.mutation.UpdateTime(); !ok && !amrmuo.mutation.UpdateTimeCleared() {
		v := asmarketrolemenu.UpdateDefaultUpdateTime()
		amrmuo.mutation.SetUpdateTime(v)
	}
}

func (amrmuo *AsMarketRoleMenuUpdateOne) sqlSave(ctx context.Context) (_node *AsMarketRoleMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketrolemenu.Table,
			Columns: asmarketrolemenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketrolemenu.FieldID,
			},
		},
	}
	id, ok := amrmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`schema: missing "AsMarketRoleMenu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := amrmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketrolemenu.FieldID)
		for _, f := range fields {
			if !asmarketrolemenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
			}
			if f != asmarketrolemenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := amrmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amrmuo.mutation.IsDeleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldIsDeleted,
		})
	}
	if value, ok := amrmuo.mutation.AddedIsDeleted(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldIsDeleted,
		})
	}
	if value, ok := amrmuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if value, ok := amrmuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if amrmuo.mutation.StatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldStatus,
		})
	}
	if value, ok := amrmuo.mutation.CreateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if value, ok := amrmuo.mutation.AddedCreateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if amrmuo.mutation.CreateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldCreateUser,
		})
	}
	if value, ok := amrmuo.mutation.UpdateUser(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if value, ok := amrmuo.mutation.AddedUpdateUser(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if amrmuo.mutation.UpdateUserCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: asmarketrolemenu.FieldUpdateUser,
		})
	}
	if amrmuo.mutation.CreateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketrolemenu.FieldCreateTime,
		})
	}
	if value, ok := amrmuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asmarketrolemenu.FieldUpdateTime,
		})
	}
	if amrmuo.mutation.UpdateTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: asmarketrolemenu.FieldUpdateTime,
		})
	}
	if amrmuo.mutation.MenuxCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.MenuxTable,
			Columns: []string{asmarketrolemenu.MenuxColumn},
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
	if nodes := amrmuo.mutation.MenuxIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.MenuxTable,
			Columns: []string{asmarketrolemenu.MenuxColumn},
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
	if amrmuo.mutation.RolexCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.RolexTable,
			Columns: []string{asmarketrolemenu.RolexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapprole.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := amrmuo.mutation.RolexIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asmarketrolemenu.RolexTable,
			Columns: []string{asmarketrolemenu.RolexColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: asmarketapprole.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AsMarketRoleMenu{config: amrmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, amrmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{asmarketrolemenu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}