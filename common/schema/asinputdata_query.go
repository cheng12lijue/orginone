// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asinputdata"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsInputDataQuery is the builder for querying AsInputData entities.
type AsInputDataQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsInputData
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsInputDataQuery builder.
func (aidq *AsInputDataQuery) Where(ps ...predicate.AsInputData) *AsInputDataQuery {
	aidq.predicates = append(aidq.predicates, ps...)
	return aidq
}

// Limit adds a limit step to the query.
func (aidq *AsInputDataQuery) Limit(limit int) *AsInputDataQuery {
	aidq.limit = &limit
	return aidq
}

// Offset adds an offset step to the query.
func (aidq *AsInputDataQuery) Offset(offset int) *AsInputDataQuery {
	aidq.offset = &offset
	return aidq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aidq *AsInputDataQuery) Unique(unique bool) *AsInputDataQuery {
	aidq.unique = &unique
	return aidq
}

// Order adds an order step to the query.
func (aidq *AsInputDataQuery) Order(o ...OrderFunc) *AsInputDataQuery {
	aidq.order = append(aidq.order, o...)
	return aidq
}

// First returns the first AsInputData entity from the query.
// Returns a *NotFoundError when no AsInputData was found.
func (aidq *AsInputDataQuery) First(ctx context.Context) (*AsInputData, error) {
	nodes, err := aidq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asinputdata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aidq *AsInputDataQuery) FirstX(ctx context.Context) *AsInputData {
	node, err := aidq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsInputData ID from the query.
// Returns a *NotFoundError when no AsInputData ID was found.
func (aidq *AsInputDataQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aidq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asinputdata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aidq *AsInputDataQuery) FirstIDX(ctx context.Context) int64 {
	id, err := aidq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsInputData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsInputData entity is found.
// Returns a *NotFoundError when no AsInputData entities are found.
func (aidq *AsInputDataQuery) Only(ctx context.Context) (*AsInputData, error) {
	nodes, err := aidq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asinputdata.Label}
	default:
		return nil, &NotSingularError{asinputdata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aidq *AsInputDataQuery) OnlyX(ctx context.Context) *AsInputData {
	node, err := aidq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsInputData ID in the query.
// Returns a *NotSingularError when more than one AsInputData ID is found.
// Returns a *NotFoundError when no entities are found.
func (aidq *AsInputDataQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = aidq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = &NotSingularError{asinputdata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aidq *AsInputDataQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := aidq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsInputDataSlice.
func (aidq *AsInputDataQuery) All(ctx context.Context) ([]*AsInputData, error) {
	if err := aidq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aidq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aidq *AsInputDataQuery) AllX(ctx context.Context) []*AsInputData {
	nodes, err := aidq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsInputData IDs.
func (aidq *AsInputDataQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := aidq.Select(asinputdata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aidq *AsInputDataQuery) IDsX(ctx context.Context) []int64 {
	ids, err := aidq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aidq *AsInputDataQuery) Count(ctx context.Context) (int64, error) {
	if err := aidq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aidq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aidq *AsInputDataQuery) CountX(ctx context.Context) int64 {
	count, err := aidq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aidq *AsInputDataQuery) Exist(ctx context.Context) (bool, error) {
	if err := aidq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aidq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aidq *AsInputDataQuery) ExistX(ctx context.Context) bool {
	exist, err := aidq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsInputDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aidq *AsInputDataQuery) Clone() *AsInputDataQuery {
	if aidq == nil {
		return nil
	}
	return &AsInputDataQuery{
		config:     aidq.config,
		limit:      aidq.limit,
		offset:     aidq.offset,
		order:      append([]OrderFunc{}, aidq.order...),
		predicates: append([]predicate.AsInputData{}, aidq.predicates...),
		// clone intermediate query.
		sql:    aidq.sql.Clone(),
		path:   aidq.path,
		unique: aidq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FileID int64 `json:"fileId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsInputData.Query().
//		GroupBy(asinputdata.FieldFileID).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (aidq *AsInputDataQuery) GroupBy(field string, fields ...string) *AsInputDataGroupBy {
	group := &AsInputDataGroupBy{config: aidq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aidq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aidq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FileID int64 `json:"fileId"`
//	}
//
//	client.AsInputData.Query().
//		Select(asinputdata.FieldFileID).
//		Scan(ctx, &v)
//
func (aidq *AsInputDataQuery) Select(fields ...string) *AsInputDataSelect {
	aidq.fields = append(aidq.fields, fields...)
	return &AsInputDataSelect{AsInputDataQuery: aidq}
}

func (aidq *AsInputDataQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aidq.fields {
		if !asinputdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if aidq.path != nil {
		prev, err := aidq.path(ctx)
		if err != nil {
			return err
		}
		aidq.sql = prev
	}
	return nil
}

func (aidq *AsInputDataQuery) sqlAll(ctx context.Context) ([]*AsInputData, error) {
	var (
		nodes = []*AsInputData{}
		_spec = aidq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsInputData{config: aidq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, aidq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aidq *AsInputDataQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := aidq.querySpec()
	_spec.Node.Columns = aidq.fields
	if len(aidq.fields) > 0 {
		_spec.Unique = aidq.unique != nil && *aidq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, aidq.driver, _spec)
	return int64(c), err
}

func (aidq *AsInputDataQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aidq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (aidq *AsInputDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asinputdata.Table,
			Columns: asinputdata.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asinputdata.FieldID,
			},
		},
		From:   aidq.sql,
		Unique: true,
	}
	if unique := aidq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aidq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asinputdata.FieldID)
		for i := range fields {
			if fields[i] != asinputdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aidq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aidq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aidq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aidq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aidq *AsInputDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aidq.driver.Dialect())
	t1 := builder.Table(asinputdata.Table)
	columns := aidq.fields
	if len(columns) == 0 {
		columns = asinputdata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aidq.sql != nil {
		selector = aidq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aidq.unique != nil && *aidq.unique {
		selector.Distinct()
	}
	for _, p := range aidq.predicates {
		p(selector)
	}
	for _, p := range aidq.order {
		p(selector)
	}
	if offset := aidq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aidq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsInputDataGroupBy is the group-by builder for AsInputData entities.
type AsInputDataGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aidgb *AsInputDataGroupBy) Aggregate(fns ...AggregateFunc) *AsInputDataGroupBy {
	aidgb.fns = append(aidgb.fns, fns...)
	return aidgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aidgb *AsInputDataGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aidgb.path(ctx)
	if err != nil {
		return err
	}
	aidgb.sql = query
	return aidgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aidgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aidgb.fields) > 1 {
		return nil, errors.New("schema: AsInputDataGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aidgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) StringsX(ctx context.Context) []string {
	v, err := aidgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aidgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) StringX(ctx context.Context) string {
	v, err := aidgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aidgb.fields) > 1 {
		return nil, errors.New("schema: AsInputDataGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aidgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) IntsX(ctx context.Context) []int {
	v, err := aidgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aidgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) IntX(ctx context.Context) int {
	v, err := aidgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aidgb.fields) > 1 {
		return nil, errors.New("schema: AsInputDataGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aidgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aidgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aidgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aidgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aidgb.fields) > 1 {
		return nil, errors.New("schema: AsInputDataGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aidgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aidgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aidgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) BoolX(ctx context.Context) bool {
	v, err := aidgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(aidgb.fields) > 1 {
		return nil, errors.New("schema: AsInputDataGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := aidgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := aidgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aidgb *AsInputDataGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = aidgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (aidgb *AsInputDataGroupBy) Int64X(ctx context.Context) int64 {
	v, err := aidgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aidgb *AsInputDataGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aidgb.fields {
		if !asinputdata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aidgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aidgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aidgb *AsInputDataGroupBy) sqlQuery() *sql.Selector {
	selector := aidgb.sql.Select()
	aggregation := make([]string, 0, len(aidgb.fns))
	for _, fn := range aidgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aidgb.fields)+len(aidgb.fns))
		for _, f := range aidgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aidgb.fields...)...)
}

// AsInputDataSelect is the builder for selecting fields of AsInputData entities.
type AsInputDataSelect struct {
	*AsInputDataQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aids *AsInputDataSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aids.prepareQuery(ctx); err != nil {
		return err
	}
	aids.sql = aids.AsInputDataQuery.sqlQuery(ctx)
	return aids.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aids *AsInputDataSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aids.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aids.fields) > 1 {
		return nil, errors.New("schema: AsInputDataSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aids *AsInputDataSelect) StringsX(ctx context.Context) []string {
	v, err := aids.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aids.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aids *AsInputDataSelect) StringX(ctx context.Context) string {
	v, err := aids.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aids.fields) > 1 {
		return nil, errors.New("schema: AsInputDataSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aids *AsInputDataSelect) IntsX(ctx context.Context) []int {
	v, err := aids.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aids.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aids *AsInputDataSelect) IntX(ctx context.Context) int {
	v, err := aids.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aids.fields) > 1 {
		return nil, errors.New("schema: AsInputDataSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aids *AsInputDataSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aids.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aids.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aids *AsInputDataSelect) Float64X(ctx context.Context) float64 {
	v, err := aids.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aids.fields) > 1 {
		return nil, errors.New("schema: AsInputDataSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aids *AsInputDataSelect) BoolsX(ctx context.Context) []bool {
	v, err := aids.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aids.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aids *AsInputDataSelect) BoolX(ctx context.Context) bool {
	v, err := aids.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(aids.fields) > 1 {
		return nil, errors.New("schema: AsInputDataSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := aids.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (aids *AsInputDataSelect) Int64sX(ctx context.Context) []int64 {
	v, err := aids.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (aids *AsInputDataSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = aids.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asinputdata.Label}
	default:
		err = fmt.Errorf("schema: AsInputDataSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (aids *AsInputDataSelect) Int64X(ctx context.Context) int64 {
	v, err := aids.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aids *AsInputDataSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aids.sql.Query()
	if err := aids.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}