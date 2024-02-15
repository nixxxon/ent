// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nixxxon/entdemo/ent/predicate"
	"github.com/nixxxon/entdemo/ent/todohack"
)

// TodoHackQuery is the builder for querying TodoHack entities.
type TodoHackQuery struct {
	config
	ctx        *QueryContext
	order      []todohack.OrderOption
	inters     []Interceptor
	predicates []predicate.TodoHack
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*TodoHack) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TodoHackQuery builder.
func (thq *TodoHackQuery) Where(ps ...predicate.TodoHack) *TodoHackQuery {
	thq.predicates = append(thq.predicates, ps...)
	return thq
}

// Limit the number of records to be returned by this query.
func (thq *TodoHackQuery) Limit(limit int) *TodoHackQuery {
	thq.ctx.Limit = &limit
	return thq
}

// Offset to start from.
func (thq *TodoHackQuery) Offset(offset int) *TodoHackQuery {
	thq.ctx.Offset = &offset
	return thq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (thq *TodoHackQuery) Unique(unique bool) *TodoHackQuery {
	thq.ctx.Unique = &unique
	return thq
}

// Order specifies how the records should be ordered.
func (thq *TodoHackQuery) Order(o ...todohack.OrderOption) *TodoHackQuery {
	thq.order = append(thq.order, o...)
	return thq
}

// First returns the first TodoHack entity from the query.
// Returns a *NotFoundError when no TodoHack was found.
func (thq *TodoHackQuery) First(ctx context.Context) (*TodoHack, error) {
	nodes, err := thq.Limit(1).All(setContextOp(ctx, thq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{todohack.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (thq *TodoHackQuery) FirstX(ctx context.Context) *TodoHack {
	node, err := thq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TodoHack ID from the query.
// Returns a *NotFoundError when no TodoHack ID was found.
func (thq *TodoHackQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thq.Limit(1).IDs(setContextOp(ctx, thq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{todohack.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (thq *TodoHackQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := thq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TodoHack entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TodoHack entity is found.
// Returns a *NotFoundError when no TodoHack entities are found.
func (thq *TodoHackQuery) Only(ctx context.Context) (*TodoHack, error) {
	nodes, err := thq.Limit(2).All(setContextOp(ctx, thq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{todohack.Label}
	default:
		return nil, &NotSingularError{todohack.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (thq *TodoHackQuery) OnlyX(ctx context.Context) *TodoHack {
	node, err := thq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TodoHack ID in the query.
// Returns a *NotSingularError when more than one TodoHack ID is found.
// Returns a *NotFoundError when no entities are found.
func (thq *TodoHackQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = thq.Limit(2).IDs(setContextOp(ctx, thq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{todohack.Label}
	default:
		err = &NotSingularError{todohack.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (thq *TodoHackQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := thq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TodoHacks.
func (thq *TodoHackQuery) All(ctx context.Context) ([]*TodoHack, error) {
	ctx = setContextOp(ctx, thq.ctx, "All")
	if err := thq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TodoHack, *TodoHackQuery]()
	return withInterceptors[[]*TodoHack](ctx, thq, qr, thq.inters)
}

// AllX is like All, but panics if an error occurs.
func (thq *TodoHackQuery) AllX(ctx context.Context) []*TodoHack {
	nodes, err := thq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TodoHack IDs.
func (thq *TodoHackQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if thq.ctx.Unique == nil && thq.path != nil {
		thq.Unique(true)
	}
	ctx = setContextOp(ctx, thq.ctx, "IDs")
	if err = thq.Select(todohack.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (thq *TodoHackQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := thq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (thq *TodoHackQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, thq.ctx, "Count")
	if err := thq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, thq, querierCount[*TodoHackQuery](), thq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (thq *TodoHackQuery) CountX(ctx context.Context) int {
	count, err := thq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (thq *TodoHackQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, thq.ctx, "Exist")
	switch _, err := thq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (thq *TodoHackQuery) ExistX(ctx context.Context) bool {
	exist, err := thq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TodoHackQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (thq *TodoHackQuery) Clone() *TodoHackQuery {
	if thq == nil {
		return nil
	}
	return &TodoHackQuery{
		config:     thq.config,
		ctx:        thq.ctx.Clone(),
		order:      append([]todohack.OrderOption{}, thq.order...),
		inters:     append([]Interceptor{}, thq.inters...),
		predicates: append([]predicate.TodoHack{}, thq.predicates...),
		// clone intermediate query.
		sql:  thq.sql.Clone(),
		path: thq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TodoHack.Query().
//		GroupBy(todohack.FieldHistoryTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (thq *TodoHackQuery) GroupBy(field string, fields ...string) *TodoHackGroupBy {
	thq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TodoHackGroupBy{build: thq}
	grbuild.flds = &thq.ctx.Fields
	grbuild.label = todohack.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.TodoHack.Query().
//		Select(todohack.FieldHistoryTime).
//		Scan(ctx, &v)
func (thq *TodoHackQuery) Select(fields ...string) *TodoHackSelect {
	thq.ctx.Fields = append(thq.ctx.Fields, fields...)
	sbuild := &TodoHackSelect{TodoHackQuery: thq}
	sbuild.label = todohack.Label
	sbuild.flds, sbuild.scan = &thq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TodoHackSelect configured with the given aggregations.
func (thq *TodoHackQuery) Aggregate(fns ...AggregateFunc) *TodoHackSelect {
	return thq.Select().Aggregate(fns...)
}

func (thq *TodoHackQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range thq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, thq); err != nil {
				return err
			}
		}
	}
	for _, f := range thq.ctx.Fields {
		if !todohack.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if thq.path != nil {
		prev, err := thq.path(ctx)
		if err != nil {
			return err
		}
		thq.sql = prev
	}
	return nil
}

func (thq *TodoHackQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TodoHack, error) {
	var (
		nodes = []*TodoHack{}
		_spec = thq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TodoHack).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TodoHack{config: thq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(thq.modifiers) > 0 {
		_spec.Modifiers = thq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, thq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range thq.loadTotal {
		if err := thq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (thq *TodoHackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := thq.querySpec()
	if len(thq.modifiers) > 0 {
		_spec.Modifiers = thq.modifiers
	}
	_spec.Node.Columns = thq.ctx.Fields
	if len(thq.ctx.Fields) > 0 {
		_spec.Unique = thq.ctx.Unique != nil && *thq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, thq.driver, _spec)
}

func (thq *TodoHackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(todohack.Table, todohack.Columns, sqlgraph.NewFieldSpec(todohack.FieldID, field.TypeUUID))
	_spec.From = thq.sql
	if unique := thq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if thq.path != nil {
		_spec.Unique = true
	}
	if fields := thq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todohack.FieldID)
		for i := range fields {
			if fields[i] != todohack.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := thq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := thq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := thq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := thq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (thq *TodoHackQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(thq.driver.Dialect())
	t1 := builder.Table(todohack.Table)
	columns := thq.ctx.Fields
	if len(columns) == 0 {
		columns = todohack.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if thq.sql != nil {
		selector = thq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if thq.ctx.Unique != nil && *thq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range thq.predicates {
		p(selector)
	}
	for _, p := range thq.order {
		p(selector)
	}
	if offset := thq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := thq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TodoHackGroupBy is the group-by builder for TodoHack entities.
type TodoHackGroupBy struct {
	selector
	build *TodoHackQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (thgb *TodoHackGroupBy) Aggregate(fns ...AggregateFunc) *TodoHackGroupBy {
	thgb.fns = append(thgb.fns, fns...)
	return thgb
}

// Scan applies the selector query and scans the result into the given value.
func (thgb *TodoHackGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, thgb.build.ctx, "GroupBy")
	if err := thgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoHackQuery, *TodoHackGroupBy](ctx, thgb.build, thgb, thgb.build.inters, v)
}

func (thgb *TodoHackGroupBy) sqlScan(ctx context.Context, root *TodoHackQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(thgb.fns))
	for _, fn := range thgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*thgb.flds)+len(thgb.fns))
		for _, f := range *thgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*thgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := thgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TodoHackSelect is the builder for selecting fields of TodoHack entities.
type TodoHackSelect struct {
	*TodoHackQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ths *TodoHackSelect) Aggregate(fns ...AggregateFunc) *TodoHackSelect {
	ths.fns = append(ths.fns, fns...)
	return ths
}

// Scan applies the selector query and scans the result into the given value.
func (ths *TodoHackSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ths.ctx, "Select")
	if err := ths.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TodoHackQuery, *TodoHackSelect](ctx, ths.TodoHackQuery, ths, ths.inters, v)
}

func (ths *TodoHackSelect) sqlScan(ctx context.Context, root *TodoHackQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ths.fns))
	for _, fn := range ths.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ths.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ths.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}