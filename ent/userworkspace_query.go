// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"slack-application/ent/predicate"
	"slack-application/ent/userworkspace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserWorkspaceQuery is the builder for querying UserWorkspace entities.
type UserWorkspaceQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.UserWorkspace
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserWorkspaceQuery builder.
func (uwq *UserWorkspaceQuery) Where(ps ...predicate.UserWorkspace) *UserWorkspaceQuery {
	uwq.predicates = append(uwq.predicates, ps...)
	return uwq
}

// Limit the number of records to be returned by this query.
func (uwq *UserWorkspaceQuery) Limit(limit int) *UserWorkspaceQuery {
	uwq.ctx.Limit = &limit
	return uwq
}

// Offset to start from.
func (uwq *UserWorkspaceQuery) Offset(offset int) *UserWorkspaceQuery {
	uwq.ctx.Offset = &offset
	return uwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uwq *UserWorkspaceQuery) Unique(unique bool) *UserWorkspaceQuery {
	uwq.ctx.Unique = &unique
	return uwq
}

// Order specifies how the records should be ordered.
func (uwq *UserWorkspaceQuery) Order(o ...OrderFunc) *UserWorkspaceQuery {
	uwq.order = append(uwq.order, o...)
	return uwq
}

// First returns the first UserWorkspace entity from the query.
// Returns a *NotFoundError when no UserWorkspace was found.
func (uwq *UserWorkspaceQuery) First(ctx context.Context) (*UserWorkspace, error) {
	nodes, err := uwq.Limit(1).All(setContextOp(ctx, uwq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userworkspace.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) FirstX(ctx context.Context) *UserWorkspace {
	node, err := uwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserWorkspace ID from the query.
// Returns a *NotFoundError when no UserWorkspace ID was found.
func (uwq *UserWorkspaceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uwq.Limit(1).IDs(setContextOp(ctx, uwq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userworkspace.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) FirstIDX(ctx context.Context) int {
	id, err := uwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserWorkspace entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserWorkspace entity is found.
// Returns a *NotFoundError when no UserWorkspace entities are found.
func (uwq *UserWorkspaceQuery) Only(ctx context.Context) (*UserWorkspace, error) {
	nodes, err := uwq.Limit(2).All(setContextOp(ctx, uwq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userworkspace.Label}
	default:
		return nil, &NotSingularError{userworkspace.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) OnlyX(ctx context.Context) *UserWorkspace {
	node, err := uwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserWorkspace ID in the query.
// Returns a *NotSingularError when more than one UserWorkspace ID is found.
// Returns a *NotFoundError when no entities are found.
func (uwq *UserWorkspaceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uwq.Limit(2).IDs(setContextOp(ctx, uwq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userworkspace.Label}
	default:
		err = &NotSingularError{userworkspace.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) OnlyIDX(ctx context.Context) int {
	id, err := uwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserWorkspaces.
func (uwq *UserWorkspaceQuery) All(ctx context.Context) ([]*UserWorkspace, error) {
	ctx = setContextOp(ctx, uwq.ctx, "All")
	if err := uwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserWorkspace, *UserWorkspaceQuery]()
	return withInterceptors[[]*UserWorkspace](ctx, uwq, qr, uwq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) AllX(ctx context.Context) []*UserWorkspace {
	nodes, err := uwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserWorkspace IDs.
func (uwq *UserWorkspaceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uwq.ctx.Unique == nil && uwq.path != nil {
		uwq.Unique(true)
	}
	ctx = setContextOp(ctx, uwq.ctx, "IDs")
	if err = uwq.Select(userworkspace.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) IDsX(ctx context.Context) []int {
	ids, err := uwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uwq *UserWorkspaceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uwq.ctx, "Count")
	if err := uwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uwq, querierCount[*UserWorkspaceQuery](), uwq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) CountX(ctx context.Context) int {
	count, err := uwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uwq *UserWorkspaceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uwq.ctx, "Exist")
	switch _, err := uwq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uwq *UserWorkspaceQuery) ExistX(ctx context.Context) bool {
	exist, err := uwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserWorkspaceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uwq *UserWorkspaceQuery) Clone() *UserWorkspaceQuery {
	if uwq == nil {
		return nil
	}
	return &UserWorkspaceQuery{
		config:     uwq.config,
		ctx:        uwq.ctx.Clone(),
		order:      append([]OrderFunc{}, uwq.order...),
		inters:     append([]Interceptor{}, uwq.inters...),
		predicates: append([]predicate.UserWorkspace{}, uwq.predicates...),
		// clone intermediate query.
		sql:  uwq.sql.Clone(),
		path: uwq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserWorkspace.Query().
//		GroupBy(userworkspace.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uwq *UserWorkspaceQuery) GroupBy(field string, fields ...string) *UserWorkspaceGroupBy {
	uwq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserWorkspaceGroupBy{build: uwq}
	grbuild.flds = &uwq.ctx.Fields
	grbuild.label = userworkspace.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.UserWorkspace.Query().
//		Select(userworkspace.FieldUserID).
//		Scan(ctx, &v)
func (uwq *UserWorkspaceQuery) Select(fields ...string) *UserWorkspaceSelect {
	uwq.ctx.Fields = append(uwq.ctx.Fields, fields...)
	sbuild := &UserWorkspaceSelect{UserWorkspaceQuery: uwq}
	sbuild.label = userworkspace.Label
	sbuild.flds, sbuild.scan = &uwq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserWorkspaceSelect configured with the given aggregations.
func (uwq *UserWorkspaceQuery) Aggregate(fns ...AggregateFunc) *UserWorkspaceSelect {
	return uwq.Select().Aggregate(fns...)
}

func (uwq *UserWorkspaceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uwq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uwq); err != nil {
				return err
			}
		}
	}
	for _, f := range uwq.ctx.Fields {
		if !userworkspace.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uwq.path != nil {
		prev, err := uwq.path(ctx)
		if err != nil {
			return err
		}
		uwq.sql = prev
	}
	return nil
}

func (uwq *UserWorkspaceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserWorkspace, error) {
	var (
		nodes = []*UserWorkspace{}
		_spec = uwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserWorkspace).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserWorkspace{config: uwq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uwq *UserWorkspaceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uwq.querySpec()
	_spec.Node.Columns = uwq.ctx.Fields
	if len(uwq.ctx.Fields) > 0 {
		_spec.Unique = uwq.ctx.Unique != nil && *uwq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uwq.driver, _spec)
}

func (uwq *UserWorkspaceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userworkspace.Table, userworkspace.Columns, sqlgraph.NewFieldSpec(userworkspace.FieldID, field.TypeInt))
	_spec.From = uwq.sql
	if unique := uwq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uwq.path != nil {
		_spec.Unique = true
	}
	if fields := uwq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userworkspace.FieldID)
		for i := range fields {
			if fields[i] != userworkspace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uwq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uwq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uwq *UserWorkspaceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uwq.driver.Dialect())
	t1 := builder.Table(userworkspace.Table)
	columns := uwq.ctx.Fields
	if len(columns) == 0 {
		columns = userworkspace.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uwq.sql != nil {
		selector = uwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uwq.ctx.Unique != nil && *uwq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uwq.predicates {
		p(selector)
	}
	for _, p := range uwq.order {
		p(selector)
	}
	if offset := uwq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uwq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserWorkspaceGroupBy is the group-by builder for UserWorkspace entities.
type UserWorkspaceGroupBy struct {
	selector
	build *UserWorkspaceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uwgb *UserWorkspaceGroupBy) Aggregate(fns ...AggregateFunc) *UserWorkspaceGroupBy {
	uwgb.fns = append(uwgb.fns, fns...)
	return uwgb
}

// Scan applies the selector query and scans the result into the given value.
func (uwgb *UserWorkspaceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uwgb.build.ctx, "GroupBy")
	if err := uwgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserWorkspaceQuery, *UserWorkspaceGroupBy](ctx, uwgb.build, uwgb, uwgb.build.inters, v)
}

func (uwgb *UserWorkspaceGroupBy) sqlScan(ctx context.Context, root *UserWorkspaceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uwgb.fns))
	for _, fn := range uwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uwgb.flds)+len(uwgb.fns))
		for _, f := range *uwgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uwgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uwgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserWorkspaceSelect is the builder for selecting fields of UserWorkspace entities.
type UserWorkspaceSelect struct {
	*UserWorkspaceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uws *UserWorkspaceSelect) Aggregate(fns ...AggregateFunc) *UserWorkspaceSelect {
	uws.fns = append(uws.fns, fns...)
	return uws
}

// Scan applies the selector query and scans the result into the given value.
func (uws *UserWorkspaceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uws.ctx, "Select")
	if err := uws.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserWorkspaceQuery, *UserWorkspaceSelect](ctx, uws.UserWorkspaceQuery, uws, uws.inters, v)
}

func (uws *UserWorkspaceSelect) sqlScan(ctx context.Context, root *UserWorkspaceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uws.fns))
	for _, fn := range uws.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uws.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
