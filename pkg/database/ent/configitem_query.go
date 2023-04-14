// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/configitem"
	"github.com/crowdsecurity/crowdsec/pkg/database/ent/predicate"
)

// ConfigItemQuery is the builder for querying ConfigItem entities.
type ConfigItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ConfigItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConfigItemQuery builder.
func (ciq *ConfigItemQuery) Where(ps ...predicate.ConfigItem) *ConfigItemQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit adds a limit step to the query.
func (ciq *ConfigItemQuery) Limit(limit int) *ConfigItemQuery {
	ciq.limit = &limit
	return ciq
}

// Offset adds an offset step to the query.
func (ciq *ConfigItemQuery) Offset(offset int) *ConfigItemQuery {
	ciq.offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *ConfigItemQuery) Unique(unique bool) *ConfigItemQuery {
	ciq.unique = &unique
	return ciq
}

// Order adds an order step to the query.
func (ciq *ConfigItemQuery) Order(o ...OrderFunc) *ConfigItemQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// First returns the first ConfigItem entity from the query.
// Returns a *NotFoundError when no ConfigItem was found.
func (ciq *ConfigItemQuery) First(ctx context.Context) (*ConfigItem, error) {
	nodes, err := ciq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{configitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *ConfigItemQuery) FirstX(ctx context.Context) *ConfigItem {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ConfigItem ID from the query.
// Returns a *NotFoundError when no ConfigItem ID was found.
func (ciq *ConfigItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{configitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *ConfigItemQuery) FirstIDX(ctx context.Context) int {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ConfigItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ConfigItem entity is found.
// Returns a *NotFoundError when no ConfigItem entities are found.
func (ciq *ConfigItemQuery) Only(ctx context.Context) (*ConfigItem, error) {
	nodes, err := ciq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{configitem.Label}
	default:
		return nil, &NotSingularError{configitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *ConfigItemQuery) OnlyX(ctx context.Context) *ConfigItem {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ConfigItem ID in the query.
// Returns a *NotSingularError when more than one ConfigItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *ConfigItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ciq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{configitem.Label}
	default:
		err = &NotSingularError{configitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *ConfigItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ConfigItems.
func (ciq *ConfigItemQuery) All(ctx context.Context) ([]*ConfigItem, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ciq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ciq *ConfigItemQuery) AllX(ctx context.Context) []*ConfigItem {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ConfigItem IDs.
func (ciq *ConfigItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ciq.Select(configitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *ConfigItemQuery) IDsX(ctx context.Context) []int {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *ConfigItemQuery) Count(ctx context.Context) (int, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ciq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *ConfigItemQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *ConfigItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := ciq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ciq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *ConfigItemQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConfigItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *ConfigItemQuery) Clone() *ConfigItemQuery {
	if ciq == nil {
		return nil
	}
	return &ConfigItemQuery{
		config:     ciq.config,
		limit:      ciq.limit,
		offset:     ciq.offset,
		order:      append([]OrderFunc{}, ciq.order...),
		predicates: append([]predicate.ConfigItem{}, ciq.predicates...),
		// clone intermediate query.
		sql:    ciq.sql.Clone(),
		path:   ciq.path,
		unique: ciq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ConfigItem.Query().
//		GroupBy(configitem.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ciq *ConfigItemQuery) GroupBy(field string, fields ...string) *ConfigItemGroupBy {
	grbuild := &ConfigItemGroupBy{config: ciq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ciq.sqlQuery(ctx), nil
	}
	grbuild.label = configitem.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at"`
//	}
//
//	client.ConfigItem.Query().
//		Select(configitem.FieldCreatedAt).
//		Scan(ctx, &v)
func (ciq *ConfigItemQuery) Select(fields ...string) *ConfigItemSelect {
	ciq.fields = append(ciq.fields, fields...)
	selbuild := &ConfigItemSelect{ConfigItemQuery: ciq}
	selbuild.label = configitem.Label
	selbuild.flds, selbuild.scan = &ciq.fields, selbuild.Scan
	return selbuild
}

func (ciq *ConfigItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ciq.fields {
		if !configitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *ConfigItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ConfigItem, error) {
	var (
		nodes = []*ConfigItem{}
		_spec = ciq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ConfigItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ConfigItem{config: ciq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ciq *ConfigItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	_spec.Node.Columns = ciq.fields
	if len(ciq.fields) > 0 {
		_spec.Unique = ciq.unique != nil && *ciq.unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *ConfigItemQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := ciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (ciq *ConfigItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   configitem.Table,
			Columns: configitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configitem.FieldID,
			},
		},
		From:   ciq.sql,
		Unique: true,
	}
	if unique := ciq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ciq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, configitem.FieldID)
		for i := range fields {
			if fields[i] != configitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *ConfigItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(configitem.Table)
	columns := ciq.fields
	if len(columns) == 0 {
		columns = configitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.unique != nil && *ciq.unique {
		selector.Distinct()
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConfigItemGroupBy is the group-by builder for ConfigItem entities.
type ConfigItemGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *ConfigItemGroupBy) Aggregate(fns ...AggregateFunc) *ConfigItemGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the group-by query and scans the result into the given value.
func (cigb *ConfigItemGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cigb.path(ctx)
	if err != nil {
		return err
	}
	cigb.sql = query
	return cigb.sqlScan(ctx, v)
}

func (cigb *ConfigItemGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cigb.fields {
		if !configitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cigb *ConfigItemGroupBy) sqlQuery() *sql.Selector {
	selector := cigb.sql.Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cigb.fields)+len(cigb.fns))
		for _, f := range cigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cigb.fields...)...)
}

// ConfigItemSelect is the builder for selecting fields of ConfigItem entities.
type ConfigItemSelect struct {
	*ConfigItemQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cis *ConfigItemSelect) Scan(ctx context.Context, v any) error {
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	cis.sql = cis.ConfigItemQuery.sqlQuery(ctx)
	return cis.sqlScan(ctx, v)
}

func (cis *ConfigItemSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := cis.sql.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
