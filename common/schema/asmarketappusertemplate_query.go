// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asmarketappcomponenttemplate"
	"orginone/common/schema/asmarketappusertemplate"
	"orginone/common/schema/asuser"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsMarketAppUserTemplateQuery is the builder for querying AsMarketAppUserTemplate entities.
type AsMarketAppUserTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsMarketAppUserTemplate
	// eager-loading edges.
	withUserx     *AsUserQuery
	withTemplatex *AsMarketAppComponentTemplateQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsMarketAppUserTemplateQuery builder.
func (amautq *AsMarketAppUserTemplateQuery) Where(ps ...predicate.AsMarketAppUserTemplate) *AsMarketAppUserTemplateQuery {
	amautq.predicates = append(amautq.predicates, ps...)
	return amautq
}

// Limit adds a limit step to the query.
func (amautq *AsMarketAppUserTemplateQuery) Limit(limit int) *AsMarketAppUserTemplateQuery {
	amautq.limit = &limit
	return amautq
}

// Offset adds an offset step to the query.
func (amautq *AsMarketAppUserTemplateQuery) Offset(offset int) *AsMarketAppUserTemplateQuery {
	amautq.offset = &offset
	return amautq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (amautq *AsMarketAppUserTemplateQuery) Unique(unique bool) *AsMarketAppUserTemplateQuery {
	amautq.unique = &unique
	return amautq
}

// Order adds an order step to the query.
func (amautq *AsMarketAppUserTemplateQuery) Order(o ...OrderFunc) *AsMarketAppUserTemplateQuery {
	amautq.order = append(amautq.order, o...)
	return amautq
}

// QueryUserx chains the current query on the "userx" edge.
func (amautq *AsMarketAppUserTemplateQuery) QueryUserx() *AsUserQuery {
	query := &AsUserQuery{config: amautq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amautq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amautq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asmarketappusertemplate.Table, asmarketappusertemplate.FieldID, selector),
			sqlgraph.To(asuser.Table, asuser.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asmarketappusertemplate.UserxTable, asmarketappusertemplate.UserxColumn),
		)
		fromU = sqlgraph.SetNeighbors(amautq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTemplatex chains the current query on the "templatex" edge.
func (amautq *AsMarketAppUserTemplateQuery) QueryTemplatex() *AsMarketAppComponentTemplateQuery {
	query := &AsMarketAppComponentTemplateQuery{config: amautq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := amautq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := amautq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asmarketappusertemplate.Table, asmarketappusertemplate.FieldID, selector),
			sqlgraph.To(asmarketappcomponenttemplate.Table, asmarketappcomponenttemplate.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asmarketappusertemplate.TemplatexTable, asmarketappusertemplate.TemplatexColumn),
		)
		fromU = sqlgraph.SetNeighbors(amautq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsMarketAppUserTemplate entity from the query.
// Returns a *NotFoundError when no AsMarketAppUserTemplate was found.
func (amautq *AsMarketAppUserTemplateQuery) First(ctx context.Context) (*AsMarketAppUserTemplate, error) {
	nodes, err := amautq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asmarketappusertemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) FirstX(ctx context.Context) *AsMarketAppUserTemplate {
	node, err := amautq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsMarketAppUserTemplate ID from the query.
// Returns a *NotFoundError when no AsMarketAppUserTemplate ID was found.
func (amautq *AsMarketAppUserTemplateQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amautq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asmarketappusertemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) FirstIDX(ctx context.Context) int64 {
	id, err := amautq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsMarketAppUserTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsMarketAppUserTemplate entity is found.
// Returns a *NotFoundError when no AsMarketAppUserTemplate entities are found.
func (amautq *AsMarketAppUserTemplateQuery) Only(ctx context.Context) (*AsMarketAppUserTemplate, error) {
	nodes, err := amautq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asmarketappusertemplate.Label}
	default:
		return nil, &NotSingularError{asmarketappusertemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) OnlyX(ctx context.Context) *AsMarketAppUserTemplate {
	node, err := amautq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsMarketAppUserTemplate ID in the query.
// Returns a *NotSingularError when more than one AsMarketAppUserTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (amautq *AsMarketAppUserTemplateQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = amautq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = &NotSingularError{asmarketappusertemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := amautq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsMarketAppUserTemplates.
func (amautq *AsMarketAppUserTemplateQuery) All(ctx context.Context) ([]*AsMarketAppUserTemplate, error) {
	if err := amautq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return amautq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) AllX(ctx context.Context) []*AsMarketAppUserTemplate {
	nodes, err := amautq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsMarketAppUserTemplate IDs.
func (amautq *AsMarketAppUserTemplateQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := amautq.Select(asmarketappusertemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) IDsX(ctx context.Context) []int64 {
	ids, err := amautq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (amautq *AsMarketAppUserTemplateQuery) Count(ctx context.Context) (int64, error) {
	if err := amautq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return amautq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) CountX(ctx context.Context) int64 {
	count, err := amautq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (amautq *AsMarketAppUserTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := amautq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return amautq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (amautq *AsMarketAppUserTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := amautq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsMarketAppUserTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (amautq *AsMarketAppUserTemplateQuery) Clone() *AsMarketAppUserTemplateQuery {
	if amautq == nil {
		return nil
	}
	return &AsMarketAppUserTemplateQuery{
		config:        amautq.config,
		limit:         amautq.limit,
		offset:        amautq.offset,
		order:         append([]OrderFunc{}, amautq.order...),
		predicates:    append([]predicate.AsMarketAppUserTemplate{}, amautq.predicates...),
		withUserx:     amautq.withUserx.Clone(),
		withTemplatex: amautq.withTemplatex.Clone(),
		// clone intermediate query.
		sql:    amautq.sql.Clone(),
		path:   amautq.path,
		unique: amautq.unique,
	}
}

// WithUserx tells the query-builder to eager-load the nodes that are connected to
// the "userx" edge. The optional arguments are used to configure the query builder of the edge.
func (amautq *AsMarketAppUserTemplateQuery) WithUserx(opts ...func(*AsUserQuery)) *AsMarketAppUserTemplateQuery {
	query := &AsUserQuery{config: amautq.config}
	for _, opt := range opts {
		opt(query)
	}
	amautq.withUserx = query
	return amautq
}

// WithTemplatex tells the query-builder to eager-load the nodes that are connected to
// the "templatex" edge. The optional arguments are used to configure the query builder of the edge.
func (amautq *AsMarketAppUserTemplateQuery) WithTemplatex(opts ...func(*AsMarketAppComponentTemplateQuery)) *AsMarketAppUserTemplateQuery {
	query := &AsMarketAppComponentTemplateQuery{config: amautq.config}
	for _, opt := range opts {
		opt(query)
	}
	amautq.withTemplatex = query
	return amautq
}

// ThenUserx tells the query-builder to eager-load the nodes that are connected to
// the "userx" edge. The optional arguments are used to configure the query builder of the edge.
func (amautq *AsMarketAppUserTemplateQuery) ThenUserx(opts ...func(*AsUserQuery)) *AsMarketAppUserTemplateQuery {
	query := &AsUserQuery{config: amautq.config}
	for _, opt := range opts {
		opt(query.Where(asuser.IsDeleted(0)))
	}
	amautq.withUserx = query
	return amautq
}

// ThenTemplatex tells the query-builder to eager-load the nodes that are connected to
// the "templatex" edge. The optional arguments are used to configure the query builder of the edge.
func (amautq *AsMarketAppUserTemplateQuery) ThenTemplatex(opts ...func(*AsMarketAppComponentTemplateQuery)) *AsMarketAppUserTemplateQuery {
	query := &AsMarketAppComponentTemplateQuery{config: amautq.config}
	for _, opt := range opts {
		opt(query.Where(asmarketappcomponenttemplate.IsDeleted(0)))
	}
	amautq.withTemplatex = query
	return amautq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"userId"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsMarketAppUserTemplate.Query().
//		GroupBy(asmarketappusertemplate.FieldUserID).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (amautq *AsMarketAppUserTemplateQuery) GroupBy(field string, fields ...string) *AsMarketAppUserTemplateGroupBy {
	group := &AsMarketAppUserTemplateGroupBy{config: amautq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := amautq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return amautq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"userId"`
//	}
//
//	client.AsMarketAppUserTemplate.Query().
//		Select(asmarketappusertemplate.FieldUserID).
//		Scan(ctx, &v)
//
func (amautq *AsMarketAppUserTemplateQuery) Select(fields ...string) *AsMarketAppUserTemplateSelect {
	amautq.fields = append(amautq.fields, fields...)
	return &AsMarketAppUserTemplateSelect{AsMarketAppUserTemplateQuery: amautq}
}

func (amautq *AsMarketAppUserTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range amautq.fields {
		if !asmarketappusertemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if amautq.path != nil {
		prev, err := amautq.path(ctx)
		if err != nil {
			return err
		}
		amautq.sql = prev
	}
	return nil
}

func (amautq *AsMarketAppUserTemplateQuery) sqlAll(ctx context.Context) ([]*AsMarketAppUserTemplate, error) {
	var (
		nodes       = []*AsMarketAppUserTemplate{}
		_spec       = amautq.querySpec()
		loadedTypes = [2]bool{
			amautq.withUserx != nil,
			amautq.withTemplatex != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsMarketAppUserTemplate{config: amautq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, amautq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := amautq.withUserx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsMarketAppUserTemplate)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asuser.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Userx = n
			}
		}
	}

	if query := amautq.withTemplatex; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsMarketAppUserTemplate)
		for i := range nodes {
			fk := nodes[i].TemplateID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asmarketappcomponenttemplate.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "template_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Templatex = n
			}
		}
	}

	return nodes, nil
}

func (amautq *AsMarketAppUserTemplateQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := amautq.querySpec()
	_spec.Node.Columns = amautq.fields
	if len(amautq.fields) > 0 {
		_spec.Unique = amautq.unique != nil && *amautq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, amautq.driver, _spec)
	return int64(c), err
}

func (amautq *AsMarketAppUserTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := amautq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (amautq *AsMarketAppUserTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asmarketappusertemplate.Table,
			Columns: asmarketappusertemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asmarketappusertemplate.FieldID,
			},
		},
		From:   amautq.sql,
		Unique: true,
	}
	if unique := amautq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := amautq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asmarketappusertemplate.FieldID)
		for i := range fields {
			if fields[i] != asmarketappusertemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := amautq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := amautq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := amautq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := amautq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (amautq *AsMarketAppUserTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(amautq.driver.Dialect())
	t1 := builder.Table(asmarketappusertemplate.Table)
	columns := amautq.fields
	if len(columns) == 0 {
		columns = asmarketappusertemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if amautq.sql != nil {
		selector = amautq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if amautq.unique != nil && *amautq.unique {
		selector.Distinct()
	}
	for _, p := range amautq.predicates {
		p(selector)
	}
	for _, p := range amautq.order {
		p(selector)
	}
	if offset := amautq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := amautq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsMarketAppUserTemplateGroupBy is the group-by builder for AsMarketAppUserTemplate entities.
type AsMarketAppUserTemplateGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (amautgb *AsMarketAppUserTemplateGroupBy) Aggregate(fns ...AggregateFunc) *AsMarketAppUserTemplateGroupBy {
	amautgb.fns = append(amautgb.fns, fns...)
	return amautgb
}

// Scan applies the group-by query and scans the result into the given value.
func (amautgb *AsMarketAppUserTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := amautgb.path(ctx)
	if err != nil {
		return err
	}
	amautgb.sql = query
	return amautgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := amautgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(amautgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := amautgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) StringsX(ctx context.Context) []string {
	v, err := amautgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amautgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) StringX(ctx context.Context) string {
	v, err := amautgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(amautgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := amautgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) IntsX(ctx context.Context) []int {
	v, err := amautgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amautgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) IntX(ctx context.Context) int {
	v, err := amautgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(amautgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := amautgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := amautgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amautgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) Float64X(ctx context.Context) float64 {
	v, err := amautgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(amautgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := amautgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := amautgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amautgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) BoolX(ctx context.Context) bool {
	v, err := amautgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(amautgb.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := amautgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := amautgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (amautgb *AsMarketAppUserTemplateGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amautgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amautgb *AsMarketAppUserTemplateGroupBy) Int64X(ctx context.Context) int64 {
	v, err := amautgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amautgb *AsMarketAppUserTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range amautgb.fields {
		if !asmarketappusertemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := amautgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := amautgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amautgb *AsMarketAppUserTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := amautgb.sql.Select()
	aggregation := make([]string, 0, len(amautgb.fns))
	for _, fn := range amautgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(amautgb.fields)+len(amautgb.fns))
		for _, f := range amautgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(amautgb.fields...)...)
}

// AsMarketAppUserTemplateSelect is the builder for selecting fields of AsMarketAppUserTemplate entities.
type AsMarketAppUserTemplateSelect struct {
	*AsMarketAppUserTemplateQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (amauts *AsMarketAppUserTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := amauts.prepareQuery(ctx); err != nil {
		return err
	}
	amauts.sql = amauts.AsMarketAppUserTemplateQuery.sqlQuery(ctx)
	return amauts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) ScanX(ctx context.Context, v interface{}) {
	if err := amauts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Strings(ctx context.Context) ([]string, error) {
	if len(amauts.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := amauts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) StringsX(ctx context.Context) []string {
	v, err := amauts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amauts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) StringX(ctx context.Context) string {
	v, err := amauts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Ints(ctx context.Context) ([]int, error) {
	if len(amauts.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := amauts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) IntsX(ctx context.Context) []int {
	v, err := amauts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amauts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) IntX(ctx context.Context) int {
	v, err := amauts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(amauts.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := amauts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) Float64sX(ctx context.Context) []float64 {
	v, err := amauts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amauts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) Float64X(ctx context.Context) float64 {
	v, err := amauts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(amauts.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := amauts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) BoolsX(ctx context.Context) []bool {
	v, err := amauts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amauts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) BoolX(ctx context.Context) bool {
	v, err := amauts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(amauts.fields) > 1 {
		return nil, errors.New("schema: AsMarketAppUserTemplateSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := amauts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) Int64sX(ctx context.Context) []int64 {
	v, err := amauts.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (amauts *AsMarketAppUserTemplateSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = amauts.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asmarketappusertemplate.Label}
	default:
		err = fmt.Errorf("schema: AsMarketAppUserTemplateSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (amauts *AsMarketAppUserTemplateSelect) Int64X(ctx context.Context) int64 {
	v, err := amauts.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amauts *AsMarketAppUserTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := amauts.sql.Query()
	if err := amauts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}