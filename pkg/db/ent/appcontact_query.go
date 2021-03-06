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
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appcontact"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppContactQuery is the builder for querying AppContact entities.
type AppContactQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppContact
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppContactQuery builder.
func (acq *AppContactQuery) Where(ps ...predicate.AppContact) *AppContactQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit adds a limit step to the query.
func (acq *AppContactQuery) Limit(limit int) *AppContactQuery {
	acq.limit = &limit
	return acq
}

// Offset adds an offset step to the query.
func (acq *AppContactQuery) Offset(offset int) *AppContactQuery {
	acq.offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AppContactQuery) Unique(unique bool) *AppContactQuery {
	acq.unique = &unique
	return acq
}

// Order adds an order step to the query.
func (acq *AppContactQuery) Order(o ...OrderFunc) *AppContactQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first AppContact entity from the query.
// Returns a *NotFoundError when no AppContact was found.
func (acq *AppContactQuery) First(ctx context.Context) (*AppContact, error) {
	nodes, err := acq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appcontact.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AppContactQuery) FirstX(ctx context.Context) *AppContact {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppContact ID from the query.
// Returns a *NotFoundError when no AppContact ID was found.
func (acq *AppContactQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appcontact.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AppContactQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppContact entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppContact entity is found.
// Returns a *NotFoundError when no AppContact entities are found.
func (acq *AppContactQuery) Only(ctx context.Context) (*AppContact, error) {
	nodes, err := acq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appcontact.Label}
	default:
		return nil, &NotSingularError{appcontact.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AppContactQuery) OnlyX(ctx context.Context) *AppContact {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppContact ID in the query.
// Returns a *NotSingularError when more than one AppContact ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AppContactQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = acq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = &NotSingularError{appcontact.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AppContactQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppContacts.
func (acq *AppContactQuery) All(ctx context.Context) ([]*AppContact, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return acq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (acq *AppContactQuery) AllX(ctx context.Context) []*AppContact {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppContact IDs.
func (acq *AppContactQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := acq.Select(appcontact.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AppContactQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AppContactQuery) Count(ctx context.Context) (int, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return acq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AppContactQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AppContactQuery) Exist(ctx context.Context) (bool, error) {
	if err := acq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return acq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AppContactQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppContactQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AppContactQuery) Clone() *AppContactQuery {
	if acq == nil {
		return nil
	}
	return &AppContactQuery{
		config:     acq.config,
		limit:      acq.limit,
		offset:     acq.offset,
		order:      append([]OrderFunc{}, acq.order...),
		predicates: append([]predicate.AppContact{}, acq.predicates...),
		// clone intermediate query.
		sql:    acq.sql.Clone(),
		path:   acq.path,
		unique: acq.unique,
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
//	client.AppContact.Query().
//		GroupBy(appcontact.FieldAppID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (acq *AppContactQuery) GroupBy(field string, fields ...string) *AppContactGroupBy {
	group := &AppContactGroupBy{config: acq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := acq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return acq.sqlQuery(ctx), nil
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
//	client.AppContact.Query().
//		Select(appcontact.FieldAppID).
//		Scan(ctx, &v)
//
func (acq *AppContactQuery) Select(fields ...string) *AppContactSelect {
	acq.fields = append(acq.fields, fields...)
	return &AppContactSelect{AppContactQuery: acq}
}

func (acq *AppContactQuery) prepareQuery(ctx context.Context) error {
	for _, f := range acq.fields {
		if !appcontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AppContactQuery) sqlAll(ctx context.Context) ([]*AppContact, error) {
	var (
		nodes = []*AppContact{}
		_spec = acq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AppContact{config: acq.config}
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
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acq *AppContactQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	_spec.Node.Columns = acq.fields
	if len(acq.fields) > 0 {
		_spec.Unique = acq.unique != nil && *acq.unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AppContactQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := acq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (acq *AppContactQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appcontact.Table,
			Columns: appcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcontact.FieldID,
			},
		},
		From:   acq.sql,
		Unique: true,
	}
	if unique := acq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := acq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appcontact.FieldID)
		for i := range fields {
			if fields[i] != appcontact.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AppContactQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(appcontact.Table)
	columns := acq.fields
	if len(columns) == 0 {
		columns = appcontact.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.unique != nil && *acq.unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AppContactGroupBy is the group-by builder for AppContact entities.
type AppContactGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AppContactGroupBy) Aggregate(fns ...AggregateFunc) *AppContactGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the group-by query and scans the result into the given value.
func (acgb *AppContactGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := acgb.path(ctx)
	if err != nil {
		return err
	}
	acgb.sql = query
	return acgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acgb *AppContactGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := acgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AppContactGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acgb *AppContactGroupBy) StringsX(ctx context.Context) []string {
	v, err := acgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acgb *AppContactGroupBy) StringX(ctx context.Context) string {
	v, err := acgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AppContactGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acgb *AppContactGroupBy) IntsX(ctx context.Context) []int {
	v, err := acgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acgb *AppContactGroupBy) IntX(ctx context.Context) int {
	v, err := acgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AppContactGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acgb *AppContactGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := acgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acgb *AppContactGroupBy) Float64X(ctx context.Context) float64 {
	v, err := acgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(acgb.fields) > 1 {
		return nil, errors.New("ent: AppContactGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := acgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acgb *AppContactGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := acgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (acgb *AppContactGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acgb *AppContactGroupBy) BoolX(ctx context.Context) bool {
	v, err := acgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acgb *AppContactGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range acgb.fields {
		if !appcontact.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := acgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (acgb *AppContactGroupBy) sqlQuery() *sql.Selector {
	selector := acgb.sql.Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(acgb.fields)+len(acgb.fns))
		for _, f := range acgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(acgb.fields...)...)
}

// AppContactSelect is the builder for selecting fields of AppContact entities.
type AppContactSelect struct {
	*AppContactQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AppContactSelect) Scan(ctx context.Context, v interface{}) error {
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	acs.sql = acs.AppContactQuery.sqlQuery(ctx)
	return acs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (acs *AppContactSelect) ScanX(ctx context.Context, v interface{}) {
	if err := acs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Strings(ctx context.Context) ([]string, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AppContactSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (acs *AppContactSelect) StringsX(ctx context.Context) []string {
	v, err := acs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = acs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (acs *AppContactSelect) StringX(ctx context.Context) string {
	v, err := acs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Ints(ctx context.Context) ([]int, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AppContactSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (acs *AppContactSelect) IntsX(ctx context.Context) []int {
	v, err := acs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = acs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (acs *AppContactSelect) IntX(ctx context.Context) int {
	v, err := acs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AppContactSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (acs *AppContactSelect) Float64sX(ctx context.Context) []float64 {
	v, err := acs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = acs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (acs *AppContactSelect) Float64X(ctx context.Context) float64 {
	v, err := acs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(acs.fields) > 1 {
		return nil, errors.New("ent: AppContactSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := acs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (acs *AppContactSelect) BoolsX(ctx context.Context) []bool {
	v, err := acs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (acs *AppContactSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = acs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{appcontact.Label}
	default:
		err = fmt.Errorf("ent: AppContactSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (acs *AppContactSelect) BoolX(ctx context.Context) bool {
	v, err := acs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (acs *AppContactSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := acs.sql.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
