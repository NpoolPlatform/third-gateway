// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appsmstemplate"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppSMSTemplateUpdate is the builder for updating AppSMSTemplate entities.
type AppSMSTemplateUpdate struct {
	config
	hooks    []Hook
	mutation *AppSMSTemplateMutation
}

// Where appends a list predicates to the AppSMSTemplateUpdate builder.
func (astu *AppSMSTemplateUpdate) Where(ps ...predicate.AppSMSTemplate) *AppSMSTemplateUpdate {
	astu.mutation.Where(ps...)
	return astu
}

// SetAppID sets the "app_id" field.
func (astu *AppSMSTemplateUpdate) SetAppID(u uuid.UUID) *AppSMSTemplateUpdate {
	astu.mutation.SetAppID(u)
	return astu
}

// SetLangID sets the "lang_id" field.
func (astu *AppSMSTemplateUpdate) SetLangID(u uuid.UUID) *AppSMSTemplateUpdate {
	astu.mutation.SetLangID(u)
	return astu
}

// SetSubject sets the "subject" field.
func (astu *AppSMSTemplateUpdate) SetSubject(s string) *AppSMSTemplateUpdate {
	astu.mutation.SetSubject(s)
	return astu
}

// SetMessage sets the "message" field.
func (astu *AppSMSTemplateUpdate) SetMessage(s string) *AppSMSTemplateUpdate {
	astu.mutation.SetMessage(s)
	return astu
}

// SetCreateAt sets the "create_at" field.
func (astu *AppSMSTemplateUpdate) SetCreateAt(u uint32) *AppSMSTemplateUpdate {
	astu.mutation.ResetCreateAt()
	astu.mutation.SetCreateAt(u)
	return astu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (astu *AppSMSTemplateUpdate) SetNillableCreateAt(u *uint32) *AppSMSTemplateUpdate {
	if u != nil {
		astu.SetCreateAt(*u)
	}
	return astu
}

// AddCreateAt adds u to the "create_at" field.
func (astu *AppSMSTemplateUpdate) AddCreateAt(u int32) *AppSMSTemplateUpdate {
	astu.mutation.AddCreateAt(u)
	return astu
}

// SetUpdateAt sets the "update_at" field.
func (astu *AppSMSTemplateUpdate) SetUpdateAt(u uint32) *AppSMSTemplateUpdate {
	astu.mutation.ResetUpdateAt()
	astu.mutation.SetUpdateAt(u)
	return astu
}

// AddUpdateAt adds u to the "update_at" field.
func (astu *AppSMSTemplateUpdate) AddUpdateAt(u int32) *AppSMSTemplateUpdate {
	astu.mutation.AddUpdateAt(u)
	return astu
}

// Mutation returns the AppSMSTemplateMutation object of the builder.
func (astu *AppSMSTemplateUpdate) Mutation() *AppSMSTemplateMutation {
	return astu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (astu *AppSMSTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	astu.defaults()
	if len(astu.hooks) == 0 {
		affected, err = astu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSMSTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			astu.mutation = mutation
			affected, err = astu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(astu.hooks) - 1; i >= 0; i-- {
			if astu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = astu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, astu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (astu *AppSMSTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := astu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (astu *AppSMSTemplateUpdate) Exec(ctx context.Context) error {
	_, err := astu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (astu *AppSMSTemplateUpdate) ExecX(ctx context.Context) {
	if err := astu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (astu *AppSMSTemplateUpdate) defaults() {
	if _, ok := astu.mutation.UpdateAt(); !ok {
		v := appsmstemplate.UpdateDefaultUpdateAt()
		astu.mutation.SetUpdateAt(v)
	}
}

func (astu *AppSMSTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsmstemplate.Table,
			Columns: appsmstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appsmstemplate.FieldID,
			},
		},
	}
	if ps := astu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := astu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldAppID,
		})
	}
	if value, ok := astu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldLangID,
		})
	}
	if value, ok := astu.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldSubject,
		})
	}
	if value, ok := astu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldMessage,
		})
	}
	if value, ok := astu.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldCreateAt,
		})
	}
	if value, ok := astu.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldCreateAt,
		})
	}
	if value, ok := astu.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldUpdateAt,
		})
	}
	if value, ok := astu.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldUpdateAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, astu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appsmstemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AppSMSTemplateUpdateOne is the builder for updating a single AppSMSTemplate entity.
type AppSMSTemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AppSMSTemplateMutation
}

// SetAppID sets the "app_id" field.
func (astuo *AppSMSTemplateUpdateOne) SetAppID(u uuid.UUID) *AppSMSTemplateUpdateOne {
	astuo.mutation.SetAppID(u)
	return astuo
}

// SetLangID sets the "lang_id" field.
func (astuo *AppSMSTemplateUpdateOne) SetLangID(u uuid.UUID) *AppSMSTemplateUpdateOne {
	astuo.mutation.SetLangID(u)
	return astuo
}

// SetSubject sets the "subject" field.
func (astuo *AppSMSTemplateUpdateOne) SetSubject(s string) *AppSMSTemplateUpdateOne {
	astuo.mutation.SetSubject(s)
	return astuo
}

// SetMessage sets the "message" field.
func (astuo *AppSMSTemplateUpdateOne) SetMessage(s string) *AppSMSTemplateUpdateOne {
	astuo.mutation.SetMessage(s)
	return astuo
}

// SetCreateAt sets the "create_at" field.
func (astuo *AppSMSTemplateUpdateOne) SetCreateAt(u uint32) *AppSMSTemplateUpdateOne {
	astuo.mutation.ResetCreateAt()
	astuo.mutation.SetCreateAt(u)
	return astuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (astuo *AppSMSTemplateUpdateOne) SetNillableCreateAt(u *uint32) *AppSMSTemplateUpdateOne {
	if u != nil {
		astuo.SetCreateAt(*u)
	}
	return astuo
}

// AddCreateAt adds u to the "create_at" field.
func (astuo *AppSMSTemplateUpdateOne) AddCreateAt(u int32) *AppSMSTemplateUpdateOne {
	astuo.mutation.AddCreateAt(u)
	return astuo
}

// SetUpdateAt sets the "update_at" field.
func (astuo *AppSMSTemplateUpdateOne) SetUpdateAt(u uint32) *AppSMSTemplateUpdateOne {
	astuo.mutation.ResetUpdateAt()
	astuo.mutation.SetUpdateAt(u)
	return astuo
}

// AddUpdateAt adds u to the "update_at" field.
func (astuo *AppSMSTemplateUpdateOne) AddUpdateAt(u int32) *AppSMSTemplateUpdateOne {
	astuo.mutation.AddUpdateAt(u)
	return astuo
}

// Mutation returns the AppSMSTemplateMutation object of the builder.
func (astuo *AppSMSTemplateUpdateOne) Mutation() *AppSMSTemplateMutation {
	return astuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (astuo *AppSMSTemplateUpdateOne) Select(field string, fields ...string) *AppSMSTemplateUpdateOne {
	astuo.fields = append([]string{field}, fields...)
	return astuo
}

// Save executes the query and returns the updated AppSMSTemplate entity.
func (astuo *AppSMSTemplateUpdateOne) Save(ctx context.Context) (*AppSMSTemplate, error) {
	var (
		err  error
		node *AppSMSTemplate
	)
	astuo.defaults()
	if len(astuo.hooks) == 0 {
		node, err = astuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSMSTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			astuo.mutation = mutation
			node, err = astuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(astuo.hooks) - 1; i >= 0; i-- {
			if astuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = astuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, astuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (astuo *AppSMSTemplateUpdateOne) SaveX(ctx context.Context) *AppSMSTemplate {
	node, err := astuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (astuo *AppSMSTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := astuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (astuo *AppSMSTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := astuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (astuo *AppSMSTemplateUpdateOne) defaults() {
	if _, ok := astuo.mutation.UpdateAt(); !ok {
		v := appsmstemplate.UpdateDefaultUpdateAt()
		astuo.mutation.SetUpdateAt(v)
	}
}

func (astuo *AppSMSTemplateUpdateOne) sqlSave(ctx context.Context) (_node *AppSMSTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsmstemplate.Table,
			Columns: appsmstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appsmstemplate.FieldID,
			},
		},
	}
	id, ok := astuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppSMSTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := astuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appsmstemplate.FieldID)
		for _, f := range fields {
			if !appsmstemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appsmstemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := astuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := astuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldAppID,
		})
	}
	if value, ok := astuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldLangID,
		})
	}
	if value, ok := astuo.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldSubject,
		})
	}
	if value, ok := astuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldMessage,
		})
	}
	if value, ok := astuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldCreateAt,
		})
	}
	if value, ok := astuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldCreateAt,
		})
	}
	if value, ok := astuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldUpdateAt,
		})
	}
	if value, ok := astuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldUpdateAt,
		})
	}
	_node = &AppSMSTemplate{config: astuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, astuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appsmstemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
