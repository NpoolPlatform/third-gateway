// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-gateway/pkg/db/ent/appuseremailtemplate"
	"github.com/google/uuid"
)

// AppUserEmailTemplateCreate is the builder for creating a AppUserEmailTemplate entity.
type AppUserEmailTemplateCreate struct {
	config
	mutation *AppUserEmailTemplateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (auetc *AppUserEmailTemplateCreate) SetAppID(u uuid.UUID) *AppUserEmailTemplateCreate {
	auetc.mutation.SetAppID(u)
	return auetc
}

// SetUserID sets the "user_id" field.
func (auetc *AppUserEmailTemplateCreate) SetUserID(u uuid.UUID) *AppUserEmailTemplateCreate {
	auetc.mutation.SetUserID(u)
	return auetc
}

// SetLangID sets the "lang_id" field.
func (auetc *AppUserEmailTemplateCreate) SetLangID(u uuid.UUID) *AppUserEmailTemplateCreate {
	auetc.mutation.SetLangID(u)
	return auetc
}

// SetSubject sets the "subject" field.
func (auetc *AppUserEmailTemplateCreate) SetSubject(s string) *AppUserEmailTemplateCreate {
	auetc.mutation.SetSubject(s)
	return auetc
}

// SetBody sets the "body" field.
func (auetc *AppUserEmailTemplateCreate) SetBody(s string) *AppUserEmailTemplateCreate {
	auetc.mutation.SetBody(s)
	return auetc
}

// SetSender sets the "sender" field.
func (auetc *AppUserEmailTemplateCreate) SetSender(s string) *AppUserEmailTemplateCreate {
	auetc.mutation.SetSender(s)
	return auetc
}

// SetCreateAt sets the "create_at" field.
func (auetc *AppUserEmailTemplateCreate) SetCreateAt(u uint32) *AppUserEmailTemplateCreate {
	auetc.mutation.SetCreateAt(u)
	return auetc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auetc *AppUserEmailTemplateCreate) SetNillableCreateAt(u *uint32) *AppUserEmailTemplateCreate {
	if u != nil {
		auetc.SetCreateAt(*u)
	}
	return auetc
}

// SetUpdateAt sets the "update_at" field.
func (auetc *AppUserEmailTemplateCreate) SetUpdateAt(u uint32) *AppUserEmailTemplateCreate {
	auetc.mutation.SetUpdateAt(u)
	return auetc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (auetc *AppUserEmailTemplateCreate) SetNillableUpdateAt(u *uint32) *AppUserEmailTemplateCreate {
	if u != nil {
		auetc.SetUpdateAt(*u)
	}
	return auetc
}

// SetID sets the "id" field.
func (auetc *AppUserEmailTemplateCreate) SetID(u uuid.UUID) *AppUserEmailTemplateCreate {
	auetc.mutation.SetID(u)
	return auetc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (auetc *AppUserEmailTemplateCreate) SetNillableID(u *uuid.UUID) *AppUserEmailTemplateCreate {
	if u != nil {
		auetc.SetID(*u)
	}
	return auetc
}

// Mutation returns the AppUserEmailTemplateMutation object of the builder.
func (auetc *AppUserEmailTemplateCreate) Mutation() *AppUserEmailTemplateMutation {
	return auetc.mutation
}

// Save creates the AppUserEmailTemplate in the database.
func (auetc *AppUserEmailTemplateCreate) Save(ctx context.Context) (*AppUserEmailTemplate, error) {
	var (
		err  error
		node *AppUserEmailTemplate
	)
	auetc.defaults()
	if len(auetc.hooks) == 0 {
		if err = auetc.check(); err != nil {
			return nil, err
		}
		node, err = auetc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserEmailTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auetc.check(); err != nil {
				return nil, err
			}
			auetc.mutation = mutation
			if node, err = auetc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auetc.hooks) - 1; i >= 0; i-- {
			if auetc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auetc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auetc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auetc *AppUserEmailTemplateCreate) SaveX(ctx context.Context) *AppUserEmailTemplate {
	v, err := auetc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auetc *AppUserEmailTemplateCreate) Exec(ctx context.Context) error {
	_, err := auetc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auetc *AppUserEmailTemplateCreate) ExecX(ctx context.Context) {
	if err := auetc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auetc *AppUserEmailTemplateCreate) defaults() {
	if _, ok := auetc.mutation.CreateAt(); !ok {
		v := appuseremailtemplate.DefaultCreateAt()
		auetc.mutation.SetCreateAt(v)
	}
	if _, ok := auetc.mutation.UpdateAt(); !ok {
		v := appuseremailtemplate.DefaultUpdateAt()
		auetc.mutation.SetUpdateAt(v)
	}
	if _, ok := auetc.mutation.ID(); !ok {
		v := appuseremailtemplate.DefaultID()
		auetc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auetc *AppUserEmailTemplateCreate) check() error {
	if _, ok := auetc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUserEmailTemplate.app_id"`)}
	}
	if _, ok := auetc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppUserEmailTemplate.user_id"`)}
	}
	if _, ok := auetc.mutation.LangID(); !ok {
		return &ValidationError{Name: "lang_id", err: errors.New(`ent: missing required field "AppUserEmailTemplate.lang_id"`)}
	}
	if _, ok := auetc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "AppUserEmailTemplate.subject"`)}
	}
	if _, ok := auetc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "AppUserEmailTemplate.body"`)}
	}
	if _, ok := auetc.mutation.Sender(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required field "AppUserEmailTemplate.sender"`)}
	}
	if _, ok := auetc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppUserEmailTemplate.create_at"`)}
	}
	if _, ok := auetc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppUserEmailTemplate.update_at"`)}
	}
	return nil
}

func (auetc *AppUserEmailTemplateCreate) sqlSave(ctx context.Context) (*AppUserEmailTemplate, error) {
	_node, _spec := auetc.createSpec()
	if err := sqlgraph.CreateNode(ctx, auetc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (auetc *AppUserEmailTemplateCreate) createSpec() (*AppUserEmailTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUserEmailTemplate{config: auetc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuseremailtemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuseremailtemplate.FieldID,
			},
		}
	)
	_spec.OnConflict = auetc.conflict
	if id, ok := auetc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := auetc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuseremailtemplate.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := auetc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuseremailtemplate.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := auetc.mutation.LangID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuseremailtemplate.FieldLangID,
		})
		_node.LangID = value
	}
	if value, ok := auetc.mutation.Subject(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuseremailtemplate.FieldSubject,
		})
		_node.Subject = value
	}
	if value, ok := auetc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuseremailtemplate.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := auetc.mutation.Sender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuseremailtemplate.FieldSender,
		})
		_node.Sender = value
	}
	if value, ok := auetc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuseremailtemplate.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := auetc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuseremailtemplate.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserEmailTemplate.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserEmailTemplateUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (auetc *AppUserEmailTemplateCreate) OnConflict(opts ...sql.ConflictOption) *AppUserEmailTemplateUpsertOne {
	auetc.conflict = opts
	return &AppUserEmailTemplateUpsertOne{
		create: auetc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserEmailTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (auetc *AppUserEmailTemplateCreate) OnConflictColumns(columns ...string) *AppUserEmailTemplateUpsertOne {
	auetc.conflict = append(auetc.conflict, sql.ConflictColumns(columns...))
	return &AppUserEmailTemplateUpsertOne{
		create: auetc,
	}
}

type (
	// AppUserEmailTemplateUpsertOne is the builder for "upsert"-ing
	//  one AppUserEmailTemplate node.
	AppUserEmailTemplateUpsertOne struct {
		create *AppUserEmailTemplateCreate
	}

	// AppUserEmailTemplateUpsert is the "OnConflict" setter.
	AppUserEmailTemplateUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AppUserEmailTemplateUpsert) SetAppID(v uuid.UUID) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateAppID() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppUserEmailTemplateUpsert) SetUserID(v uuid.UUID) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateUserID() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldUserID)
	return u
}

// SetLangID sets the "lang_id" field.
func (u *AppUserEmailTemplateUpsert) SetLangID(v uuid.UUID) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldLangID, v)
	return u
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateLangID() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldLangID)
	return u
}

// SetSubject sets the "subject" field.
func (u *AppUserEmailTemplateUpsert) SetSubject(v string) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldSubject, v)
	return u
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateSubject() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldSubject)
	return u
}

// SetBody sets the "body" field.
func (u *AppUserEmailTemplateUpsert) SetBody(v string) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldBody, v)
	return u
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateBody() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldBody)
	return u
}

// SetSender sets the "sender" field.
func (u *AppUserEmailTemplateUpsert) SetSender(v string) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldSender, v)
	return u
}

// UpdateSender sets the "sender" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateSender() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldSender)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserEmailTemplateUpsert) SetCreateAt(v uint32) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateCreateAt() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserEmailTemplateUpsert) AddCreateAt(v uint32) *AppUserEmailTemplateUpsert {
	u.Add(appuseremailtemplate.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserEmailTemplateUpsert) SetUpdateAt(v uint32) *AppUserEmailTemplateUpsert {
	u.Set(appuseremailtemplate.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsert) UpdateUpdateAt() *AppUserEmailTemplateUpsert {
	u.SetExcluded(appuseremailtemplate.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserEmailTemplateUpsert) AddUpdateAt(v uint32) *AppUserEmailTemplateUpsert {
	u.Add(appuseremailtemplate.FieldUpdateAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUserEmailTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuseremailtemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserEmailTemplateUpsertOne) UpdateNewValues() *AppUserEmailTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuseremailtemplate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUserEmailTemplate.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserEmailTemplateUpsertOne) Ignore() *AppUserEmailTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserEmailTemplateUpsertOne) DoNothing() *AppUserEmailTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserEmailTemplateCreate.OnConflict
// documentation for more info.
func (u *AppUserEmailTemplateUpsertOne) Update(set func(*AppUserEmailTemplateUpsert)) *AppUserEmailTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserEmailTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserEmailTemplateUpsertOne) SetAppID(v uuid.UUID) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateAppID() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserEmailTemplateUpsertOne) SetUserID(v uuid.UUID) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateUserID() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateUserID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppUserEmailTemplateUpsertOne) SetLangID(v uuid.UUID) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateLangID() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetSubject sets the "subject" field.
func (u *AppUserEmailTemplateUpsertOne) SetSubject(v string) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateSubject() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateSubject()
	})
}

// SetBody sets the "body" field.
func (u *AppUserEmailTemplateUpsertOne) SetBody(v string) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateBody() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateBody()
	})
}

// SetSender sets the "sender" field.
func (u *AppUserEmailTemplateUpsertOne) SetSender(v string) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetSender(v)
	})
}

// UpdateSender sets the "sender" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateSender() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateSender()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserEmailTemplateUpsertOne) SetCreateAt(v uint32) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserEmailTemplateUpsertOne) AddCreateAt(v uint32) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateCreateAt() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserEmailTemplateUpsertOne) SetUpdateAt(v uint32) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserEmailTemplateUpsertOne) AddUpdateAt(v uint32) *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertOne) UpdateUpdateAt() *AppUserEmailTemplateUpsertOne {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *AppUserEmailTemplateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserEmailTemplateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserEmailTemplateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserEmailTemplateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserEmailTemplateUpsertOne.ID is not supported by MySQL driver. Use AppUserEmailTemplateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserEmailTemplateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserEmailTemplateCreateBulk is the builder for creating many AppUserEmailTemplate entities in bulk.
type AppUserEmailTemplateCreateBulk struct {
	config
	builders []*AppUserEmailTemplateCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUserEmailTemplate entities in the database.
func (auetcb *AppUserEmailTemplateCreateBulk) Save(ctx context.Context) ([]*AppUserEmailTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(auetcb.builders))
	nodes := make([]*AppUserEmailTemplate, len(auetcb.builders))
	mutators := make([]Mutator, len(auetcb.builders))
	for i := range auetcb.builders {
		func(i int, root context.Context) {
			builder := auetcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserEmailTemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, auetcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = auetcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, auetcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, auetcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (auetcb *AppUserEmailTemplateCreateBulk) SaveX(ctx context.Context) []*AppUserEmailTemplate {
	v, err := auetcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auetcb *AppUserEmailTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := auetcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auetcb *AppUserEmailTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := auetcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserEmailTemplate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserEmailTemplateUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (auetcb *AppUserEmailTemplateCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserEmailTemplateUpsertBulk {
	auetcb.conflict = opts
	return &AppUserEmailTemplateUpsertBulk{
		create: auetcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserEmailTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (auetcb *AppUserEmailTemplateCreateBulk) OnConflictColumns(columns ...string) *AppUserEmailTemplateUpsertBulk {
	auetcb.conflict = append(auetcb.conflict, sql.ConflictColumns(columns...))
	return &AppUserEmailTemplateUpsertBulk{
		create: auetcb,
	}
}

// AppUserEmailTemplateUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUserEmailTemplate nodes.
type AppUserEmailTemplateUpsertBulk struct {
	create *AppUserEmailTemplateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUserEmailTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuseremailtemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserEmailTemplateUpsertBulk) UpdateNewValues() *AppUserEmailTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuseremailtemplate.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUserEmailTemplate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserEmailTemplateUpsertBulk) Ignore() *AppUserEmailTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserEmailTemplateUpsertBulk) DoNothing() *AppUserEmailTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserEmailTemplateCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserEmailTemplateUpsertBulk) Update(set func(*AppUserEmailTemplateUpsert)) *AppUserEmailTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserEmailTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserEmailTemplateUpsertBulk) SetAppID(v uuid.UUID) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateAppID() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserEmailTemplateUpsertBulk) SetUserID(v uuid.UUID) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateUserID() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateUserID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppUserEmailTemplateUpsertBulk) SetLangID(v uuid.UUID) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateLangID() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetSubject sets the "subject" field.
func (u *AppUserEmailTemplateUpsertBulk) SetSubject(v string) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateSubject() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateSubject()
	})
}

// SetBody sets the "body" field.
func (u *AppUserEmailTemplateUpsertBulk) SetBody(v string) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetBody(v)
	})
}

// UpdateBody sets the "body" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateBody() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateBody()
	})
}

// SetSender sets the "sender" field.
func (u *AppUserEmailTemplateUpsertBulk) SetSender(v string) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetSender(v)
	})
}

// UpdateSender sets the "sender" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateSender() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateSender()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserEmailTemplateUpsertBulk) SetCreateAt(v uint32) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserEmailTemplateUpsertBulk) AddCreateAt(v uint32) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateCreateAt() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserEmailTemplateUpsertBulk) SetUpdateAt(v uint32) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserEmailTemplateUpsertBulk) AddUpdateAt(v uint32) *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserEmailTemplateUpsertBulk) UpdateUpdateAt() *AppUserEmailTemplateUpsertBulk {
	return u.Update(func(s *AppUserEmailTemplateUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *AppUserEmailTemplateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserEmailTemplateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserEmailTemplateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserEmailTemplateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
