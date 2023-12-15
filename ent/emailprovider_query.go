// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/huuhoait/simple-admin-message-center/ent/emailprovider"
	"github.com/huuhoait/simple-admin-message-center/ent/predicate"
)

// EmailProviderQuery is the builder for querying EmailProvider entities.
type EmailProviderQuery struct {
	config
	ctx        *QueryContext
	order      []emailprovider.OrderOption
	inters     []Interceptor
	predicates []predicate.EmailProvider
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailProviderQuery builder.
func (epq *EmailProviderQuery) Where(ps ...predicate.EmailProvider) *EmailProviderQuery {
	epq.predicates = append(epq.predicates, ps...)
	return epq
}

// Limit the number of records to be returned by this query.
func (epq *EmailProviderQuery) Limit(limit int) *EmailProviderQuery {
	epq.ctx.Limit = &limit
	return epq
}

// Offset to start from.
func (epq *EmailProviderQuery) Offset(offset int) *EmailProviderQuery {
	epq.ctx.Offset = &offset
	return epq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (epq *EmailProviderQuery) Unique(unique bool) *EmailProviderQuery {
	epq.ctx.Unique = &unique
	return epq
}

// Order specifies how the records should be ordered.
func (epq *EmailProviderQuery) Order(o ...emailprovider.OrderOption) *EmailProviderQuery {
	epq.order = append(epq.order, o...)
	return epq
}

// First returns the first EmailProvider entity from the query.
// Returns a *NotFoundError when no EmailProvider was found.
func (epq *EmailProviderQuery) First(ctx context.Context) (*EmailProvider, error) {
	nodes, err := epq.Limit(1).All(setContextOp(ctx, epq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailprovider.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (epq *EmailProviderQuery) FirstX(ctx context.Context) *EmailProvider {
	node, err := epq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailProvider ID from the query.
// Returns a *NotFoundError when no EmailProvider ID was found.
func (epq *EmailProviderQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = epq.Limit(1).IDs(setContextOp(ctx, epq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailprovider.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (epq *EmailProviderQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := epq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailProvider entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailProvider entity is found.
// Returns a *NotFoundError when no EmailProvider entities are found.
func (epq *EmailProviderQuery) Only(ctx context.Context) (*EmailProvider, error) {
	nodes, err := epq.Limit(2).All(setContextOp(ctx, epq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailprovider.Label}
	default:
		return nil, &NotSingularError{emailprovider.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (epq *EmailProviderQuery) OnlyX(ctx context.Context) *EmailProvider {
	node, err := epq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailProvider ID in the query.
// Returns a *NotSingularError when more than one EmailProvider ID is found.
// Returns a *NotFoundError when no entities are found.
func (epq *EmailProviderQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = epq.Limit(2).IDs(setContextOp(ctx, epq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailprovider.Label}
	default:
		err = &NotSingularError{emailprovider.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (epq *EmailProviderQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := epq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailProviders.
func (epq *EmailProviderQuery) All(ctx context.Context) ([]*EmailProvider, error) {
	ctx = setContextOp(ctx, epq.ctx, "All")
	if err := epq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmailProvider, *EmailProviderQuery]()
	return withInterceptors[[]*EmailProvider](ctx, epq, qr, epq.inters)
}

// AllX is like All, but panics if an error occurs.
func (epq *EmailProviderQuery) AllX(ctx context.Context) []*EmailProvider {
	nodes, err := epq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailProvider IDs.
func (epq *EmailProviderQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if epq.ctx.Unique == nil && epq.path != nil {
		epq.Unique(true)
	}
	ctx = setContextOp(ctx, epq.ctx, "IDs")
	if err = epq.Select(emailprovider.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (epq *EmailProviderQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := epq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (epq *EmailProviderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, epq.ctx, "Count")
	if err := epq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, epq, querierCount[*EmailProviderQuery](), epq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (epq *EmailProviderQuery) CountX(ctx context.Context) int {
	count, err := epq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (epq *EmailProviderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, epq.ctx, "Exist")
	switch _, err := epq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (epq *EmailProviderQuery) ExistX(ctx context.Context) bool {
	exist, err := epq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailProviderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (epq *EmailProviderQuery) Clone() *EmailProviderQuery {
	if epq == nil {
		return nil
	}
	return &EmailProviderQuery{
		config:     epq.config,
		ctx:        epq.ctx.Clone(),
		order:      append([]emailprovider.OrderOption{}, epq.order...),
		inters:     append([]Interceptor{}, epq.inters...),
		predicates: append([]predicate.EmailProvider{}, epq.predicates...),
		// clone intermediate query.
		sql:  epq.sql.Clone(),
		path: epq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailProvider.Query().
//		GroupBy(emailprovider.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (epq *EmailProviderQuery) GroupBy(field string, fields ...string) *EmailProviderGroupBy {
	epq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmailProviderGroupBy{build: epq}
	grbuild.flds = &epq.ctx.Fields
	grbuild.label = emailprovider.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.EmailProvider.Query().
//		Select(emailprovider.FieldCreatedAt).
//		Scan(ctx, &v)
func (epq *EmailProviderQuery) Select(fields ...string) *EmailProviderSelect {
	epq.ctx.Fields = append(epq.ctx.Fields, fields...)
	sbuild := &EmailProviderSelect{EmailProviderQuery: epq}
	sbuild.label = emailprovider.Label
	sbuild.flds, sbuild.scan = &epq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmailProviderSelect configured with the given aggregations.
func (epq *EmailProviderQuery) Aggregate(fns ...AggregateFunc) *EmailProviderSelect {
	return epq.Select().Aggregate(fns...)
}

func (epq *EmailProviderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range epq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, epq); err != nil {
				return err
			}
		}
	}
	for _, f := range epq.ctx.Fields {
		if !emailprovider.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if epq.path != nil {
		prev, err := epq.path(ctx)
		if err != nil {
			return err
		}
		epq.sql = prev
	}
	return nil
}

func (epq *EmailProviderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailProvider, error) {
	var (
		nodes = []*EmailProvider{}
		_spec = epq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailProvider).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailProvider{config: epq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, epq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (epq *EmailProviderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := epq.querySpec()
	_spec.Node.Columns = epq.ctx.Fields
	if len(epq.ctx.Fields) > 0 {
		_spec.Unique = epq.ctx.Unique != nil && *epq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, epq.driver, _spec)
}

func (epq *EmailProviderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(emailprovider.Table, emailprovider.Columns, sqlgraph.NewFieldSpec(emailprovider.FieldID, field.TypeUint64))
	_spec.From = epq.sql
	if unique := epq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if epq.path != nil {
		_spec.Unique = true
	}
	if fields := epq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailprovider.FieldID)
		for i := range fields {
			if fields[i] != emailprovider.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := epq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := epq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := epq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := epq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (epq *EmailProviderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(epq.driver.Dialect())
	t1 := builder.Table(emailprovider.Table)
	columns := epq.ctx.Fields
	if len(columns) == 0 {
		columns = emailprovider.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if epq.sql != nil {
		selector = epq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if epq.ctx.Unique != nil && *epq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range epq.predicates {
		p(selector)
	}
	for _, p := range epq.order {
		p(selector)
	}
	if offset := epq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := epq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmailProviderGroupBy is the group-by builder for EmailProvider entities.
type EmailProviderGroupBy struct {
	selector
	build *EmailProviderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (epgb *EmailProviderGroupBy) Aggregate(fns ...AggregateFunc) *EmailProviderGroupBy {
	epgb.fns = append(epgb.fns, fns...)
	return epgb
}

// Scan applies the selector query and scans the result into the given value.
func (epgb *EmailProviderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, epgb.build.ctx, "GroupBy")
	if err := epgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailProviderQuery, *EmailProviderGroupBy](ctx, epgb.build, epgb, epgb.build.inters, v)
}

func (epgb *EmailProviderGroupBy) sqlScan(ctx context.Context, root *EmailProviderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(epgb.fns))
	for _, fn := range epgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*epgb.flds)+len(epgb.fns))
		for _, f := range *epgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*epgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := epgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmailProviderSelect is the builder for selecting fields of EmailProvider entities.
type EmailProviderSelect struct {
	*EmailProviderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (eps *EmailProviderSelect) Aggregate(fns ...AggregateFunc) *EmailProviderSelect {
	eps.fns = append(eps.fns, fns...)
	return eps
}

// Scan applies the selector query and scans the result into the given value.
func (eps *EmailProviderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, eps.ctx, "Select")
	if err := eps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailProviderQuery, *EmailProviderSelect](ctx, eps.EmailProviderQuery, eps, eps.inters, v)
}

func (eps *EmailProviderSelect) sqlScan(ctx context.Context, root *EmailProviderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(eps.fns))
	for _, fn := range eps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*eps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
