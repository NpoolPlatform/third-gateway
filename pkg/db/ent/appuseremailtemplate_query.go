// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appuseremailtemplate"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserEmailTemplateQuery is the builder for querying AppUserEmailTemplate entities.
type AppUserEmailTemplateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUserEmailTemplate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserEmailTemplateQuery builder.
func (auetq *AppUserEmailTemplateQuery) Where(ps ...predicate.AppUserEmailTemplate) *AppUserEmailTemplateQuery {
	auetq.predicates = append(auetq.predicates, ps...)
	return auetq
}

// Limit adds a limit step to the query.
func (auetq *AppUserEmailTemplateQuery) Limit(limit int) *AppUserEmailTemplateQuery {
	auetq.limit = &limit
	return auetq
}

// Offset adds an offset step to the query.
func (auetq *AppUserEmailTemplateQuery) Offset(offset int) *AppUserEmailTemplateQuery {
	auetq.offset = &offset
	return auetq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auetq *AppUserEmailTemplateQuery) Unique(unique bool) *AppUserEmailTemplateQuery {
	auetq.unique = &unique
	return auetq
}

// Order adds an order step to the query.
func (auetq *AppUserEmailTemplateQuery) Order(o ...OrderFunc) *AppUserEmailTemplateQuery {
	auetq.order = append(auetq.order, o...)
	return auetq
}

// First returns the first AppUserEmailTemplate entity from the query.
// Returns a *NotFoundError when no AppUserEmailTemplate was found.
func (auetq *AppUserEmailTemplateQuery) First(ctx context.Context) (*AppUserEmailTemplate, error) {
	nodes, err := auetq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuseremailtemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) FirstX(ctx context.Context) *AppUserEmailTemplate {
	node, err := auetq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUserEmailTemplate ID from the query.
// Returns a *NotFoundError when no AppUserEmailTemplate ID was found.
func (auetq *AppUserEmailTemplateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auetq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuseremailtemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := auetq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUserEmailTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AppUserEmailTemplate entity is not found.
// Returns a *NotFoundError when no AppUserEmailTemplate entities are found.
func (auetq *AppUserEmailTemplateQuery) Only(ctx context.Context) (*AppUserEmailTemplate, error) {
	nodes, err := auetq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuseremailtemplate.Label}
	default:
		return nil, &NotSingularError{appuseremailtemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) OnlyX(ctx context.Context) *AppUserEmailTemplate {
	node, err := auetq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUserEmailTemplate ID in the query.
// Returns a *NotSingularError when exactly one AppUserEmailTemplate ID is not found.
// Returns a *NotFoundError when no entities are found.
func (auetq *AppUserEmailTemplateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auetq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = &NotSingularError{appuseremailtemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := auetq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUserEmailTemplates.
func (auetq *AppUserEmailTemplateQuery) All(ctx context.Context) ([]*AppUserEmailTemplate, error) {
	if err := auetq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auetq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) AllX(ctx context.Context) []*AppUserEmailTemplate {
	nodes, err := auetq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUserEmailTemplate IDs.
func (auetq *AppUserEmailTemplateQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := auetq.Select(appuseremailtemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := auetq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auetq *AppUserEmailTemplateQuery) Count(ctx context.Context) (int, error) {
	if err := auetq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auetq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) CountX(ctx context.Context) int {
	count, err := auetq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auetq *AppUserEmailTemplateQuery) Exist(ctx context.Context) (bool, error) {
	if err := auetq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auetq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auetq *AppUserEmailTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := auetq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserEmailTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auetq *AppUserEmailTemplateQuery) Clone() *AppUserEmailTemplateQuery {
	if auetq == nil {
		return nil
	}
	return &AppUserEmailTemplateQuery{
		config:     auetq.config,
		limit:      auetq.limit,
		offset:     auetq.offset,
		order:      append([]OrderFunc{}, auetq.order...),
		predicates: append([]predicate.AppUserEmailTemplate{}, auetq.predicates...),
		// clone intermediate query.
		sql:  auetq.sql.Clone(),
		path: auetq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppUserEmailTemplate.Query().
//		GroupBy(appuseremailtemplate.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (auetq *AppUserEmailTemplateQuery) GroupBy(field string, fields ...string) *AppUserEmailTemplateGroupBy {
	group := &AppUserEmailTemplateGroupBy{config: auetq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auetq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auetq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		AppID uuid.UUID `json:"app_id,omitempty"`
//	}
//
//	client.AppUserEmailTemplate.Query().
//		Select(appuseremailtemplate.FieldAppID).
//		Scan(ctx, &v)
//
func (auetq *AppUserEmailTemplateQuery) Select(fields ...string) *AppUserEmailTemplateSelect {
	auetq.fields = append(auetq.fields, fields...)
	return &AppUserEmailTemplateSelect{AppUserEmailTemplateQuery: auetq}
}

func (auetq *AppUserEmailTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range auetq.fields {
		if !appuseremailtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auetq.path != nil {
		prev, err := auetq.path(ctx)
		if err != nil {
			return err
		}
		auetq.sql = prev
	}
	return nil
}

func (auetq *AppUserEmailTemplateQuery) sqlAll(ctx context.Context) ([]*AppUserEmailTemplate, error) {
	var (
		nodes = []*AppUserEmailTemplate{}
		_spec = auetq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppUserEmailTemplate{config: auetq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, auetq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auetq *AppUserEmailTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auetq.querySpec()
	_spec.Node.Columns = auetq.fields
	if len(auetq.fields) > 0 {
		_spec.Unique = auetq.unique != nil && *auetq.unique
	}
	return sqlgraph.CountNodes(ctx, auetq.driver, _spec)
}

func (auetq *AppUserEmailTemplateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auetq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (auetq *AppUserEmailTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuseremailtemplate.Table,
			Columns: appuseremailtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuseremailtemplate.FieldID,
			},
		},
		From:   auetq.sql,
		Unique: true,
	}
	if unique := auetq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := auetq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuseremailtemplate.FieldID)
		for i := range fields {
			if fields[i] != appuseremailtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auetq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auetq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auetq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auetq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auetq *AppUserEmailTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auetq.driver.Dialect())
	t1 := builder.Table(appuseremailtemplate.Table)
	columns := auetq.fields
	if len(columns) == 0 {
		columns = appuseremailtemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auetq.sql != nil {
		selector = auetq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auetq.unique != nil && *auetq.unique {
		selector.Distinct()
	}
	for _, p := range auetq.predicates {
		p(selector)
	}
	for _, p := range auetq.order {
		p(selector)
	}
	if offset := auetq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auetq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppUserEmailTemplateGroupBy is the group-by builder for AppUserEmailTemplate entities.
type AppUserEmailTemplateGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (auetgb *AppUserEmailTemplateGroupBy) Aggregate(fns ...AggregateFunc) *AppUserEmailTemplateGroupBy {
	auetgb.fns = append(auetgb.fns, fns...)
	return auetgb
}

// Scan applies the group-by query and scans the result into the given value.
func (auetgb *AppUserEmailTemplateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := auetgb.path(ctx)
	if err != nil {
		return err
	}
	auetgb.sql = query
	return auetgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := auetgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(auetgb.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := auetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) StringsX(ctx context.Context) []string {
	v, err := auetgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auetgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) StringX(ctx context.Context) string {
	v, err := auetgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(auetgb.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := auetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) IntsX(ctx context.Context) []int {
	v, err := auetgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auetgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) IntX(ctx context.Context) int {
	v, err := auetgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(auetgb.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := auetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := auetgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auetgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) Float64X(ctx context.Context) float64 {
	v, err := auetgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(auetgb.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := auetgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := auetgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (auetgb *AppUserEmailTemplateGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auetgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auetgb *AppUserEmailTemplateGroupBy) BoolX(ctx context.Context) bool {
	v, err := auetgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auetgb *AppUserEmailTemplateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range auetgb.fields {
		if !appuseremailtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := auetgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := auetgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (auetgb *AppUserEmailTemplateGroupBy) sqlQuery() *sql.Selector {
	selector := auetgb.sql.Select()
	aggregation := make([]string, 0, len(auetgb.fns))
	for _, fn := range auetgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(auetgb.fields)+len(auetgb.fns))
		for _, f := range auetgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(auetgb.fields...)...)
}

// AppUserEmailTemplateSelect is the builder for selecting fields of AppUserEmailTemplate entities.
type AppUserEmailTemplateSelect struct {
	*AppUserEmailTemplateQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (auets *AppUserEmailTemplateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := auets.prepareQuery(ctx); err != nil {
		return err
	}
	auets.sql = auets.AppUserEmailTemplateQuery.sqlQuery(ctx)
	return auets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) ScanX(ctx context.Context, v interface{}) {
	if err := auets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Strings(ctx context.Context) ([]string, error) {
	if len(auets.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := auets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) StringsX(ctx context.Context) []string {
	v, err := auets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = auets.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) StringX(ctx context.Context) string {
	v, err := auets.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Ints(ctx context.Context) ([]int, error) {
	if len(auets.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := auets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) IntsX(ctx context.Context) []int {
	v, err := auets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = auets.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) IntX(ctx context.Context) int {
	v, err := auets.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(auets.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := auets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) Float64sX(ctx context.Context) []float64 {
	v, err := auets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = auets.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) Float64X(ctx context.Context) float64 {
	v, err := auets.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(auets.fields) > 1 {
		return nil, errors.New("ent: AppUserEmailTemplateSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := auets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) BoolsX(ctx context.Context) []bool {
	v, err := auets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (auets *AppUserEmailTemplateSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = auets.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appuseremailtemplate.Label}
	default:
		err = fmt.Errorf("ent: AppUserEmailTemplateSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (auets *AppUserEmailTemplateSelect) BoolX(ctx context.Context) bool {
	v, err := auets.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (auets *AppUserEmailTemplateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := auets.sql.Query()
	if err := auets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
